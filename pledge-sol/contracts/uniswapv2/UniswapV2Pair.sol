// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

import "@openzeppelin/contracts/token/ERC20/ERC20.sol"; // 导入 OpenZeppelin 的 ERC20 实现
import "@openzeppelin/contracts/token/ERC20/IERC20.sol"; // 导入 OpenZeppelin 的 ERC20 接口
import "./libraries/Math.sol"; // 导入数学库
import "./libraries/UQ112x112.sol"; // 导入定点数库
import "./interfaces/IUniswapV2Pair.sol"; // 导入交易对接口
import "./interfaces/IUniswapV2Callee.sol"; // 导入闪电贷回调接口
import "./interfaces/IUniswapV2Factory.sol"; // 导入工厂接口

/**
 * @title UniswapV2Pair 交易对合约
 * @notice 实现代币交易对的核心逻辑
 * @dev 包含流动性管理、交易执行、价格预言机等功能
 *      这是 UniswapV2 的核心合约，实现了 AMM（自动做市商）机制
 *      使用恒定乘积公式 x * y = k 来确定交易价格
 */
contract UniswapV2Pair is ERC20, IUniswapV2Pair {
    using UQ112x112 for uint224; // 为 uint224 类型启用 UQ112x112 库的方法

    // 最小流动性，永久锁定
    // 首次添加流动性时，会铸造 1000 个 LP 代币并锁定
    // 这是为了防止流动性池被清空，确保始终有流动性
    uint256 public constant MINIMUM_LIQUIDITY = 10 ** 3;

    // 选择器常量，用于识别 ERC20 的 transfer 函数
    // 通过函数签名计算得到：keccak256("transfer(address,uint256)") 的前 4 字节
    bytes4 private constant SELECTOR = bytes4(keccak256(bytes("transfer(address,uint256)")));

    // 工厂合约地址：创建此交易对的工厂合约
    address public factory;
    // 代币0地址：交易对中的第一个代币（地址较小的）
    address public token0;
    // 代币1地址：交易对中的第二个代币（地址较大的）
    address public token1;

    // 代币0储备量：当前池中代币0的数量
    // 使用 uint112 节省 gas，同时足够大（约 5 * 10^33）
    uint112 private reserve0;
    // 代币1储备量：当前池中代币1的数量
    uint112 private reserve1;
    // 最后更新时间戳：上次储备量更新的区块时间
    // 使用 uint32 节省存储空间，可以表示约 136 年
    uint32 private blockTimestampLast;

    // 价格累积器0：token1/token0 的累积价格（用于 TWAP）
    // TWAP = Time Weighted Average Price（时间加权平均价格）
    uint256 public price0CumulativeLast;
    // 价格累积器1：token0/token1 的累积价格（用于 TWAP）
    uint256 public price1CumulativeLast;
    // 上次铸造时的 K 值：储备量乘积 reserve0 * reserve1
    // 用于计算协议手续费
    uint256 public kLast;

    // 铸造事件：当添加流动性时触发
    event Mint(address indexed sender, uint256 amount0, uint256 amount1);
    // 销毁事件：当移除流动性时触发
    event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to);
    // 交换事件：当执行交易时触发
    event Swap(
        address indexed sender, // 发起者
        uint256 amount0In, // 输入的代币0数量
        uint256 amount1In, // 输入的代币1数量
        uint256 amount0Out, // 输出的代币0数量
        uint256 amount1Out, // 输出的代币1数量
        address indexed to // 接收者
    );
    // 同步事件：当储备量更新时触发
    event Sync(uint112 reserve0, uint112 reserve1);

    /**
     * @notice 构造函数
     * @dev 初始化 LP 代币名称和符号，设置工厂地址
     *      LP 代币名称为 "Uniswap V2"，符号为 "UNI-V2"
     */
    constructor() ERC20("Uniswap V2", "UNI-V2") {
        factory = msg.sender; // 记录创建者（工厂合约）地址
    }

    /**
     * @notice 初始化函数，仅由工厂合约调用一次
     * @param _token0 代币0地址
     * @param _token1 代币1地址
     * @dev 在工厂合约创建交易对后调用，设置两个代币地址
     */
    function initialize(address _token0, address _token1) external {
        // 只有工厂合约可以调用
        require(msg.sender == factory, "UniswapV2: FORBIDDEN");
        token0 = _token0; // 设置代币0地址
        token1 = _token1; // 设置代币1地址
    }

    /**
     * @notice 获取储备量
     * @return _reserve0 代币0储备量
     * @return _reserve1 代币1储备量
     * @return _blockTimestampLast 最后更新时间戳
     * @dev 返回当前的储备量和最后更新时间
     */
    function getReserves() public view returns (uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast) {
        _reserve0 = reserve0; // 返回代币0储备量
        _reserve1 = reserve1; // 返回代币1储备量
        _blockTimestampLast = blockTimestampLast; // 返回最后更新时间戳
    }

    /**
     * @notice 安全转账函数
     * @param token 代币地址
     * @param to 接收地址
     * @param amount 转账金额
     * @dev 使用低级 call 调用代币的 transfer 函数
     *      检查返回值确保转账成功
     */
    function _safeTransfer(address token, address to, uint256 amount) private {
        // 调用代币的 transfer 函数
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(SELECTOR, to, amount));
        // 检查调用是否成功，以及返回值是否为 true（或没有返回值）
        require(success && (data.length == 0 || abi.decode(data, (bool))), "UniswapV2: TRANSFER_FAILED");
    }

    /**
     * @notice 更新储备量和价格累积器
     * @param balance0 代币0当前余额
     * @param balance1 代币1当前余额
     * @param _reserve0 当前储备量0
     * @param _reserve1 当前储备量1
     * @dev 在每次储备量变化时调用，更新价格预言机数据
     */
    function _update(uint256 balance0, uint256 balance1, uint112 _reserve0, uint112 _reserve1) private {
        // 确保余额不超过 uint112 最大值
        require(balance0 <= type(uint112).max && balance1 <= type(uint112).max, "UniswapV2: OVERFLOW");
        // 获取当前区块时间戳（取模 2^32 防止溢出）
        uint32 blockTimestamp = uint32(block.timestamp % 2 ** 32);
        uint32 timeElapsed;
        // 计算时间差（使用 unchecked 因为时间戳可能溢出）
        unchecked {
            timeElapsed = blockTimestamp - blockTimestampLast;
        }
        // 如果时间差大于 0 且储备量不为 0，更新价格累积器
        if (timeElapsed > 0 && _reserve0 != 0 && _reserve1 != 0) {
            // 更新价格累积器0：price0 = reserve1 / reserve0
            // 使用定点数运算，然后乘以时间差
            price0CumulativeLast += uint256(UQ112x112.encode(_reserve1).uqdiv(_reserve0)) * timeElapsed;
            // 更新价格累积器1：price1 = reserve0 / reserve1
            price1CumulativeLast += uint256(UQ112x112.encode(_reserve0).uqdiv(_reserve1)) * timeElapsed;
        }
        // 更新储备量
        reserve0 = uint112(balance0);
        reserve1 = uint112(balance1);
        // 更新最后更新时间戳
        blockTimestampLast = blockTimestamp;
        // 触发同步事件
        emit Sync(reserve0, reserve1);
    }

    /**
     * @notice 铸造协议费用（如果设置了费用开关）
     * @param _reserve0 当前储备量0
     * @param _reserve1 当前储备量1
     * @return feeOn 是否开启费用
     * @dev 如果 feeTo 设置了地址，则收取交易手续费的 1/6 作为协议费
     *      协议费通过铸造 LP 代币给 feeTo 地址实现
     */
    function _mintFee(uint112 _reserve0, uint112 _reserve1) private returns (bool feeOn) {
        // 获取费用接收地址
        address feeTo = IUniswapV2Factory(factory).feeTo();
        // 判断是否开启费用
        feeOn = feeTo != address(0);
        uint256 _kLast = kLast; // 读取上次 K 值
        if (feeOn) {
            // 如果开启了费用
            if (_kLast != 0) {
                // 计算当前 K 值的平方根
                uint256 rootK = Math.sqrt(uint256(_reserve0) * uint256(_reserve1));
                // 计算上次 K 值的平方根
                uint256 rootKLast = Math.sqrt(_kLast);
                // 如果 K 值增加了（有交易手续费累积）
                if (rootK > rootKLast) {
                    // 计算协议费用对应的 LP 代币数量
                    // 公式：liquidity = totalSupply * (rootK - rootKLast) / (rootK * 5 + rootKLast)
                    // 这相当于收取增长部分的 1/6
                    uint256 numerator = totalSupply() * (rootK - rootKLast);
                    uint256 denominator = rootK * 5 + rootKLast;
                    uint256 liquidity = numerator / denominator;
                    // 如果有流动性，铸造给 feeTo 地址
                    if (liquidity > 0) _mint(feeTo, liquidity);
                }
            }
        } else if (_kLast != 0) {
            // 如果关闭了费用，清空 kLast
            kLast = 0;
        }
    }

    /**
     * @notice 铸造 LP 代币
     * @param to 接收地址
     * @return liquidity 铸造的流动性数量
     * @dev 当用户向交易对添加流动性时调用
     *      用户需要先将代币转入交易对合约，然后调用此函数
     */
    function mint(address to) external returns (uint256 liquidity) {
        // 获取当前储备量
        (uint112 _reserve0, uint112 _reserve1,) = getReserves();
        // 获取当前代币余额
        uint256 balance0 = IERC20(token0).balanceOf(address(this));
        uint256 balance1 = IERC20(token1).balanceOf(address(this));
        // 计算新增的代币数量
        uint256 amount0 = balance0 - _reserve0;
        uint256 amount1 = balance1 - _reserve1;

        // 尝试铸造协议费用
        bool feeOn = _mintFee(_reserve0, _reserve1);
        uint256 _totalSupply = totalSupply();
        if (_totalSupply == 0) {
            // 首次添加流动性
            // 流动性 = sqrt(amount0 * amount1) - MINIMUM_LIQUIDITY
            liquidity = Math.sqrt(amount0 * amount1) - MINIMUM_LIQUIDITY;
            // 永久锁定最小流动性到 0xdead 地址
            // 这是为了防止流动性池被清空
            _mint(address(0xdead), MINIMUM_LIQUIDITY);
        } else {
            // 后续添加流动性
            // 流动性按比例计算，取较小值
            // liquidity = min(amount0 * totalSupply / reserve0, amount1 * totalSupply / reserve1)
            liquidity = Math.min(amount0 * _totalSupply / _reserve0, amount1 * _totalSupply / _reserve1);
        }
        // 确保铸造的流动性大于 0
        require(liquidity > 0, "UniswapV2: INSUFFICIENT_LIQUIDITY_MINTED");
        // 铸造 LP 代币给接收地址
        _mint(to, liquidity);

        // 更新储备量
        _update(balance0, balance1, _reserve0, _reserve1);
        // 如果开启了费用，更新 kLast
        if (feeOn) kLast = uint256(reserve0) * uint256(reserve1);
        // 触发铸造事件
        emit Mint(msg.sender, amount0, amount1);
    }

    /**
     * @notice 销毁 LP 代币并提取代币
     * @param to 接收地址
     * @return amount0 提取的代币0数量
     * @return amount1 提取的代币1数量
     * @dev 当用户移除流动性时调用
     *      用户需要先将 LP 代币转入交易对合约，然后调用此函数
     */
    function burn(address to) external returns (uint256 amount0, uint256 amount1) {
        // 获取当前储备量
        (uint112 _reserve0, uint112 _reserve1,) = getReserves();
        // 获取代币地址（本地变量节省 gas）
        address _token0 = token0;
        address _token1 = token1;
        // 获取当前代币余额
        uint256 balance0 = IERC20(_token0).balanceOf(address(this));
        uint256 balance1 = IERC20(_token1).balanceOf(address(this));
        // 获取合约中的 LP 代币数量（用户需要先转入）
        uint256 liquidity = balanceOf(address(this));

        // 尝试铸造协议费用
        bool feeOn = _mintFee(_reserve0, _reserve1);
        uint256 _totalSupply = totalSupply();
        // 按比例计算提取的代币数量
        // amount = liquidity * balance / totalSupply
        amount0 = liquidity * balance0 / _totalSupply;
        amount1 = liquidity * balance1 / _totalSupply;
        // 确保提取数量大于 0
        require(amount0 > 0 && amount1 > 0, "UniswapV2: INSUFFICIENT_LIQUIDITY_BURNED");
        // 销毁 LP 代币
        _burn(address(this), liquidity);
        // 转移代币给接收地址
        _safeTransfer(_token0, to, amount0);
        _safeTransfer(_token1, to, amount1);
        // 获取更新后的余额
        balance0 = IERC20(_token0).balanceOf(address(this));
        balance1 = IERC20(_token1).balanceOf(address(this));

        // 更新储备量
        _update(balance0, balance1, _reserve0, _reserve1);
        // 如果开启了费用，更新 kLast
        if (feeOn) kLast = uint256(reserve0) * uint256(reserve1);
        // 触发销毁事件
        emit Burn(msg.sender, amount0, amount1, to);
    }

    /**
     * @notice 代币交换
     * @param amount0Out 输出的代币0数量
     * @param amount1Out 输出的代币1数量
     * @param to 接收地址
     * @param data 回调数据（用于闪电贷，如果非空则触发闪电贷回调）
     * @dev 执行代币交换，支持普通交换和闪电贷
     *      闪电贷：先借出代币，在回调中归还，无需抵押
     */
    function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata data) external {
        // 确保至少有一个输出
        require(amount0Out > 0 || amount1Out > 0, "UniswapV2: INSUFFICIENT_OUTPUT_AMOUNT");
        // 获取当前储备量
        (uint112 _reserve0, uint112 _reserve1,) = getReserves();
        // 确保输出不超过储备量
        require(amount0Out < _reserve0 && amount1Out < _reserve1, "UniswapV2: INSUFFICIENT_LIQUIDITY");

        uint256 balance0; // 代币0余额
        uint256 balance1; // 代币1余额
        {
            // 使用局部作用域节省 gas
            address _token0 = token0;
            address _token1 = token1;
            // 确保接收地址不是代币合约地址
            require(to != _token0 && to != _token1, "UniswapV2: INVALID_TO");
            // 转移输出的代币0
            if (amount0Out > 0) _safeTransfer(_token0, to, amount0Out);
            // 转移输出的代币1
            if (amount1Out > 0) _safeTransfer(_token1, to, amount1Out);
            // 如果有回调数据，执行闪电贷回调
            if (data.length > 0) IUniswapV2Callee(to).uniswapV2Call(msg.sender, amount0Out, amount1Out, data);
            // 获取当前余额
            balance0 = IERC20(_token0).balanceOf(address(this));
            balance1 = IERC20(_token1).balanceOf(address(this));
        }
        // 计算输入的代币数量
        // 如果余额 > (储备量 - 输出量)，说明有代币输入
        uint256 amount0In = balance0 > _reserve0 - amount0Out ? balance0 - (_reserve0 - amount0Out) : 0;
        uint256 amount1In = balance1 > _reserve1 - amount1Out ? balance1 - (_reserve1 - amount1Out) : 0;
        // 确保至少有一个输入
        require(amount0In > 0 || amount1In > 0, "UniswapV2: INSUFFICIENT_INPUT_AMOUNT");
        {
            // 验证 K 值（扣除 0.3% 手续费后）
            // 调整后的余额 = (余额 * 1000 - 输入量 * 3)
            // 这相当于扣除 0.3% 的手续费
            uint256 balance0Adjusted = balance0 * 1000 - amount0In * 3;
            uint256 balance1Adjusted = balance1 * 1000 - amount1In * 3;
            // 确保 K 值不减少：balance0Adjusted * balance1Adjusted >= reserve0 * reserve1 * 1000^2
            require(balance0Adjusted * balance1Adjusted >= uint256(_reserve0) * uint256(_reserve1) * 1000 ** 2, "UniswapV2: K");
        }

        // 更新储备量
        _update(balance0, balance1, _reserve0, _reserve1);
        // 触发交换事件
        emit Swap(msg.sender, amount0In, amount1In, amount0Out, amount1Out, to);
    }

    /**
     * @notice 强制同步储备量
     * @param to 接收地址
     * @dev 用于处理直接转账导致的不平衡
     *      将多余的代币发送给指定地址
     *      这可以防止有人通过直接转账操纵价格
     */
    function skim(address to) external {
        address _token0 = token0;
        address _token1 = token1;
        // 将余额与储备量的差额发送给指定地址
        _safeTransfer(_token0, to, IERC20(_token0).balanceOf(address(this)) - reserve0);
        _safeTransfer(_token1, to, IERC20(_token1).balanceOf(address(this)) - reserve1);
    }

    /**
     * @notice 同步储备量到当前余额
     * @dev 更新储备量以匹配当前代币余额
     *      当有人直接向交易对转账时，需要调用此函数更新储备量
     */
    function sync() external {
        _update(IERC20(token0).balanceOf(address(this)), IERC20(token1).balanceOf(address(this)), reserve0, reserve1);
    }
}
