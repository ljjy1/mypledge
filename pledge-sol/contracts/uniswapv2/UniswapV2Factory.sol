// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

import "./UniswapV2Pair.sol"; // 导入交易对合约
import "./interfaces/IUniswapV2Factory.sol"; // 导入工厂接口

/**
 * @title UniswapV2Factory 工厂合约
 * @notice 创建和管理交易对
 * @dev 使用 CREATE2 创建交易对，确保地址可预测
 *      工厂合约是 UniswapV2 的核心，负责创建和管理所有交易对
 *      每个交易对由两个代币组成，且每个代币对只有一个交易对
 */
contract UniswapV2Factory is IUniswapV2Factory {
    // 费用接收地址：协议手续费将发送到此地址
    // 如果设置为 address(0)，则不收取协议手续费
    address public feeTo;
    // 管理员地址：有权设置 feeTo 和转移管理员权限
    address public feeToSetter;

    // 交易对映射：token0 => token1 => pair
    // 用于快速查找两个代币对应的交易对地址
    mapping(address => mapping(address => address)) public getPair;
    // 所有交易对列表：存储所有已创建的交易对地址
    address[] public allPairs;
    // Pair 合约 init code hash，用于 CREATE2 地址计算
    bytes32 public initCodeHash;

    // 交易对创建事件：当新交易对创建时触发
    event PairCreated(address indexed token0, address indexed token1, address pair, uint256);

    /**
     * @notice 构造函数
     * @param _feeToSetter 管理员地址
     * @dev 初始化管理员地址，feeTo 默认为 address(0)
     */
    constructor(address _feeToSetter) {
        feeToSetter = _feeToSetter; // 设置管理员地址
        initCodeHash = keccak256(type(UniswapV2Pair).creationCode);
    }

    /**
     * @notice 获取所有交易对数量
     * @return 交易对数量
     * @dev 返回 allPairs 数组的长度
     */
    function allPairsLength() external view returns (uint256) {
        return allPairs.length; // 返回交易对数组长度
    }

    /**
     * @notice 创建交易对
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @return pair 交易对地址
     * @dev 使用 CREATE2 创建交易对，确保地址可预测
     *      CREATE2 允许在部署前就知道合约地址
     *      这对于计算交易对地址非常重要
     */
    function createPair(address tokenA, address tokenB) external returns (address pair) {
        // 确保两个代币地址不相同
        require(tokenA != tokenB, "UniswapV2: IDENTICAL_ADDRESSES");
        // 对代币地址排序：token0 是较小的地址，token1 是较大的地址
        // 排序确保无论用户以何种顺序提供代币，都能找到相同的交易对
        (address token0, address token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        // 确保 token0 不是零地址
        require(token0 != address(0), "UniswapV2: ZERO_ADDRESS");
        // 确保交易对不存在
        require(getPair[token0][token1] == address(0), "UniswapV2: PAIR_EXISTS");

        // 使用 CREATE2 创建交易对
        // 获取交易对合约的创建字节码
        bytes memory bytecode = type(UniswapV2Pair).creationCode;
        // 计算 salt：两个代币地址的哈希
        // salt 用于 CREATE2 地址计算，确保相同代币对产生相同地址
        bytes32 salt = keccak256(abi.encodePacked(token0, token1));
        // 使用内联汇编调用 CREATE2 操作码
        // CREATE2 语法：create2(value, offset, length, salt)
        // value: 发送的 ETH 数量（这里为 0）
        // offset: 字节码在内存中的起始位置
        // length: 字节码长度
        // salt: 用于地址计算的盐值
        assembly {
            pair := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }
        // 初始化交易对：设置两个代币地址
        UniswapV2Pair(pair).initialize(token0, token1);

        // 记录交易对：双向映射，方便从任一代币查找交易对
        getPair[token0][token1] = pair;
        getPair[token1][token0] = pair;
        // 将交易对添加到列表
        allPairs.push(pair);

        // 触发交易对创建事件
        emit PairCreated(token0, token1, pair, allPairs.length);
    }

    /**
     * @notice 设置费用接收地址
     * @param _feeTo 费用接收地址
     * @dev 只有管理员可以调用
     *      设置为 address(0) 表示不收取协议手续费
     *      设置为有效地址表示收取协议手续费（交易手续费的 1/6）
     */
    function setFeeTo(address _feeTo) external {
        // 验证调用者是管理员
        require(msg.sender == feeToSetter, "UniswapV2: FORBIDDEN");
        // 设置费用接收地址
        feeTo = _feeTo;
    }

    /**
     * @notice 设置管理员地址
     * @param _feeToSetter 新管理员地址
     * @dev 只有当前管理员可以调用
     *      用于转移管理员权限
     */
    function setFeeToSetter(address _feeToSetter) external {
        // 验证调用者是管理员
        require(msg.sender == feeToSetter, "UniswapV2: FORBIDDEN");
        // 设置新管理员地址
        feeToSetter = _feeToSetter;
    }
}
