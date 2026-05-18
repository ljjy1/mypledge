// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// ============ Safe 多签辅助库 ============

import {SafeHelper} from "../mocks/SafeHelper.sol";
// ============ 被测试合约 ============

import {PledgePool} from "./PledgePool.sol";
// ============ 依赖合约 ============

import {DebtToken} from "./DebtToken.sol";
import {BscPledgeOracle} from "./BscPledgeOracle.sol";
// ============ Mock / 基础设施 ============

import {MockTestERC20} from "../mocks/MockTestERC20.sol";
import {WETH} from "../uniswapv2/WETH.sol";
import {UniswapV2Factory} from "../uniswapv2/UniswapV2Factory.sol";
import {UniswapV2Router02} from "../uniswapv2/UniswapV2Router02.sol";
// ============ OpenZeppelin & Safe 类型 ============

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Enum} from "@safe-global/safe-smart-account/contracts/libraries/Enum.sol";

/**
 * @title PledgePoolTest
 * @notice PledgePool 借贷池合约的完整 Solidity 测试套件
 * @dev 继承 SafeHelper，使用 Safe 多签钱包作为合约 owner，模拟生产环境的多签管理
 *      覆盖测试范围与 PledgePool.ts 中的 39 个 Mocha 测试完全一致
 */
contract PledgePoolTest is SafeHelper {

    // ==================== 测试常量 ====================

    // 基础小数位数：1e8，用于抵押率、利率等百分比参数
    uint256 constant BASE_DECIMAL = 1e8;
    // 利率精度（1e8 中的 5e7 = 5%）
    uint256 constant INTEREST_RATE = 5e7;
    // 抵押率：2e8 = 200%，即借出的 lendToken 需要 2 倍价值的抵押品
    uint256 constant MORTGAGE_RATE = 2e8;
    // 清算阈值：7e7 = 70%，当抵押率低于此值时触发清算
    uint256 constant LIQUIDATE_THRESHOLD = 7e7;
    // 最小存入金额：100 * 1e18
    uint256 constant MIN_AMOUNT = 100e18;
    // 池子最大募集上限：1,000,000 * 1e18
    uint256 constant MAX_SUPPLY = 1_000_000e18;
    // 测试代币总供应量：10,000,000 * 1e18
    uint256 constant TOKEN_SUPPLY = 10_000_000e18;

    // ==================== 合约实例 ====================

    // Uniswap V2 基础设施
    WETH internal weth;
    UniswapV2Factory internal uniFactory;
    UniswapV2Router02 internal router;
    // 测试代币 — tokenA 为出借资产（lendToken），tokenB 为质押资产（borrowToken）
    MockTestERC20 internal tokenA;
    MockTestERC20 internal tokenB;
    // 预言机合约
    BscPledgeOracle internal oracle;
    // 债务凭证代币 — 出借侧 SP Token，借款侧 JP Token
    DebtToken internal lendDebt;
    DebtToken internal borrowDebt;
    // 被测试的核心借贷池
    PledgePool internal pool;

    // ==================== 测试账号 ====================

    // 使用 makeAddr 创建可追溯的确定性地址，方便调试
    address internal lender = makeAddr("lender");
    address internal lender2 = makeAddr("lender2");
    address internal borrower = makeAddr("borrower");
    address internal borrower2 = makeAddr("borrower2");
    address internal feeCollector = makeAddr("feeCollector");
    address internal liquidator = makeAddr("liquidator");
    address internal random = makeAddr("random");

    // ==================== 池子数据 ====================

    // 默认池 ID（部署后自动创建的第一个池子，ID 为 0）
    uint256 internal pid;
    // 结算时间（到达后可从 MATCH 转为 EXECUTION/UNDONE）
    uint256 internal settleTime;
    // 结束时间（到达后可执行 finish）
    uint256 internal endTime;

    // ==================== 事件声明（用于 vm.expectEmit 断言） ====================

    // 管理事件
    event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee);
    event SetOracle(address indexed oldOracle, address indexed newOracle);
    event SetSwapRouterAddress(address indexed oldSwapAddress, address indexed newSwapAddress);
    event SetFeeAddress(address indexed oldFeeAddress, address indexed newFeeAddress);
    event SetMinAmount(uint256 indexed oldMinAmount, uint256 indexed newMinAmount);
    event SetGlobalPaused(bool indexed oldPaused, bool indexed newPaused);
    // 池操作事件
    event CreatePledgePool(uint256 indexed pid);
    event Lend(uint256 indexed pid, address indexed token, address indexed lender_, uint256 amount);
    event Borrow(uint256 indexed pid, address indexed token, address indexed borrower_, uint256 amount);
    event SettlePool(uint256 indexed pid, uint256 settleAmountLend, uint256 settleAmountBorrow);
    event FinishPool(uint256 indexed pid, uint256 finishAmountLend, uint256 finishAmountBorrow);
    event LiquidatePool(uint256 indexed pid, uint256 liquidationAmountLend, uint256 liquidationAmountBorrow);
    event ClaimLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender_, uint256 amount);
    event ClaimBorrow(uint256 indexed pid, address indexed token, address indexed borrower_, uint256 lendDebtTokenAmount, uint256 lendTokenAmount);
    event DestroyLendDebtToken(uint256 indexed pid, address indexed token, address indexed lender_, uint256 burnAmount, uint256 redeemAmount);
    event DestroyBorrowDebtToken(uint256 indexed pid, address indexed token, address indexed borrower_, uint256 burnAmount, uint256 redeemAmount);
    event EmergencyWithdrawLend(uint256 indexed pid, address indexed token, address indexed lender_, uint256 amount);
    event EmergencyWithdrawBorrow(uint256 indexed pid, address indexed token, address indexed borrower_, uint256 amount);
    event RefundLend(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount);
    event RefundBorrow(uint256 indexed pid, address indexed token, address indexed refunder, uint256 amount);
    // 代币事件
    event MinterAdded(address indexed minter, bool status);
    event DebtMinted(address indexed account, uint256 amount);
    event DebtBurned(address indexed account, uint256 amount);

    // ==================== setUp — 每次测试前的完整部署 ====================

    /**
     * @notice 每个测试用例前的初始化函数
     * @dev 完整部署一套测试环境：Uniswap、代币、预言机、债务代币、Safe 多签、PledgePool
     *      与 TypeScript 中的 deployBase() 功能完全一致
     */
    function setUp() public {
        // 部署 Safe 多签钱包（双签名阈值），使用 SafeHelper 的内置函数
        _deploySafe(0xA11CE, 0xB0B, 2);

        // ============ 部署 Uniswap V2 基础设施 ============

        weth = new WETH();
        uniFactory = new UniswapV2Factory(address(this));
        router = new UniswapV2Router02(address(uniFactory), address(weth));

        // ============ 部署测试代币 ============

        // 部署模拟 ERC20 代币，初始供应量全部铸造给当前合约（测试合约）
        tokenA = new MockTestERC20("LendToken", "LEND", TOKEN_SUPPLY);
        tokenB = new MockTestERC20("BorrowToken", "BORR", TOKEN_SUPPLY);

        // 授权 Uniswap 路由器从合约转移代币，用于添加流动性
        tokenA.approve(address(router), MAX_SUPPLY);
        tokenB.approve(address(router), MAX_SUPPLY);
        // 在 Uniswap 上为 tokenA/tokenB 交易对添加流动性（各 500,000 * 1e18）
        router.addLiquidity(
            address(tokenA), address(tokenB),
            500_000e18, 500_000e18,
            0, 0, address(this), block.timestamp + 1 hours
        );

        // ============ 部署预言机 ============

        // 预言机 owner 为测试合约本身，这样可以直接设置价格（无需多签）
        oracle = new BscPledgeOracle(address(this));
        // 设置 tokenA 和 tokenB 的价格均为 1 USD（精度 1e18）
        oracle.setPrice(address(tokenA), 1e18);
        oracle.setPrice(address(tokenB), 1e18);

        // ============ 部署债务代币 ============

        // 部署出借债务代币 SP Token，owner 为测试合约本身
        lendDebt = new DebtToken("SP Token", "sp", address(this));
        // 部署质押债务代币 JP Token，owner 为测试合约本身
        borrowDebt = new DebtToken("JP Token", "jp", address(this));

        // ============ 部署 PledgePool ============

        // 部署核心借贷池，owner 为 Safe 多签地址
        pool = new PledgePool(address(oracle), address(router), payable(feeCollector), safeAddress);
        // 授权 PledgePool 可以铸造和销毁两种债务代币
        lendDebt.setMinter(address(pool), true);
        borrowDebt.setMinter(address(pool), true);

        // ============ 创建默认借贷池 ============

        // 结算时间设为 100 天后（确保测试开始时池子处于 MATCH 状态）
        settleTime = block.timestamp + 100 days;
        // 结束时间设为结算后 30 天
        endTime = settleTime + 30 days;

        // 通过 Safe 多签创建第一个借贷池（pid = 0）
        _executeViaSafe(address(pool), abi.encodeCall(
            pool.createPledgePool,
            (PledgePool.CreatePoolParams({
                settleTime: settleTime,
                endTime: endTime,
                interestRate: INTEREST_RATE,
                maxSupply: MAX_SUPPLY,
                mortgageRate: MORTGAGE_RATE,
                lendToken: address(tokenA),
                borrowToken: address(tokenB),
                lendDebtToken: address(lendDebt),
                borrowDebtToken: address(borrowDebt),
                autoLiquidateThreshold: LIQUIDATE_THRESHOLD
            }))
        ));
        pid = 0;

        // ============ 分发代币给测试账号 ============

        // 给出借人分发 tokenA（lendToken）
        tokenA.transfer(lender, 500_000e18);
        tokenA.transfer(lender2, 500_000e18);
        // 给借款人分发 tokenB（borrowToken）
        tokenB.transfer(borrower, 500_000e18);
        tokenB.transfer(borrower2, 500_000e18);
        // 给清算人分发 ETH（gas 费）
        vm.deal(liquidator, 100 ether);
        vm.deal(random, 100 ether);

        // 给测试账号授权 PledgePool 合约可以转移代币
        vm.startPrank(lender);
        tokenA.approve(address(pool), type(uint256).max);
        vm.stopPrank();

        vm.startPrank(lender2);
        tokenA.approve(address(pool), type(uint256).max);
        vm.stopPrank();

        vm.startPrank(borrower);
        tokenB.approve(address(pool), type(uint256).max);
        vm.stopPrank();

        vm.startPrank(borrower2);
        tokenB.approve(address(pool), type(uint256).max);
        vm.stopPrank();
    }

    // ==================== 内部辅助函数 ====================

    /// @dev 出借人存入 tokenA，借款人质押 tokenB 的快捷函数
    function _depositAndBorrow(uint256 lendAmount, uint256 borrowAmount) internal {
        vm.prank(lender);
        pool.lend(pid, lendAmount);

        vm.prank(borrower);
        pool.borrow(pid, borrowAmount);
    }

    /// @dev 推进时间并结算（通过 Safe 多签），适合仅有一种资产的池子
    function _settlePool() internal {
        vm.warp(settleTime + 1);
        _executeViaSafe(address(pool), abi.encodeCall(pool.settlePool, (pid)));
    }

    /// @dev 推进时间到结束时间之后并执行 finishPool（通过 Safe 多签）
    function _finishPool() internal {
        vm.warp(endTime + 1);
        _executeViaSafe(address(pool), abi.encodeCall(pool.finishPool, (pid)));
    }

    /// @dev 设置出借和借款手续费的快捷函数（通过 Safe 多签）
    function _setFee(uint256 lendFee_, uint256 borrowFee_) internal {
        _executeViaSafe(address(pool), abi.encodeCall(pool.setFee, (lendFee_, borrowFee_)));
    }

    /// @dev 在 same-token 池子场景下，给借款人分发 tokenA 并授权
    function _setupBorrowerForSameTokenPool() internal {
        // 同币种池子使用 tokenA 作为抵押品，需要给借款人分发 tokenA
        tokenA.transfer(borrower, 500_000e18);
        vm.startPrank(borrower);
        tokenA.approve(address(pool), type(uint256).max);
        vm.stopPrank();
    }

    // ==================== 元组解构辅助函数 ====================

    /// @dev 安全读取池子 lendSupply，避免 inline 解构导致 stack too deep
    function _lendSupply(uint256 _pid) internal view returns (uint256) {
        (,,,, uint256 s,,,,,,,,) = pool.pledgePoolInfoList(_pid);
        return s;
    }

    /// @dev 安全读取池子 borrowSupply
    function _borrowSupply(uint256 _pid) internal view returns (uint256) {
        (,,,,, uint256 s,,,,,,,) = pool.pledgePoolInfoList(_pid);
        return s;
    }

    /// @dev 安全读取结算借出量 settleAmountLend
    function _settleAmountLend(uint256 _pid) internal view returns (uint256) {
        (uint256 s,,,,,) = pool.poolDataInfoList(_pid);
        return s;
    }

    /// @dev 安全读取结算质押量 settleAmountBorrow
    function _settleAmountBorrow(uint256 _pid) internal view returns (uint256) {
        (,uint256 s,,,,) = pool.poolDataInfoList(_pid);
        return s;
    }

    /// @dev 安全读取 finishAmountLend
    function _finishAmountLend(uint256 _pid) internal view returns (uint256) {
        (,,uint256 s,,,) = pool.poolDataInfoList(_pid);
        return s;
    }

    /// @dev 安全读取 finishAmountBorrow
    function _finishAmountBorrow(uint256 _pid) internal view returns (uint256) {
        (,,,uint256 s,,) = pool.poolDataInfoList(_pid);
        return s;
    }

    /// @dev 安全读取 liquidationAmountLend
    function _liquidationAmountLend(uint256 _pid) internal view returns (uint256) {
        (,,,,uint256 s,) = pool.poolDataInfoList(_pid);
        return s;
    }

    /// @dev 安全读取用户出借金额 lendAmount
    function _lendAmount(address user, uint256 _pid) internal view returns (uint256) {
        (uint256 s,,,) = pool.lendInfoMap(user, _pid);
        return s;
    }

    // ============================================================
    // 1. 部署 & 初始状态
    // 测试部署后合约的基本配置是否正确
    // ============================================================

    /**
     * @notice 验证部署后 owner、oracle、swapRouter、feeAddress 等关键字段是否正确设置
     */
    function test_deployment_setsOwnerOracleRouter() public {
        // owner 必须是 Safe 多签地址
        assertEq(pool.owner(), safeAddress);
        // oracle 地址
        assertEq(pool.oracle(), address(oracle));
        // swap 路由地址
        assertEq(pool.swapRouter(), address(router));
        // 手续费地址
        assertEq(pool.feeAddress(), feeCollector);
        // 初始出借手续费为 0
        assertEq(pool.lendFee(), 0);
        // 初始借款手续费为 0
        assertEq(pool.borrowFee(), 0);
        // 默认最小金额
        assertEq(pool.minAmount(), MIN_AMOUNT);
        // 全局暂停初始为 false
        assertFalse(pool.globalPaused());
    }

    // ============================================================
    // 2. 管理函数
    // 测试 onlyOwner 权限控制、参数校验和状态变更
    // ============================================================

    /**
     * @notice 非 owner 调用 setFee 应被 OwnableUnauthorizedAccount 拒绝
     */
    function test_setFee_revertsForNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        pool.setFee(1, 1);
    }

    /**
     * @notice 通过 Safe 多签成功设置手续费
     */
    function test_setFee_updatesViaSafe() public {
        // 通过 Safe 设置出借手续费 0.5%（5e6/1e8），借款手续费 0.3%（3e6/1e8）
        _setFee(5e6, 3e6);
        assertEq(pool.lendFee(), 5e6);
        assertEq(pool.borrowFee(), 3e6);
    }

    /**
     * @notice 通过 Safe 设置出借手续费为 0 应被拒绝
     */
    function test_setFee_revertsZeroLendFee() public {
        vm.prank(safeAddress);
        vm.expectRevert("lendFee must be greater than 0 and less than or equal to 1e8");
        pool.setFee(0, 1);
    }

    /**
     * @notice 通过 Safe 设置借款手续费超过 1e8 应被拒绝
     */
    function test_setFee_revertsBorrowFeeTooHigh() public {
        vm.prank(safeAddress);
        vm.expectRevert("borrowFee must be greater than 0 and less than or equal to 1e8");
        pool.setFee(1, BASE_DECIMAL + 1);
    }

    /**
     * @notice 通过 Safe 设置 setOracle 为零地址应被拒绝
     */
    function test_setOracle_revertsZeroAddress() public {
        vm.prank(safeAddress);
        vm.expectRevert("Oracle address cannot be empty");
        pool.setOracle(address(0));
    }

    /**
     * @notice 通过 Safe 成功设置 setOracle
     */
    function test_setOracle_updatesAddress() public {
        vm.expectEmit(true, true, false, true);
        emit SetOracle(address(oracle), random);
        _executeViaSafe(address(pool), abi.encodeCall(pool.setOracle, (random)));
        assertEq(pool.oracle(), random);
    }

    /**
     * @notice 通过 Safe 成功设置 setSwapRouter
     */
    function test_setSwapRouter_updatesAddress() public {
        _executeViaSafe(address(pool), abi.encodeCall(pool.setSwapRouter, (random)));
        assertEq(pool.swapRouter(), random);
    }

    /**
     * @notice 通过 Safe 成功设置 setFeeAddress
     */
    function test_setFeeAddress_updatesAddress() public {
        _executeViaSafe(address(pool), abi.encodeCall(pool.setFeeAddress, (payable(random))));
        assertEq(pool.feeAddress(), random);
    }

    /**
     * @notice 通过 Safe 成功设置 setMinAmount
     */
    function test_setMinAmount_updatesAmount() public {
        _executeViaSafe(address(pool), abi.encodeCall(pool.setMinAmount, (200e18)));
        assertEq(pool.minAmount(), 200e18);
    }

    /**
     * @notice 通过 Safe 切换全局暂停状态
     */
    function test_setGlobalPaused_togglesState() public {
        // 第一次调用：开启暂停
        _executeViaSafe(address(pool), abi.encodeCall(pool.setGlobalPaused, ()));
        assertTrue(pool.globalPaused());
        // 第二次调用：关闭暂停
        _executeViaSafe(address(pool), abi.encodeCall(pool.setGlobalPaused, ()));
        assertFalse(pool.globalPaused());
    }

    // ============================================================
    // 3. 池创建
    // 测试 createPledgePool 的参数校验和字段初始化
    // ============================================================

    /**
     * @notice 创建一个新的借贷池并验证字段正确初始化
     */
    function test_createPool_createsWithCorrectFields() public {
        // 设置新的池子参数（结算时间为 1 天后，便于测试）
        uint256 newSettleTime = block.timestamp + 1 days;
        uint256 newEndTime = newSettleTime + 30 days;

        vm.expectEmit(true, false, false, true);
        emit CreatePledgePool(1);
        _executeViaSafe(address(pool), abi.encodeCall(
            pool.createPledgePool,
            (PledgePool.CreatePoolParams({
                settleTime: newSettleTime,
                endTime: newEndTime,
                interestRate: INTEREST_RATE,
                maxSupply: MAX_SUPPLY,
                mortgageRate: MORTGAGE_RATE,
                lendToken: address(tokenA),
                borrowToken: address(tokenB),
                lendDebtToken: address(lendDebt),
                borrowDebtToken: address(borrowDebt),
                autoLiquidateThreshold: LIQUIDATE_THRESHOLD
            }))
        ));

        // 读取新池子的信息
        (uint256 settleTime_, uint256 endTime_,,,,,,,, PledgePool.PoolState state_,,,) = pool.pledgePoolInfoList(1);
        assertEq(settleTime_, newSettleTime);
        assertEq(endTime_, newEndTime);
        assertTrue(state_ == PledgePool.PoolState.MATCH);
        assertEq(settleTime_, newSettleTime);
        assertEq(endTime_, newEndTime);
        assertTrue(state_ == PledgePool.PoolState.MATCH);
    }

    /**
     * @notice 验证 createPledgePool 参数校验：settleTime 为 0 应被拒绝
     */
    function test_createPool_revertsZeroSettleTime() public {
        PledgePool.CreatePoolParams memory params = _defaultPoolParams();
        params.settleTime = 0;
        vm.prank(safeAddress);
        vm.expectRevert("SettleTime must be greater than 0");
        pool.createPledgePool(params);
    }

    /**
     * @notice 验证 endTime 为 0 应被拒绝
     */
    function test_createPool_revertsZeroEndTime() public {
        PledgePool.CreatePoolParams memory params = _defaultPoolParams();
        params.endTime = 0;
        vm.prank(safeAddress);
        vm.expectRevert("EndTime must be greater than 0");
        pool.createPledgePool(params);
    }

    /**
     * @notice 验证 endTime 等于 settleTime 应被拒绝
     */
    function test_createPool_revertsEndTimeEqSettleTime() public {
        PledgePool.CreatePoolParams memory params = _defaultPoolParams();
        params.endTime = params.settleTime;
        vm.prank(safeAddress);
        vm.expectRevert("EndTime must be greater than SettleTime");
        pool.createPledgePool(params);
    }

    /// @dev 构造默认池子参数的快捷函数
    function _defaultPoolParams() internal view returns (PledgePool.CreatePoolParams memory) {
        return PledgePool.CreatePoolParams({
            settleTime: block.timestamp + 1 days,
            endTime: block.timestamp + 1 days + 30 days,
            interestRate: INTEREST_RATE,
            maxSupply: MAX_SUPPLY,
            mortgageRate: MORTGAGE_RATE,
            lendToken: address(tokenA),
            borrowToken: address(tokenB),
            lendDebtToken: address(lendDebt),
            borrowDebtToken: address(borrowDebt),
            autoLiquidateThreshold: LIQUIDATE_THRESHOLD
        });
    }

    // ============================================================
    // 4. 存入
    // 测试 lend 函数的正常流程、金额校验和时间限制
    // ============================================================

    /**
     * @notice 正常的 lend 操作：状态更新和事件发送
     */
    function test_lend_updatesStateAndEmitsEvent() public {
        uint256 amount = 1000e18;

        vm.expectEmit(true, true, true, true);
        emit Lend(pid, address(tokenA), lender, amount);
        vm.prank(lender);
        pool.lend(pid, amount);

        // 验证池子信息中的 lendSupply 和出借金额
        assertEq(_lendSupply(pid), amount);
        assertEq(_lendAmount(lender, pid), amount);
    }

    /**
     * @notice lend 低于最小金额应被拒绝
     */
    function test_lend_revertsBelowMinAmount() public {
        vm.prank(lender);
        vm.expectRevert("ERC20 must be greater than minAmount");
        pool.lend(pid, MIN_AMOUNT - 1);
    }

    /**
     * @notice lend 超过最大供应量应被拒绝
     */
    function test_lend_revertsExceedsMaxSupply() public {
        vm.prank(lender);
        vm.expectRevert("Exceeds the maximum limit");
        pool.lend(pid, MAX_SUPPLY + 1);
    }

    /**
     * @notice lend 在结算时间之后应被拒绝
     */
    function test_lend_revertsAfterSettleTime() public {
        vm.warp(settleTime + 1);

        vm.prank(lender);
        vm.expectRevert("Less than this time");
        pool.lend(pid, 1000e18);
    }

    // ============================================================
    // 5. 质押
    // 测试 borrow 函数的正常流程和时间限制
    // ============================================================

    /**
     * @notice 正常的 borrow 操作：状态更新和事件发送
     */
    function test_borrow_updatesStateAndEmitsEvent() public {
        uint256 amount = 2000e18;

        vm.expectEmit(true, true, true, true);
        emit Borrow(pid, address(tokenB), borrower, amount);
        vm.prank(borrower);
        pool.borrow(pid, amount);

        // 验证池子信息中的 borrowSupply
        assertEq(_borrowSupply(pid), amount);
    }

    /**
     * @notice borrow 在结算时间之后应被拒绝
     */
    function test_borrow_revertsAfterSettleTime() public {
        vm.warp(settleTime + 1);

        vm.prank(borrower);
        vm.expectRevert("Less than this time");
        pool.borrow(pid, 1000e18);
    }

    // ============================================================
    // 6. 结算
    // 测试 settlePool 函数的状态转换和金额计算
    // ============================================================

    /**
     * @notice 正常结算：有 lend 和 borrow 时池子进入 EXECUTION 状态
     */
    function test_settle_transitionsToExecution() public {
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 验证状态为 EXECUTION（1）
        assertEq(pool.getPoolState(pid), 1);

        // 验证结算数据正确
        assertEq(_settleAmountLend(pid), 1000e18);
        assertEq(_settleAmountBorrow(pid), 2000e18);
    }

    /**
     * @notice 零供应量结算：没有 lend 或 borrow 时池子进入 UNDONE 状态
     */
    function test_settle_zeroSupplyGoesUndone() public {
        vm.warp(settleTime + 1);
        _executeViaSafe(address(pool), abi.encodeCall(pool.settlePool, (pid)));

        // 验证状态为 UNDONE（4）
        assertEq(pool.getPoolState(pid), 4);
    }

    /**
     * @notice 在结算时间之前调用 settlePool 应被拒绝
     */
    function test_settle_revertsBeforeSettleTime() public {
        _depositAndBorrow(1000e18, 2000e18);
        // 不推进时间，当前时间仍在 settleTime 之前
        vm.prank(safeAddress);
        vm.expectRevert("Not reached settle time");
        pool.settlePool(pid);
    }

    // ============================================================
    // 7. 领取凭证
    // 测试出借人和借款人领取债务代币凭证
    // ============================================================

    /**
     * @notice 出借人领取 lendDebtToken（SP Token）
     */
    function test_claim_lendDebtToken() public {
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        vm.prank(lender);
        pool.claimLendDebtToken(pid);

        assertEq(lendDebt.balanceOf(lender), 1000e18);
    }

    /**
     * @notice 借款人领取 borrowDebtToken（JP Token）和可借出的 lendToken
     */
    function test_claim_borrow() public {
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        vm.prank(borrower);
        pool.claimBorrow(pid);

        // 验证借款人收到了 JP Token
        assertEq(borrowDebt.balanceOf(borrower), 2000e18);
    }

    // ============================================================
    // 8. finishPool — same token
    // 测试 lendToken 和 borrowToken 为同一代币时的 finish 流程
    // ============================================================

    /**
     * @notice 同币种池子的 finish 操作：验证 FINISH 状态以及 lend 金额增加（含利息）
     */
    function test_finishPool_sameToken() public {
        // 创建同币种池子（lendToken = borrowToken = tokenA）
        uint256 newSettleTime = block.timestamp + 7 days;
        uint256 newEndTime = newSettleTime + 30 days;

        _executeViaSafe(address(pool), abi.encodeCall(
            pool.createPledgePool,
            (PledgePool.CreatePoolParams({
                settleTime: newSettleTime,
                endTime: newEndTime,
                interestRate: INTEREST_RATE,
                maxSupply: MAX_SUPPLY,
                mortgageRate: MORTGAGE_RATE,
                lendToken: address(tokenA),
                borrowToken: address(tokenA),
                lendDebtToken: address(lendDebt),
                borrowDebtToken: address(borrowDebt),
                autoLiquidateThreshold: LIQUIDATE_THRESHOLD
            }))
        ));
        uint256 samePid = 1;

        // 给借款人分发 tokenA 并授权
        _setupBorrowerForSameTokenPool();

        // 出借人存入 1000 tokenA，借款人质押 2000 tokenA
        vm.prank(lender);
        pool.lend(samePid, 1000e18);

        vm.prank(borrower);
        pool.borrow(samePid, 2000e18);

        // 结算
        vm.warp(newSettleTime + 1);
        _executeViaSafe(address(pool), abi.encodeCall(pool.settlePool, (samePid)));

        // 推进时间到结束时间后并 finish
        vm.warp(newEndTime + 1);
        _executeViaSafe(address(pool), abi.encodeCall(pool.finishPool, (samePid)));

        // 验证状态为 FINISH（2）
        assertEq(pool.getPoolState(samePid), 2);

        // 验证 finishAmountLend > 本金（有利息）
        assertTrue(_finishAmountLend(samePid) > 1000e18);
    }

    // ============================================================
    // 9. finishPool — different tokens
    // 测试 lendToken 和 borrowToken 为不同代币时的 finish 流程
    // ============================================================

    /**
     * @notice 在结束时间之前调用 finishPool 应被拒绝
     */
    function test_finishPool_revertsBeforeEndTime() public {
        _setFee(1e7, 5e6);
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 当前时间在 endTime 之前
        vm.prank(safeAddress);
        vm.expectRevert("Not reached end time");
        pool.finishPool(pid);
    }

    /**
     * @notice 不同币种池子的 finish 操作（涉及 DEX 兑换）
     */
    function test_finishPool_differentTokens() public {
        _setFee(1e7, 5e6);
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 推进时间到结束时间后
        _finishPool();

        // 验证状态为 FINISH（2）
        assertEq(pool.getPoolState(pid), 2);

        // 验证 finishAmountLend > 本金（有利息收益）
        assertTrue(_finishAmountLend(pid) > 1000e18);
        // 验证 finishAmountBorrow < 总质押（部分用于兑换和手续费）
        assertTrue(_finishAmountBorrow(pid) < 2000e18);
    }

    // ============================================================
    // 10. 销毁凭证取回资产
    // 测试销毁债务代币并取回对应资产
    // ============================================================

    /**
     * @notice 出借人销毁 SP Token 取回本金 + 利息
     */
    function test_destroyTokens_lendReceivesPrincipalAndInterest() public {
        _setFee(1e7, 5e6);
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 出借人领取 SP Token
        vm.prank(lender);
        pool.claimLendDebtToken(pid);

        // 借款人领取 JP Token
        vm.prank(borrower);
        pool.claimBorrow(pid);

        // finish
        _finishPool();

        // 记录销毁前的 lender 余额
        uint256 balanceBefore = tokenA.balanceOf(lender);
        uint256 spBalance = lendDebt.balanceOf(lender);

        // 销毁所有 SP Token
        vm.prank(lender);
        pool.destroyLendDebtToken(pid, spBalance);

        // 验证收到了本金+利息（余额增加）
        assertTrue(tokenA.balanceOf(lender) > balanceBefore);
        // 验证 SP Token 已全部销毁
        assertEq(lendDebt.balanceOf(lender), 0);
    }

    /**
     * @notice 借款人销毁 JP Token 取回剩余抵押物
     */
    function test_destroyTokens_borrowRecoversCollateral() public {
        _setFee(1e7, 5e6);
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 领取凭证
        vm.prank(lender);
        pool.claimLendDebtToken(pid);
        vm.prank(borrower);
        pool.claimBorrow(pid);

        // finish
        _finishPool();

        // 销毁所有 JP Token
        uint256 jpBalance = borrowDebt.balanceOf(borrower);
        vm.prank(borrower);
        pool.destroyBorrowDebtToken(pid, jpBalance);

        // 验证 JP Token 已全部销毁
        assertEq(borrowDebt.balanceOf(borrower), 0);
    }

    // ============================================================
    // 11. 清算
    // 测试资产价格下跌时的清算流程
    // ============================================================

    /**
     * @notice checkCanLiquidate 在抵押品价值下跌时返回 true
     */
    function test_liquidation_checkReturnsTrue() public {
        _setFee(1e7, 5e6);
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 通过预言机将 tokenB 价格压低到 0.05 USD，触发清算条件
        oracle.setPrice(address(tokenB), 5e16);

        assertTrue(pool.checkCanLiquidate(pid));
    }

    /**
     * @notice liquidatePool 将池子状态变为 LIQUIDATION
     */
    function test_liquidation_setsLiquidationState() public {
        _setFee(1e7, 5e6);
        _depositAndBorrow(1000e18, 2000e18);
        _settlePool();

        // 压低 tokenB 价格触发清算
        oracle.setPrice(address(tokenB), 5e16);

        vm.expectEmit(true, false, false, false);
        emit LiquidatePool(pid, 0, 0); // 仅验证主题，不检查具体金额
        vm.prank(liquidator);
        pool.liquidatePool(pid);

        // 验证状态为 LIQUIDATION（3）
        assertEq(pool.getPoolState(pid), 3);

        // 验证 liquidationAmountLend > 0
        assertTrue(_liquidationAmountLend(pid) > 0);
    }

    // ============================================================
    // 12. 紧急提取
    // 测试池子处于 UNDONE 状态时的紧急提取功能
    // ============================================================

    /**
     * @notice 出借人在 UNDONE 状态下的紧急提取
     */
    function test_emergencyWithdraw_lendInUndone() public {
        // 出借人存入，但借款人未参与
        vm.prank(lender);
        pool.lend(pid, 1000e18);

        // 结算后池子进入 UNDONE（纯出借、无质押）
        _settlePool();
        assertEq(pool.getPoolState(pid), 4); // UNDONE

        // 紧急提取
        vm.prank(lender);
        pool.emergencyWithdrawLend(pid);

        // 验证出借人收回了全部 tokenA（初始 500,000 - 存入 1,000 + 取回 1,000 = 500,000）
        assertEq(tokenA.balanceOf(lender), 500_000e18);
    }

    /**
     * @notice 借款人在 UNDONE 状态下的紧急提取
     */
    function test_emergencyWithdraw_borrowInUndone() public {
        // 借款人质押，但出借人未参与
        vm.prank(borrower);
        pool.borrow(pid, 2000e18);

        // 结算后池子进入 UNDONE（纯质押、无出借）
        _settlePool();
        assertEq(pool.getPoolState(pid), 4); // UNDONE

        // 紧急提取
        vm.prank(borrower);
        pool.emergencyWithdrawBorrow(pid);

        // 验证借款人收回了全部 tokenB
        assertEq(tokenB.balanceOf(borrower), 500_000e18);
    }

    /**
     * @notice 非 UNDONE 状态下调用紧急提取应被拒绝
     */
    function test_emergencyWithdraw_revertsInNonUndone() public {
        // 池子处于 MATCH（0）状态
        vm.prank(lender);
        pool.lend(pid, 1000e18);

        vm.prank(lender);
        vm.expectRevert("Pool state must be UNDONE");
        pool.emergencyWithdrawLend(pid);
    }

    // ============================================================
    // 13. 完整生命周期
    // 测试从存放到销毁的完整流程
    // ============================================================

    /**
     * @notice 完整借贷生命周期：lend → borrow → settle → claim → finish → destroy
     */
    function test_fullLifecycle() public {
        // 设置手续费
        _setFee(1e7, 5e6);

        // --- 阶段 1：存入 ---
        vm.prank(lender);
        pool.lend(pid, 1000e18);
        vm.prank(lender2);
        pool.lend(pid, 500e18);
        vm.prank(borrower);
        pool.borrow(pid, 2000e18);
        vm.prank(borrower2);
        pool.borrow(pid, 1000e18);

        // 验证总存入和总质押量
        assertEq(_lendSupply(pid), 1500e18);
        assertEq(_borrowSupply(pid), 3000e18);

        // --- 阶段 2：结算 ---
        _settlePool();
        assertEq(pool.getPoolState(pid), 1); // EXECUTION

        // --- 阶段 3：领取凭证 ---
        vm.prank(lender);
        pool.claimLendDebtToken(pid);
        vm.prank(lender2);
        pool.claimLendDebtToken(pid);
        vm.prank(borrower);
        pool.claimBorrow(pid);
        vm.prank(borrower2);
        pool.claimBorrow(pid);

        // --- 阶段 4：finish ---
        _finishPool();
        assertEq(pool.getPoolState(pid), 2); // FINISH

        // --- 阶段 5：销毁凭证取回资产 ---
        // 先在 vm.prank 之前读取余额，避免 prank 被静态调用消耗
        uint256 lenderSpBalance = lendDebt.balanceOf(lender);
        uint256 lender2SpBalance = lendDebt.balanceOf(lender2);
        uint256 borrowerJpBalance = borrowDebt.balanceOf(borrower);
        uint256 borrower2JpBalance = borrowDebt.balanceOf(borrower2);

        vm.prank(lender);
        pool.destroyLendDebtToken(pid, lenderSpBalance);
        vm.prank(lender2);
        pool.destroyLendDebtToken(pid, lender2SpBalance);
        vm.prank(borrower);
        pool.destroyBorrowDebtToken(pid, borrowerJpBalance);
        vm.prank(borrower2);
        pool.destroyBorrowDebtToken(pid, borrower2JpBalance);

        // 所有债务代币已销毁
        assertEq(lendDebt.totalSupply(), 0);
        assertEq(borrowDebt.totalSupply(), 0);
        // 手续费已被收取
        assertTrue(tokenA.balanceOf(feeCollector) > 0);
    }

    // ============================================================
    // 14. 校验函数 & 边界
    // 测试各种校验函数和边界条件
    // ============================================================

    /**
     * @notice checkCanSettle 在结算时间之前返回 false
     */
    function test_checkFunctions_checkCanSettleFalseBeforeSettle() public {
        _depositAndBorrow(1000e18, 2000e18);
        // 当前时间在结算时间之前
        assertFalse(pool.checkCanSettle(pid));
    }

    /**
     * @notice checkCanFinish 在 MATCH 状态下返回 false
     */
    function test_checkFunctions_checkCanFinishFalseForMatch() public {
        // 池子处于 MATCH 状态（未结算）
        assertFalse(pool.checkCanFinish(pid));
    }

    /**
     * @notice 对不存在池子 ID 的操作触发 panic(0x32)
     */
    function test_checkFunctions_revertsNonExistentPool() public {
        vm.prank(safeAddress);
        vm.expectRevert(abi.encodeWithSignature("Panic(uint256)", 0x32));
        pool.settlePool(999);
    }

    /**
     * @notice onlyOwner 函数对非 owner 地址都正确拒绝
     */
    function test_checkFunctions_onlyOwnerRevertsForNonOwner() public {
        // setOracle
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        pool.setOracle(random);

        // setSwapRouter
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        pool.setSwapRouter(random);

        // setGlobalPaused
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        pool.setGlobalPaused();

        // settlePool
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        pool.settlePool(pid);

        // finishPool
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        pool.finishPool(pid);
    }

    /**
     * @notice lend 传入 0 金额被拒绝
     */
    function test_checkFunctions_lendZeroReverts() public {
        vm.prank(lender);
        vm.expectRevert("ERC20 must be greater than 0");
        pool.lend(pid, 0);
    }

    /**
     * @notice 全局暂停时所有操作被拒绝
     */
    function test_checkFunctions_revertsWhenPaused() public {
        // 通过 Safe 开启全局暂停
        _executeViaSafe(address(pool), abi.encodeCall(pool.setGlobalPaused, ()));
        assertTrue(pool.globalPaused());

        // 暂停后调用 lend 应被拒绝
        vm.prank(lender);
        vm.expectRevert("Global Paused");
        pool.lend(pid, 1000e18);
    }

    // ============================================================
    // 15. 退款测试
    // 验证超募退款功能
    // ============================================================

    /**
     * @notice 出借人领取超募退款
     */
    function test_refund_lendRefund() public {
        // 出借人存入 2000 tokenA，借款人质押 1000 tokenB
        // 按 1:1 价格和 200% 抵押率，settleAmountLend = 1000 / 40%
        // 实际：抵押物价值可支持最大出借量 = 1000 * 1e18 * 1 * 1e8 / (1 * 2e8) = 500e18
        // 超募部分 = 2000 - 500 = 1500
        // 所以需要借入较少的 borrow 使 matchedLend 小于 lendSupply
        vm.startPrank(lender);
        pool.lend(pid, 2000e18);
        vm.stopPrank();

        vm.startPrank(borrower);
        pool.borrow(pid, 1000e18);
        vm.stopPrank();

        _settlePool();

        // 出借人领取退款（10% 的超募比例）
        vm.prank(lender);
        pool.refundLend(pid);
    }

    /**
     * @notice 借款人领取多余质押物退款
     */
    function test_refund_borrowRefund() public {
        // 出借人存入 500 tokenA，借款人质押 2000 tokenB
        // settleAmountLend = min(500, 2000*1e8/2e8) = min(500, 1000) = 500
        // settleAmountBorrow = 500*2e8/1e8 = 1000
        // 退款 = 2000 - 1000 = 1000
        vm.startPrank(lender);
        pool.lend(pid, 500e18);
        vm.stopPrank();

        vm.startPrank(borrower);
        pool.borrow(pid, 2000e18);
        vm.stopPrank();

        _settlePool();

        // 借款人领取多余质押物退款
        vm.prank(borrower);
        pool.refundBorrow(pid);
    }

    // ============================================================
    // 16. Safe 多签安全测试
    // 验证单签名无法通过 Safe 执行管理员操作
    // ============================================================

    /**
     * @notice 测试单签名无法执行管理操作（确保多签阈值 2 生效）
     */
    function test_safe_revertsWithSingleSignature() public {
        // 构造 setFee 调用数据
        bytes memory data = abi.encodeCall(pool.setFee, (1e7, 5e6));
        // 获取 Safe 交易哈希
        bytes32 txHash = safe.getTransactionHash(
            address(pool), 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), safe.nonce()
        );
        // 仅使用第一个 owner 签名（不满足阈值 2）
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(_ownerPrivateKey1, txHash);
        bytes memory sig = abi.encodePacked(r, s, v);

        // 预期 Safe 执行失败
        vm.expectRevert();
        safe.execTransaction(
            address(pool), 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), sig
        );

        // 手续费应保持默认值 0，未被修改
        assertEq(pool.lendFee(), 0);
    }
}
