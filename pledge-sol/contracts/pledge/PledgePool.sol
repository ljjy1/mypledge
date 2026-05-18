// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./interfaces/IDebtToken.sol";
import "./interfaces/IBscPledgeOracle.sol";
import "../uniswapv2/interfaces/IUniswapV2Router02.sol";
import "../uniswapv2/interfaces/IWETH.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title PledgePool
 * @notice 质押借贷池核心合约
 * @dev 该合约是整个借贷协议的核心入口，负责：
 *      1. 创建借贷池
 *      2. 存款人存入 lendToken
 *      3. 借款人质押 borrowToken
 *      4. 结算池子状态
 *      5. 到期完成/清算
 *      6. 用户领取 sp / jp 凭证代币
 *      7. 用户最终提现本金、利息或剩余抵押物
 */
contract PledgePool is ReentrancyGuard, Ownable {
    // 为 IERC20 启用 SafeERC20 安全方法
    using SafeERC20 for IERC20;

    // 计算精度，统一按 1e18 做比例换算  1e18 = 10**18  10的18次方
    uint256 constant internal calDecimal = 1e18;
    // 手续费和利率的精度基数，按 1e8 表示 1e8 = 10**8 10的8次方
    uint256 constant internal baseDecimal = 1e8;
    // 最小参与金额，默认 100e18
    uint256 public minAmount = 100e18;
    // 一年的秒数，用于按时间比例计算利息
    uint256 constant baseYear = 365 days;

    // 全局暂停标记，true 表示暂停所有受 notPause 修饰的操作
    bool public globalPaused = false;

    // 预言机合约地址
    address  public oracle;
    // DEX 路由地址，例如 PancakeSwap Router
    address public swapRouter;
    // 平台手续费接收地址
    address payable public feeAddress;

    // 出借方手续费率 单位1e8 如果是10 就表示10/(10 ** 8)
    uint256 public lendFee;
    // 借款方手续费率 单位1e8 如果是10 就表示10/(10 ** 8)
    uint256 public borrowFee;

    /**
     * @dev 池子的生命周期状态
     * MATCH       : 撮合中/匹配中，允许存款和质押
     * EXECUTION   : 已结算进入执行期
     * FINISH      : 到期正常结束
     * LIQUIDATION : 已触发清算
     * UNDONE      : 极端情况下未成立（如一边没有资金）
     */
    enum PoolState{MATCH, EXECUTION, FINISH, LIQUIDATION, UNDONE}

    // 新建池子的默认状态为 MATCH
    PoolState constant internal defaultState = PoolState.MATCH;

    //质押借贷池的基础信息
    struct PledgePoolInfo {
        //结算时间：时间到达后池子可以从 MATCH 进入 EXECUTION/UNDONE
        uint256 settleTime;
        //结束时间：到达后池子可以 finish
        uint256 endTime;
        //固定利率，单位 1e8
        uint256 interestRate;
        //出借池最大募集上限
        uint256 maxSupply;
        //当前已存入的 lendToken 总量
        uint256 lendSupply;
        //质押池已质押的 borrowToken 总量
        uint256 borrowSupply;
        //抵押率，单位 1e8
        uint256 mortgageRate;
        //出借资产地址，例如 BUSD
        address lendToken;
        //质押资产地址，例如 BTC
        address borrowToken;
        //当前池状态
        PoolState state;
        //出借方凭证 token地址 DebtToken
        address lendDebtToken;
        //质押方凭证 token地址 DebtToken
        address borrowDebtToken;
        //自动清算阈值
        uint256 autoLiquidateThreshold;
    }

    // 所有池子的基础信息数组，下标即 pid
    PledgePoolInfo[] public pledgePoolInfoList;

    // 每个池在不同阶段的实际结算数据
    struct PoolDataInfo {
        uint256 settleAmountLend;       // 结算时实际生效的出借token总量
        uint256 settleAmountBorrow;     // 结算时实际生效的质押token总量
        uint256 finishAmountLend;       // 正常结束时，出借侧最终可分配token总量
        uint256 finishAmountBorrow;     // 正常结束时，借款侧最终可取回抵押物token总量
        uint256 liquidationAmountLend;   // 清算后，出借侧最终可分配token总量
        uint256 liquidationAmountBorrow; // 清算后，借款侧最终可取回抵押物token总量
    }

    // 所有池子的阶段数据数组，下标与 pledgePoolInfo 对齐
    PoolDataInfo[] public poolDataInfoList;

    /**
     * @dev 创建池子时的入参结构体
     */
    struct CreatePoolParams {
        uint256 settleTime;             // 结算时间
        uint256 endTime;                // 结束时间
        uint256 interestRate;           // 固定利率，单位 1e8
        uint256 maxSupply;              // 出借池最大募集上限
        uint256 mortgageRate;           // 抵押率，单位 1e8
        address lendToken;              // 出借资产地址
        address borrowToken;            // 质押资产地址
        address lendDebtToken;          // 出借凭证地址
        address borrowDebtToken;        // 借款凭证地址
        uint256 autoLiquidateThreshold; // 自动清算阈值
    }


    //出借人信息
    struct LendInfo {
        //出借人存入的lendToken数量
        uint256 lendAmount;
        //出借人已退款数据(超募退回)
        uint256 refundLendAmount;
        //是否有退款 false=未退款，true=已退款
        bool hasNoRefund;
        //是否领取 false=未领取，true=已领取
        bool hasNoClaim;
    }

    //用户出借信息 映射 用户地址 => 池id => 出借人信息
    mapping(address => mapping(uint256 => LendInfo)) public lendInfoMap;


    //借款人信息
    struct BorrowInfo {
        //质押人质押的borrowToken数量
        uint256 borrowAmount;
        //借款人已退款数量（多余部分）
        uint256 returnBorrowAmount;
        // false=未退款，true=已退款
        bool hasNoRefund;
        // false=未领取，true=已领取
        bool hasNoClaim;
    }

    //用户质押信息 映射 用户地址 => 池id => 借款人信息
    mapping(address => mapping(uint256 => BorrowInfo)) public borrowInfoMap;


    // 设置费用事件
    event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee);
    // 设置预言机事件
    event SetOracle(address indexed oldOracle, address indexed newOracle);
    // 设置路由地址事件
    event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress);
    // 设置手续费接收地址事件
    event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress);
    // 设置最小金额事件
    event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount);
    // 设置全局暂停
    event SetGlobalPaused(bool indexed oldPaused, bool indexed newPaused);
    // 创建新池子事件 pid(池基础信息数组的索引)
    event CreatePledgePool(uint256 indexed pid);
    // 借款人质押事件
    event Borrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount);
    // 借款人退款事件
    event RefundBorrow(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount);
    // 借款人领取凭证和lendToken事件
    event ClaimBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 borrowDebtTokenAmount, uint256 lendTokenAmount);
    // 出借人销毁凭证领取本金利息事件
    event DestroyLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 burnAmount, uint256 redeemAmount);
    // 出借人领取凭证事件
    event ClaimLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount);
    // 出借人紧急提取事件
    event EmergencyWithdrawLend(uint256 indexed pid, address indexed token, address indexed lender, uint256 amount);
    // 借款人销毁凭证领取剩余抵押物事件
    event DestroyBorrowDebtToken(uint256 indexed pid, address indexed token, address indexed borrower, uint256 burnAmount, uint256 redeemAmount);
    // 借款人紧急提取事件
    event EmergencyWithdrawBorrow(uint256 indexed pid, address indexed token, address indexed borrower, uint256 amount);
    // 池结算事件
    event SettlePool(uint256 indexed pid, uint256 settleAmountLend, uint256 settleAmountBorrow);
    // 池正常完成事件
    event FinishPool(uint256 indexed pid, uint256 finishAmountLend, uint256 finishAmountBorrow);
    // 池清算事件
    event LiquidatePool(uint256 indexed pid, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow);
    /**
     * 出借事件
     * pid 池基础信息数组的索引
     * token 币种地址 0代表是ETH
     * lender 出借人地址
     * amount 出借金额
     */
    event Lend(uint256 indexed pid, address indexed token,address indexed lender, uint256 amount);
    /**
     * 退款事件
     * pid 池基础信息数组的索引
     * token 币种地址 0代表是ETH
     * refunder 退款人地址
     * amount 退款金额
     */
    event RefundLend(uint256 indexed pid, address indexed token,address indexed refunder, uint256 amount);

    /**
     * 构造函数
     * _oracle 预言机地址
     * _swapRouter DEX 路由地址
     *
     */
    constructor(
        address _oracle,
        address _swapRouter,
        address payable _feeAddress,
        address _owner
    )Ownable(_owner){
        //预言机地址不能为空
        require(_oracle != address(0), "Oracle address cannot be empty");
        //DEX 路由地址不能为空
        require(_swapRouter != address(0), "SwapRouter address cannot be empty");
        //平台手续费接收地址不能为空
        require(_feeAddress != address(0), "FeeAddress address cannot be empty");

        oracle = _oracle;
        swapRouter = _swapRouter;
        feeAddress = _feeAddress;
        // 默认手续费为 0
        lendFee = 0;
        borrowFee = 0;
    }

    /**
     * 设置出借手续费和借款手续费 只能由_owner调用 (当前合约_owner是多签地址)
     * _lendFee,_borrowFee单位为1e18 传入100 就表示收 100/(10**18)的比例的手续费
     */
    function setFee(uint256 _lendFee, uint256 _borrowFee) external onlyOwner {
        //_lendFee必须大于0且小于等于1e18
        require(_lendFee > 0 && _lendFee <= baseDecimal, "lendFee must be greater than 0 and less than or equal to 1e8");
        //_borrowFee必须大于0且小于等于1e18
        require(_borrowFee > 0 && _borrowFee <= baseDecimal, "borrowFee must be greater than 0 and less than or equal to 1e8");
        //记录事件
        emit SetFee(_lendFee, _borrowFee);
        //设置
        lendFee = _lendFee;
        borrowFee = _borrowFee;
    }

    /**
     * 设置预言机地址
     */
    function setOracle(address _oracle) external onlyOwner {
        //地址不能为0
        require(_oracle != address(0), "Oracle address cannot be empty");
        //记录事件
        emit SetOracle(oracle, _oracle);
        //修改新的地址
        oracle = _oracle;
    }

    /**
     * 设置DEX 路由地址
     */
    function setSwapRouter(address _swapRouter) external onlyOwner {
        //地址不能为0
        require(_swapRouter != address(0), "SwapRouter address cannot be empty");
        //记录事件
        emit SetSwapRouterAddress(swapRouter, _swapRouter);
        //修改新的地址
        swapRouter = _swapRouter;
    }

    /**
     * 设置手续费地址
     */
    function setFeeAddress(address payable _feeAddress) external onlyOwner {
        //地址不能为0
        require(_feeAddress != address(0), "FeeAddress address cannot be empty");
        //记录事件
        emit SetFeeAddress(feeAddress, _feeAddress);
        //修改新的地址
        feeAddress = _feeAddress;
    }

    /**
     * 设置最小参与金额
     */
    function setMinAmount(uint256 _minAmount) external onlyOwner {
        //必须大于0
        require(_minAmount > 0, "minAmount must be greater than 0");
        emit SetMinAmount(minAmount, _minAmount);
        minAmount = _minAmount;
    }

    /**
     * 设置全局暂停
     */
    function setGlobalPaused() external onlyOwner {
        emit SetGlobalPaused(globalPaused, !globalPaused);
        //通过取反实现开关
        globalPaused = !globalPaused;
    }

    /**
     * 创建质押借贷池
     * @param params 池子创建参数，使用 calldata 传入以节省 Gas
     * @return pid 创建成功后返回的池子索引，从0开始
     */
    function createPledgePool(
        CreatePoolParams calldata params
    ) public onlyOwner returns (uint256){
        require(params.settleTime > 0, "SettleTime must be greater than 0");
        require(params.endTime > 0, "EndTime must be greater than 0");
        require(params.endTime > params.settleTime, "EndTime must be greater than SettleTime");
        require(params.lendDebtToken != address(0), "LendDebtToken address cannot be empty");
        require(params.borrowDebtToken != address(0), "BorrowDebtToken address cannot be empty");
        require(params.mortgageRate > 0, "MortgageRate must be greater than 0");
        require(params.interestRate > 0, "InterestRate must be greater than 0");

        uint256 pid = pledgePoolInfoList.length;
        pledgePoolInfoList.push();
        PledgePoolInfo storage pool = pledgePoolInfoList[pid];
        pool.settleTime = params.settleTime;
        pool.endTime = params.endTime;
        pool.interestRate = params.interestRate;
        pool.maxSupply = params.maxSupply;
        pool.mortgageRate = params.mortgageRate;
        pool.lendToken = params.lendToken;
        pool.borrowToken = params.borrowToken;
        pool.state = PoolState.MATCH;
        pool.lendDebtToken = params.lendDebtToken;
        pool.borrowDebtToken = params.borrowDebtToken;
        pool.autoLiquidateThreshold = params.autoLiquidateThreshold;

        poolDataInfoList.push();

        emit CreatePledgePool(pid);
        return pid;
    }

    /**
     * @dev 获取池状态，返回枚举对应的 uint256
     */
    function getPoolState(uint256 _pid) public view returns (uint256) {
        // 读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        // 返回状态编码
        return uint256(pool.state);
    }

    /**
     * 出借人存入lendToken
     * @param _pid 池子ID
     * @param _lendAmount 存入金额
     */
    function lend(uint256 _pid, uint256 _lendAmount) external
    payable //可传入ETH
    nonReentrant  //不可以重入
    notPaused   //未暂停
    beforeSettleTime(_pid)  //必须在结算时间之前
    onlyStateMatch(_pid)   //池状态必须是MATCH
    {
        //读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //读取出借信息
        LendInfo storage lendInfo = lendInfoMap[msg.sender][_pid];
        //判断池子存入资产
        if(pool.lendToken == address(0)){
            //判断ETH需要大于0
            require(msg.value > 0, "ETH must be greater than 0");
            //判断需要大于池子的最小金额
            require(msg.value >= minAmount, "ETH must be greater than minAmount");
            //不能超过池子最大上限
            require(pool.maxSupply >= pool.lendSupply + msg.value, "Exceeds the maximum limit");

            _lendAmount = msg.value;
        }else{
            //说明是ERC20 判断金额大于0
            require(_lendAmount > 0, "ERC20 must be greater than 0");
            //判断大于池子最小金额
            require(_lendAmount >= minAmount, "ERC20 must be greater than minAmount");
            //不能超过池子最大上限
            require(pool.maxSupply >= pool.lendSupply + _lendAmount, "Exceeds the maximum limit");

            IERC20 erc20 = IERC20(pool.lendToken);
            //判断用户账户余额
            require(erc20.balanceOf(msg.sender) >= _lendAmount, "ERC20 balance is not enough");
            //进行安全转账到当前合约
            erc20.safeTransferFrom(msg.sender, address(this), _lendAmount);

        }
        //添加出借人存入的lendToken数量
        lendInfo.lendAmount = lendInfo.lendAmount + _lendAmount;
        //添加池子存入的lendToken数量
        pool.lendSupply = pool.lendSupply + _lendAmount;
        // 重置用户标记，表示后续可以 claim / refund
        lendInfo.hasNoClaim = false;
        lendInfo.hasNoRefund = false;

        //记录出借事件
        emit Lend(_pid, pool.lendToken,msg.sender, _lendAmount);
    }


    /**
     * 出借人领取超募退款
     */
    function refundLend(uint256 _pid) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    afterSettleTime(_pid)  //必须在结算时间之后
    onlyStateNotMatchUndone(_pid)   //池状态必须不是 MATCH/UNDONE，而是 EXECUTION、FINISH、LIQUIDATION 之一
    {
        //读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //池阶段信息
        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        //读取出借信息
        LendInfo storage lendInfo = lendInfoMap[msg.sender][_pid];

        //必须出借过
        require(lendInfo.lendAmount > 0, "Not lend");
        //必须存在可退款金额 (池子已存入的lendToken总量 - 结算时实际出借的金额)
        require(pool.lendSupply - poolDataInfo.settleAmountLend > 0, "No refund amount");
        //不能重复退款
        require(!lendInfo.hasNoRefund, "Has no refund");

        //出借人份额比例 出借人存入的lendToken数量* 计算精度 / 池子总存入的 lendToken 总量
        uint256 userShare = lendInfo.lendAmount * calDecimal / pool.lendSupply;
        //计算退款金额  (池子总存入的lendToken总量 - 池子结算时实际生效的出借lendToken量) * 出借人份额比例 / 计算精度
        uint256 refundAmount = (pool.lendSupply - poolDataInfo.settleAmountLend) * userShare / calDecimal;
        //退还金额需要大于0
        require(refundAmount > 0, "Refund amount must be greater than 0");
        //标记已退款
        lendInfo.hasNoRefund = true;
        //添加可退款金额
        lendInfo.refundLendAmount = lendInfo.refundLendAmount + refundAmount;

        //退还资产转账
        if(pool.lendToken == address(0)){
            //退还ETH
            payable(msg.sender).transfer(refundAmount);
        }else{
            //退还ERC20
            IERC20 erc20 = IERC20(pool.lendToken);
            //判断当前账户余额
            require(erc20.balanceOf(address(this)) >= refundAmount, "ERC20 balance is not enough");
            //进行安全转账到出借人
            erc20.safeTransfer(msg.sender, refundAmount);
        }

        //记录退款事件
        emit RefundLend(_pid, pool.lendToken,msg.sender, refundAmount);
    }


    /**
     * 出借人领取借款凭证代币lendDebtToken
     */
    function claimLendDebtToken(uint256 _pid) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    afterSettleTime(_pid)  //必须在结算时间之后
    onlyStateNotMatchUndone(_pid)   //池状态必须不是 MATCH/UNDONE，而是 EXECUTION、FINISH、LIQUIDATION 之一
    {
        //读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //池阶段信息
        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        //读取出借信息
        LendInfo storage lendInfo = lendInfoMap[msg.sender][_pid];

        //必须出借过
        require(lendInfo.lendAmount > 0, "Not lend");
        //不能重复领取
        require(!lendInfo.hasNoClaim, "Has no claim");
        //出借人份额比例 出借人存入的lendToken数量 * 计算精度 / 池子总存入的 lendToken 总量
        uint256 userShare = lendInfo.lendAmount * calDecimal / pool.lendSupply;
        // 按份额计算用户应得 Token (结算后的有效出借金额 * 出借人份额比例) / 计算精度
        uint256 claimLendDebtTokenAmount = poolDataInfo.settleAmountLend * userShare / calDecimal;

        //铸造债务凭证代币DebtToken给出借人 (需要多签钱包提前设置当前合约能进行pool.lendDebtToken铸币)
        IDebtToken(pool.lendDebtToken).mint(msg.sender, claimLendDebtTokenAmount);
        //标记已经领取凭证代币
        lendInfo.hasNoClaim = true;
        //记录事件
        emit ClaimLendDebtToken(_pid, pool.lendDebtToken,msg.sender, claimLendDebtTokenAmount);
    }

    /**
     * 出借人销毁债务凭证代币lendDebtToken 取回本金和利息
     * @param _pid 池ID
     * @param _lendDebtTokenAmount 需要销毁债务凭证代币数量
     */
    function destroyLendDebtToken(uint256 _pid,uint256 _lendDebtTokenAmount) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    onlyStateFinishOrLiquidation(_pid)   //池状态必须是 FINISH 或 LIQUIDATION
    {
        //读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //池阶段信息
        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        //读取出借信息
        LendInfo storage lendInfo = lendInfoMap[msg.sender][_pid];
        //必须出借过
        require(lendInfo.lendAmount > 0, "Not lend");
        //判断销毁token大于0
        require(_lendDebtTokenAmount > 0, "Destroy amount must be greater than 0");
        //销毁token 这里不判断用户_lendDebtTokenAmount够不够 销毁的时候会判断用户token余额
        IDebtToken(pool.lendDebtToken).burn(msg.sender, _lendDebtTokenAmount);

        //计算需要销毁的token 占总token的份额比例 poolDataInfo.settleAmountLend就是实际可以铸币的endDebtToken数量
        uint256 userShare = _lendDebtTokenAmount * calDecimal / poolDataInfo.settleAmountLend;
        // 正常到期完成场景
        if(pool.state == PoolState.FINISH){
            //必须超过结束时间
            require(block.timestamp >= pool.endTime, "Not reached end time");
            // 按份额分配 finish 阶段的 lend 收益
            uint256 redeemAmount = poolDataInfo.finishAmountLend * userShare / calDecimal ;

            //领取本金和收益
            if(pool.lendToken == address(0)){
                //退还ETH
                payable(msg.sender).transfer(redeemAmount);
            }else{
                IERC20 erc20 = IERC20(pool.lendToken);
                //判断当前账户余额
                require(erc20.balanceOf(address(this)) >= redeemAmount, "ERC20 balance is not enough");
                //进行安全转账到出借人
                erc20.safeTransfer(msg.sender, redeemAmount);
            }

            //记录事件
            emit DestroyLendDebtToken(_pid, pool.lendDebtToken,msg.sender, _lendDebtTokenAmount, redeemAmount);
        }else{
            //被清算场景
            // 只要超过结算时间即可
            require(block.timestamp > pool.settleTime, "Not reached settle time");
            // 按份额分配 LIQUIDATION 阶段的 lend 收益
            uint256 redeemAmount = poolDataInfo.liquidationAmountLend * userShare / calDecimal ;

            //领取本金和收益
            if(pool.lendToken == address(0)){
                //退还ETH
                payable(msg.sender).transfer(redeemAmount);
            }else{
                IERC20 erc20 = IERC20(pool.lendToken);
                //判断当前账户余额
                require(erc20.balanceOf(address(this)) >= redeemAmount, "ERC20 balance is not enough");
                //进行安全转账到出借人
                erc20.safeTransfer(msg.sender, redeemAmount);
            }
            //记录事件
            emit DestroyLendDebtToken(_pid, pool.lendDebtToken,msg.sender, _lendDebtTokenAmount, redeemAmount);
        }
    }


    /**
     * 出借人紧急提取
     * 只有池子进入 UNDONE（未成立）状态时才允许
     */
    function emergencyWithdrawLend(uint256 _pid) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    onlyStateUndone(_pid)   //池状态必须为 UNDONE
    {
        //读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //读取出借信息
        LendInfo storage lendInfo = lendInfoMap[msg.sender][_pid];
        // 池中必须确实存在出借总额
        require(pool.lendSupply > 0, "No lend supply");
        //必须出借过
        require(lendInfo.lendAmount > 0, "No lend amount");
        //判断用户是否已经退款了
        require(!lendInfo.hasNoRefund, "Already refunded");

        //全额退回存入token
        if(pool.lendToken == address(0)){
            payable(msg.sender).transfer(lendInfo.lendAmount);
        }else{
            IERC20 erc20 = IERC20(pool.lendToken);
            require(erc20.balanceOf(address(this)) >= lendInfo.lendAmount, "ERC20 balance is not enough");
            erc20.safeTransfer(msg.sender, lendInfo.lendAmount);
        }
        // 标记为已退款
        lendInfo.hasNoRefund = true;
        //记录事件
        emit EmergencyWithdrawLend(_pid, pool.lendToken,msg.sender, lendInfo.lendAmount);
    }


    /**
     * 借款人质押 borrowToken
     * @param _pid 池ID
     * @param _borrowTokenAmount 质押数量
     */
    function borrow(uint256 _pid,uint256 _borrowTokenAmount) external payable
    nonReentrant  //不可以重入
    notPaused   //未暂停
    beforeSettleTime(_pid)  //必须在结算时间之前
    onlyStateMatch(_pid)   //池状态必须为 MATCH
    {
        //读取池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //获取借款人信息
        BorrowInfo storage borrowInfo = borrowInfoMap[msg.sender][_pid];
        if(pool.borrowToken == address(0)){
            //判断用户ETH账户余额要大于0
            require(msg.value > 0, "ETH must be greater than 0");
            _borrowTokenAmount = msg.value;
        }else{
            //_borrowTokenAmount必须大于0
            require(_borrowTokenAmount > 0, "borrowTokenAmount must be greater than 0");
            IERC20 erc20 = IERC20(pool.borrowToken);
            require(erc20.balanceOf(msg.sender) >= _borrowTokenAmount, "ERC20 balance is not enough");
            //进行转账
            erc20.safeTransferFrom(msg.sender, address(this), _borrowTokenAmount);
        }
        //出借人质押数量累加
        borrowInfo.borrowAmount += _borrowTokenAmount;
        //池中总质押数量累加
        pool.borrowSupply += _borrowTokenAmount;
        //出借人退款状态重置
        borrowInfo.hasNoRefund = false;
        //借款人无收益状态重置
        borrowInfo.hasNoClaim = false;
        //记录事件
        emit Borrow(_pid, pool.borrowToken,msg.sender, _borrowTokenAmount);
    }



    /**
     * @dev 借款人领取多余质押物退款
     * @notice 仅允许在结算后、且池状态不是 MATCH/UNDONE 时执行
     * @param _pid 池索引
     */
    function refundBorrow(uint256 _pid) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    afterSettleTime(_pid)    // 池状态必须是结算之后
    onlyStateNotMatchUndone(_pid)   //池状态必须不是 MATCH/UNDONE，而是 EXECUTION、FINISH、LIQUIDATION 之一
    {
        //池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //池阶段数据
        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        //获取借款人信息
        BorrowInfo storage borrowInfo = borrowInfoMap[msg.sender][_pid];

        //必须参与质押过
        require(borrowInfo.borrowAmount > 0, "No borrow amount");
        //需要存在可退款的质押物
        require(pool.borrowSupply > poolDataInfo.settleAmountBorrow, "No refund amount");
        //不能重复退款
        require(!borrowInfo.hasNoRefund, "Already refunded");
        //计算用户质押物的份额 (用户质押数量 * 计算精度) / 池中总质押数量
        uint256 userShare = borrowInfo.borrowAmount * calDecimal / pool.borrowSupply;
        //按份额分配多余质押物 (池中总质押数量 - 结算时实际生效的质押token总量) * 份额 / 计算精度
        uint256 refundAmount = (pool.borrowSupply - poolDataInfo.settleAmountBorrow) * userShare / calDecimal;

        //退回质押物
        if(pool.borrowToken == address(0)){
            payable(msg.sender).transfer(refundAmount);
        }else{
            IERC20 erc20 = IERC20(pool.borrowToken);
            require(erc20.balanceOf(address(this)) >= refundAmount, "ERC20 balance is not enough");
            erc20.safeTransfer(msg.sender, refundAmount);
        }
        //标记为已退款
        borrowInfo.hasNoRefund = true;
        borrowInfo.returnBorrowAmount += refundAmount;
        //记录事件
        emit RefundBorrow(_pid, pool.borrowToken,msg.sender, refundAmount);
    }


    /**
     * 借款人领取borrowDebtToken 和可以借出的lendToken
     */
    function claimBorrow(uint256 _pid) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    afterSettleTime(_pid) // 池状态必须是结算之后
    onlyStateNotMatchUndone(_pid)
    {
        //池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //池阶段数据
        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        //获取借款人信息
        BorrowInfo storage borrowInfo = borrowInfoMap[msg.sender][_pid];

        //必须参与质押过
        require(borrowInfo.borrowAmount > 0, "No borrow amount");
        //不能重复领取
        require(!borrowInfo.hasNoClaim, "Already claimed");

        //可领取borrowDebtToken总量  实际借出金额 * (抵押率 / 1e8) 1e8是抵押率的单位需要去除 例如:抵押率传入1e8 代表抵押率100% 1:1
        uint256 borrowDebtTokenAmount = poolDataInfo.settleAmountLend * pool.mortgageRate / baseDecimal;

        //计算用户占比 (用户质押数量 * 计算精度) / 池中总质押数量
        uint256 userShare = borrowInfo.borrowAmount * calDecimal / pool.borrowSupply;
        //按占比计算用户可领取borrowDebtToken  可领取borrowDebtToken总量 * 占比 / 计算精度
        uint256 claimBorrowDebtTokenAmount = borrowDebtTokenAmount * userShare / calDecimal;

        //进行铸币 (需要多签钱包提前设置当前合约能进行pool.borrowDebtToken铸币)
        IDebtToken(pool.borrowDebtToken).mint(msg.sender, claimBorrowDebtTokenAmount);

        //用户可借出的lendToken数量 实际借出金额 * 用户占比 / 计算精度
        uint256 claimLendTokenAmount = poolDataInfo.settleAmountLend * userShare / calDecimal;
        if(pool.lendToken == address(0)){
            payable(msg.sender).transfer(claimLendTokenAmount);
        }else{
            IERC20 erc20 = IERC20(pool.lendToken);
            require(erc20.balanceOf(address(this)) >= claimLendTokenAmount, "ERC20 balance is not enough");
            erc20.safeTransfer(msg.sender, claimLendTokenAmount);
        }
        // 标记已领取
        borrowInfo.hasNoClaim = true;
        //记录事件
        emit ClaimBorrow(_pid, pool.borrowDebtToken,msg.sender, claimBorrowDebtTokenAmount, claimLendTokenAmount);
    }


    /**
     * 借款人销毁borrowDebtToken 取回剩余抵押物
     */
    function destroyBorrowDebtToken(uint256 _pid,uint256 _borrowDebtTokenAmount) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    onlyStateFinishOrLiquidation(_pid)
    {
        //池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //池阶段数据
        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        //销毁数量必须大于0
        require(_borrowDebtTokenAmount > 0, "borrowDebtTokenAmount must be greater than 0");
        //销毁borrowDebtToken
        IDebtToken(pool.borrowDebtToken).burn(msg.sender, _borrowDebtTokenAmount);

        //可领取borrowDebtToken总量  实际借出金额 * (抵押率 / 1e8) 1e8是抵押率的单位需要去除 例如:抵押率传入1e8 代表抵押率100% 1:1
        uint256 totalBorrowDebtTokenAmount = poolDataInfo.settleAmountLend * pool.mortgageRate / baseDecimal;

        //计算_borrowDebtTokenAmount销毁占比 (用户销毁数量 * 计算精度) / 总量
        uint256 userShare = _borrowDebtTokenAmount * calDecimal / totalBorrowDebtTokenAmount;
        //正常完成场景
        if (pool.state == PoolState.FINISH) {
            //必须超过结束时间
            require(block.timestamp >= pool.endTime, "Not reached end time");
            //按份额分配借款侧可取回的抵押物
            uint256 redeemAmount = poolDataInfo.finishAmountBorrow * userShare / calDecimal;

            //转回抵押物
            if (pool.borrowToken == address(0)) {
                payable(msg.sender).transfer(redeemAmount);
            } else {
                IERC20 erc20 = IERC20(pool.borrowToken);
                require(erc20.balanceOf(address(this)) >= redeemAmount, "ERC20 balance is not enough");
                erc20.safeTransfer(msg.sender, redeemAmount);
            }

            emit DestroyBorrowDebtToken(_pid, pool.borrowToken, msg.sender, _borrowDebtTokenAmount, redeemAmount);
        } else {
            //清算场景
            //只需超过结算时间即可
            require(block.timestamp > pool.settleTime, "Not reached settle time");
            //按份额分配清算后借款侧可取回的抵押物
            uint256 redeemAmount = poolDataInfo.liquidationAmountBorrow * userShare / calDecimal;

            //转回抵押物
            if (pool.borrowToken == address(0)) {
                payable(msg.sender).transfer(redeemAmount);
            } else {
                IERC20 erc20 = IERC20(pool.borrowToken);
                require(erc20.balanceOf(address(this)) >= redeemAmount, "ERC20 balance is not enough");
                erc20.safeTransfer(msg.sender, redeemAmount);
            }

            emit DestroyBorrowDebtToken(_pid, pool.borrowToken, msg.sender, _borrowDebtTokenAmount, redeemAmount);
        }
    }


    /**
     * 借款人紧急提取
     * 只有池子进入 UNDONE（未成立）状态时才允许
     * @param _pid 池ID
     */
    function emergencyWithdrawBorrow(uint256 _pid) external
    nonReentrant  //不可以重入
    notPaused   //未暂停
    onlyStateUndone(_pid)   //池状态必须为 UNDONE
    {
        //池信息
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        //借款人信息
        BorrowInfo storage borrowInfo = borrowInfoMap[msg.sender][_pid];

        //池中必须存在质押总额
        require(pool.borrowSupply > 0, "No borrow supply");
        //必须参与质押过
        require(borrowInfo.borrowAmount > 0, "No borrow amount");
        //不能重复退款
        require(!borrowInfo.hasNoRefund, "Already refunded");

        //全额退回质押token
        if (pool.borrowToken == address(0)) {
            payable(msg.sender).transfer(borrowInfo.borrowAmount);
        } else {
            IERC20 erc20 = IERC20(pool.borrowToken);
            require(erc20.balanceOf(address(this)) >= borrowInfo.borrowAmount, "ERC20 balance is not enough");
            erc20.safeTransfer(msg.sender, borrowInfo.borrowAmount);
        }
        //标记为已退款
        borrowInfo.hasNoRefund = true;

        emit EmergencyWithdrawBorrow(_pid, pool.borrowToken, msg.sender, borrowInfo.borrowAmount);
    }


    /**
     * 检查是否到达可结算时间
     * @param _pid 池ID
     * @return 可以结算返回 true，否则返回 false
     */
    function checkCanSettle(uint256 _pid) public view returns (bool) {
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        if (pool.state != PoolState.MATCH) {
            return false;
        }
        if (block.timestamp < pool.settleTime) {
            return false;
        }
        if (pool.lendSupply == 0 || pool.borrowSupply == 0) {
            return false;
        }
        return true;
    }


    /**
     * @notice 对借贷池执行结算，将池子从 MATCH 转为 EXECUTION 或 UNDONE
     * @dev 根据双方资产价格计算实际匹配的出借量和质押量
     *      当 borrowToken == lendToken（含双方均为 ETH 的情况），
     *      价格比恒为 1:1，跳过预言机查询以节省 Gas 并避免对同资产重复查价
     *      当任一 token 为 address(0) 代表 ETH，预言机使用 underlying=0 查询
     * @param _pid 池ID
     */
    function settlePool(uint256 _pid) external
    onlyOwner
    {
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];

        require(pool.state == PoolState.MATCH, "Pool state must be MATCH");
        require(block.timestamp >= pool.settleTime, "Not reached settle time");

        if (pool.lendSupply == 0 || pool.borrowSupply == 0) {
            pool.state = PoolState.UNDONE;
            return;
        }

        uint256 actualLend;
        uint256 actualBorrow;

        if (pool.lendToken == pool.borrowToken) {
            // 相同资产（含双方均为 address(0) ETH 的情况）：
            // 价格比恒为 1:1，无需查询预言机，避免冗余调用及价格未配置时的除零异常
            // matchedLendByCollateral = borrowSupply * 1e8 / mortgageRate
            uint256 matchedLendByCollateral = pool.borrowSupply * baseDecimal / pool.mortgageRate;

            actualLend = pool.lendSupply <= matchedLendByCollateral ? pool.lendSupply : matchedLendByCollateral;
            actualBorrow = actualLend * pool.mortgageRate / baseDecimal;
        } else {
            uint256 borrowTokenPrice = IBscPledgeOracle(oracle).getPrice(pool.borrowToken);
            uint256 lendTokenPrice = IBscPledgeOracle(oracle).getPrice(pool.lendToken);

            // 价格必须 > 0，避免除零异常；address(0) 代表 ETH 时依赖预言机 underlying=0 价格
            require(borrowTokenPrice > 0, "settlePool: borrow token price is zero");
            require(lendTokenPrice > 0, "settlePool: lend token price is zero");

            // 质押token总量换算为出借token等值，再根据抵押率计算最大可匹配出借量
            uint256 matchedLendByCollateral = pool.borrowSupply * borrowTokenPrice * baseDecimal
                / (lendTokenPrice * pool.mortgageRate);

            actualLend = pool.lendSupply <= matchedLendByCollateral ? pool.lendSupply : matchedLendByCollateral;

            // 反算实际生效的质押量
            actualBorrow = actualLend * pool.mortgageRate * lendTokenPrice
                / (baseDecimal * borrowTokenPrice);
        }

        if (actualBorrow == 0) {
            pool.state = PoolState.UNDONE;
            return;
        }

        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];
        poolDataInfo.settleAmountLend = actualLend;
        poolDataInfo.settleAmountBorrow = actualBorrow;

        pool.state = PoolState.EXECUTION;

        emit SettlePool(_pid, actualLend, actualBorrow);
    }


    /**
     * 检查是否已经到达 finish 时间
     * @param _pid 池ID
     * @return 可以完成返回 true，否则返回 false
     */
    function checkCanFinish(uint256 _pid) public view returns (bool) {
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        if (pool.state != PoolState.EXECUTION) {
            return false;
        }
        if (block.timestamp < pool.endTime) {
            return false;
        }
        return true;
    }


    /**
     * 正常完成函数
     * 池子到期后，计算出借方和借款方的最终可分配金额
     * @param _pid 池ID
     */
    function finishPool(uint256 _pid) external
    onlyOwner
    {
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];

        require(pool.state == PoolState.EXECUTION, "Pool state must be EXECUTION");
        require(block.timestamp >= pool.endTime, "Not reached end time");

        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];

        // 利息 = 本金 * 利率 * 时长 / (一年秒数 * 利率精度)
        uint256 duration = pool.endTime - pool.settleTime;
        uint256 lendInterest = poolDataInfo.settleAmountLend * pool.interestRate * duration
            / (baseYear * baseDecimal);

        // 出借方应收总额 = 本金 + 利息
        uint256 finishAmountLend = poolDataInfo.settleAmountLend + lendInterest;

        // 平台从总应付出借金额中提取的手续费（lendFee 单位 1e8）
        uint256 lendFeeAmount = finishAmountLend * lendFee / baseDecimal;

        // 包括手续费的 swap 目标输出：出借方应收 + 手续费
        uint256 sellAmount = finishAmountLend + lendFeeAmount;

        if (pool.borrowToken == pool.lendToken) {
            // 相同资产（含双方均为 ETH）：无需 DEX 交换
            poolDataInfo.finishAmountLend = finishAmountLend;

            uint256 remainBeforeFee = poolDataInfo.settleAmountBorrow >= sellAmount
                ? poolDataInfo.settleAmountBorrow - sellAmount
                : 0;
            poolDataInfo.finishAmountBorrow = _redeemFees(borrowFee, pool.borrowToken, remainBeforeFee);

            if (lendFeeAmount > 0) {
                _transferToken(pool.lendToken, feeAddress, lendFeeAmount);
            }
        } else {
            // 不同资产：通过 Router swapExact* 出售 borrowToken 换回 lendToken
            address weth = IUniswapV2Router02(swapRouter).weth();
            address tokenIn = pool.borrowToken == address(0) ? weth : pool.borrowToken;
            address tokenOut = pool.lendToken == address(0) ? weth : pool.lendToken;

            if (tokenIn == tokenOut) {
                // 同一底层资产（如 borrowToken=WETH, lendToken=ETH），无需 AMM 交换
                // 仅做 wrap/unwrap
                if (pool.borrowToken == address(0)) {
                    IWETH(weth).deposit{value: sellAmount}();
                }

                poolDataInfo.finishAmountLend = finishAmountLend;

                if (lendFeeAmount > 0) {
                    _transferToken(pool.lendToken, feeAddress, lendFeeAmount);
                }

                uint256 remainBeforeFee = poolDataInfo.settleAmountBorrow >= sellAmount
                    ? poolDataInfo.settleAmountBorrow - sellAmount
                    : 0;
                poolDataInfo.finishAmountBorrow = _redeemFees(borrowFee, pool.borrowToken, remainBeforeFee);
            } else {
                // 通过 Router 进行单次原子 swap，使用 exact-input 模式避免滑点问题
                uint256 amountSell = _getAmountIn(pool.borrowToken, pool.lendToken, sellAmount);
                require(amountSell <= poolDataInfo.settleAmountBorrow,
                    "finishPool: insufficient borrow collateral");

                uint256 amountIn = _swapExactIn(pool.borrowToken, pool.lendToken, amountSell);
                // 至少换回本金+利息，否则滑点过大
                require(amountIn >= finishAmountLend, "finishPool: Slippage too high");

                // 超出 lendAmount 的部分作为手续费转给 feeAddress
                if (amountIn > finishAmountLend) {
                    uint256 feeAmount = amountIn - finishAmountLend;
                    _transferToken(pool.lendToken, feeAddress, feeAmount);
                    poolDataInfo.finishAmountLend = finishAmountLend;
                } else {
                    poolDataInfo.finishAmountLend = amountIn;
                }

                // 剩余抵押物扣除 borrowFee 后为借款人可取回
                uint256 remainNow = poolDataInfo.settleAmountBorrow >= amountSell
                    ? poolDataInfo.settleAmountBorrow - amountSell
                    : 0;
                poolDataInfo.finishAmountBorrow = _redeemFees(borrowFee, pool.borrowToken, remainNow);
            }
        }

        pool.state = PoolState.FINISH;

        emit FinishPool(_pid, poolDataInfo.finishAmountLend, poolDataInfo.finishAmountBorrow);
    }


    /**
     * 检查是否满足清算条件
     * 当前抵押物价值/借款价值 < autoLiquidateThreshold/baseDecimal 时返回 true
     * @param _pid 池ID
     * @return 可以清算返回 true，否则返回 false
     */
    function checkCanLiquidate(uint256 _pid) public view returns (bool) {
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];
        if (pool.state != PoolState.EXECUTION) {
            return false;
        }
        if (block.timestamp >= pool.endTime) {
            return false;
        }

        IBscPledgeOracle oracleContract = IBscPledgeOracle(oracle);
        uint256 lendTokenPrice = oracleContract.getPrice(pool.lendToken);
        uint256 borrowTokenPrice = oracleContract.getPrice(pool.borrowToken);

        //当前抵押物价值（以lendToken等值计价）
        uint256 currentCollateralValue = poolDataInfoList[_pid].settleAmountBorrow * borrowTokenPrice / lendTokenPrice;
        //当前借款价值
        uint256 borrowValue = poolDataInfoList[_pid].settleAmountLend;

        //当前抵押率 = 抵押物价值 * baseDecimal / 借款价值
        uint256 currentCollateralRate = currentCollateralValue * baseDecimal / borrowValue;

        return currentCollateralRate < pool.autoLiquidateThreshold;
    }


    /**
     * 对池子执行清算
     * 清算后，出借方可按凭证取回本金，借款方取回剩余抵押物
     * @param _pid 池ID
     */
    function liquidatePool(uint256 _pid) external
    nonReentrant
    notPaused
    {
        PledgePoolInfo storage pool = pledgePoolInfoList[_pid];

        require(pool.state == PoolState.EXECUTION, "Pool state must be EXECUTION");
        require(block.timestamp < pool.endTime, "Pool already ended, use finishPool");

        PoolDataInfo storage poolDataInfo = poolDataInfoList[_pid];

        // 利息计算与 finishPool 一致
        uint256 duration = pool.endTime - pool.settleTime;
        uint256 interest = poolDataInfo.settleAmountLend * pool.interestRate * duration
            / (baseYear * baseDecimal);

        // 出借方应收 = 本金 + 利息
        uint256 lendAmount = poolDataInfo.settleAmountLend + interest;

        // 包括 lendFee 的总目标
        uint256 lendFeeAmount = lendAmount * lendFee / baseDecimal;
        uint256 sellAmount = lendAmount + lendFeeAmount;

        if (pool.borrowToken == pool.lendToken) {
            // 相同资产：无需 DEX 交换
            poolDataInfo.liquidationAmountLend = lendAmount;

            uint256 remainBeforeFee = poolDataInfo.settleAmountBorrow >= sellAmount
                ? poolDataInfo.settleAmountBorrow - sellAmount
                : 0;
            poolDataInfo.liquidationAmountBorrow = _redeemFees(borrowFee, pool.borrowToken, remainBeforeFee);

            if (lendFeeAmount > 0) {
                _transferToken(pool.lendToken, feeAddress, lendFeeAmount);
            }
        } else {
            address weth = IUniswapV2Router02(swapRouter).weth();
            address tokenIn = pool.borrowToken == address(0) ? weth : pool.borrowToken;
            address tokenOut = pool.lendToken == address(0) ? weth : pool.lendToken;

            if (tokenIn == tokenOut) {
                // 同一底层资产，仅做 wrap/unwrap
                if (pool.borrowToken == address(0)) {
                    IWETH(weth).deposit{value: sellAmount}();
                }

                poolDataInfo.liquidationAmountLend = lendAmount;

                if (lendFeeAmount > 0) {
                    _transferToken(pool.lendToken, feeAddress, lendFeeAmount);
                }

                uint256 remainBeforeFee = poolDataInfo.settleAmountBorrow >= sellAmount
                    ? poolDataInfo.settleAmountBorrow - sellAmount
                    : 0;
                poolDataInfo.liquidationAmountBorrow = _redeemFees(borrowFee, pool.borrowToken, remainBeforeFee);
            } else {
                // 通过 Router swapExact* 出售 borrowToken
                uint256 amountSell = _getAmountIn(pool.borrowToken, pool.lendToken, sellAmount);
                require(amountSell <= poolDataInfo.settleAmountBorrow,
                    "liquidatePool: insufficient borrow collateral");

                uint256 amountIn = _swapExactIn(pool.borrowToken, pool.lendToken, amountSell);
                require(amountIn >= lendAmount, "liquidatePool: Slippage too high");

                if (amountIn > lendAmount) {
                    uint256 feeAmount = amountIn - lendAmount;
                    _transferToken(pool.lendToken, feeAddress, feeAmount);
                    poolDataInfo.liquidationAmountLend = lendAmount;
                } else {
                    poolDataInfo.liquidationAmountLend = amountIn;
                }

                uint256 remainNow = poolDataInfo.settleAmountBorrow >= amountSell
                    ? poolDataInfo.settleAmountBorrow - amountSell
                    : 0;
                poolDataInfo.liquidationAmountBorrow = _redeemFees(borrowFee, pool.borrowToken, remainNow);
            }
        }

        pool.state = PoolState.LIQUIDATION;

        emit LiquidatePool(_pid, poolDataInfo.liquidationAmountLend, poolDataInfo.liquidationAmountBorrow);
    }












    // ============ 内部辅助函数 ============

    /// @dev 内部转账函数，address(0)=ETH 否则 ERC20
    function _transferToken(address token, address to, uint256 amount) internal {
        if (token == address(0)) {
            payable(to).transfer(amount);
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    /// @dev 内部查询合约内 token 余额，address(0)=ETH 余额
    function _tokenBalance(address token) internal view returns (uint256) {
        return token == address(0) ? address(this).balance : IERC20(token).balanceOf(address(this));
    }

    /// @dev 构建 swap 路径（address(0) 替换为 WETH）
    function _getSwapPath(address borrowToken, address lendToken) internal view returns (address[] memory path) {
        address weth = IUniswapV2Router02(swapRouter).weth();
        path = new address[](2);
        path[0] = borrowToken == address(0) ? weth : borrowToken;
        path[1] = lendToken == address(0) ? weth : lendToken;
    }

    /// @dev 计算为得到 amountOut 需要卖出多少 borrowToken（含滑点缓冲）
    function _getAmountIn(address borrowToken, address lendToken, uint256 amountOut) internal view returns (uint256) {
        address[] memory path = _getSwapPath(borrowToken, lendToken);
        uint256[] memory amounts = IUniswapV2Router02(swapRouter).getAmountsIn(amountOut, path);
        return amounts[0];
    }

    /// @dev 通过 Router swapExact* 执行卖出，返回实际收到的 lendToken 数量
    function _swapExactIn(address borrowToken, address lendToken, uint256 amountIn) internal returns (uint256) {
        IUniswapV2Router02 router = IUniswapV2Router02(swapRouter);
        address[] memory path = _getSwapPath(borrowToken, lendToken);
        uint256[] memory amounts;

        if (borrowToken == address(0)) {
            // 原生 ETH → ERC20
            amounts = router.swapExactETHForTokens{value: amountIn}(
                0, path, address(this), block.timestamp + 300
            );
        } else {
            IERC20(borrowToken).forceApprove(swapRouter, amountIn);
            if (lendToken == address(0)) {
                // ERC20 → 原生 ETH
                amounts = router.swapExactTokensForETH(
                    amountIn, 0, path, address(this), block.timestamp + 300
                );
            } else {
                // ERC20 → ERC20
                amounts = router.swapExactTokensForTokens(
                    amountIn, 0, path, address(this), block.timestamp + 300
                );
            }
        }

        return amounts[amounts.length - 1];
    }

    /// @dev 扣除手续费并转给 feeAddress，返回扣费后净额
    function _redeemFees(uint256 feeRatio, address token, uint256 amount) internal returns (uint256) {
        uint256 fee = amount * feeRatio / baseDecimal;
        if (fee > 0) {
            _transferToken(token, feeAddress, fee);
        }
        return amount - fee;
    }

    // ============ Modifiers ============

    /**
    * 未暂停修饰器
    */
    modifier notPaused(){
        require(!globalPaused, "Global Paused");
        _;
    }

    /**
     * 结算时间之前修饰器
     */
    modifier beforeSettleTime(uint256 _pid){
        require(block.timestamp < pledgePoolInfoList[_pid].settleTime, "Less than this time");
        _;
    }

    /**
     * 结算时间之后修饰器
     */
    modifier afterSettleTime(uint256 _pid){
        require(block.timestamp >= pledgePoolInfoList[_pid].settleTime, "Greater than this time");
        _;
    }

    /**
     * 池状态必须是MATCH
     */
    modifier onlyStateMatch(uint256 _pid){
        require(pledgePoolInfoList[_pid].state == PoolState.MATCH, "Pool state must be MATCH");
        _;
    }

    /**
     * 池状态必须不是 MATCH/UNDONE，而是 EXECUTION、FINISH、LIQUIDATION 之一
     */
    modifier onlyStateNotMatchUndone(uint256 _pid){
        require(pledgePoolInfoList[_pid].state != PoolState.MATCH && pledgePoolInfoList[_pid].state != PoolState.UNDONE, "Pool state must not be MATCH or UNDONE");
        _;
    }

    /**
     * 池状态必须是 FINISH 或 LIQUIDATION
     */
    modifier onlyStateFinishOrLiquidation(uint256 _pid){
        require(pledgePoolInfoList[_pid].state == PoolState.FINISH || pledgePoolInfoList[_pid].state == PoolState.LIQUIDATION, "Pool state must be FINISH or LIQUIDATION");
        _;
    }

    /**
     * 池状态必须是 UNDONE
     */
    modifier onlyStateUndone(uint256 _pid){
        require(pledgePoolInfoList[_pid].state == PoolState.UNDONE, "Pool state must be UNDONE");
        _;
    }


}
