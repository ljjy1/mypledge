// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

/**
 * @title IUniswapV2Factory 接口
 * @notice 定义工厂合约的核心功能
 * @dev 工厂合约负责创建和管理交易对
 *      是 UniswapV2 协议的核心组件之一
 */
interface IUniswapV2Factory {
    /**
     * @notice 获取费用接收地址
     * @return 费用接收地址
     * @dev 协议手续费会发送到此地址
     */
    function feeTo() external view returns (address);

    /**
     * @notice 设置费用接收地址
     * @param 新的费用接收地址
     * @dev 只有管理员可以调用
     */
    function setFeeTo(address) external;

    /**
     * @notice 获取管理员地址
     * @return 管理员地址
     */
    function feeToSetter() external view returns (address);

    /**
     * @notice 设置新管理员
     * @param 新管理员地址
     * @dev 只有当前管理员可以调用
     */
    function setFeeToSetter(address) external;

    /**
     * @notice 获取交易对地址
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @return pair 交易对地址（如果不存在则返回零地址）
     */
    function getPair(address tokenA, address tokenB) external view returns (address pair);

    /**
     * @notice 获取所有交易对
     * @param index 交易对索引
     * @return pair 交易对地址
     */
    function allPairs(uint256 index) external view returns (address pair);

    /**
     * @notice 获取交易对数量
     * @return 交易对数量
     */
    function allPairsLength() external view returns (uint256);

    /**
     * @notice 创建交易对
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @return pair 交易对地址
     * @dev 使用 CREATE2 创建交易对，确保地址可预测
     */
    function createPair(address tokenA, address tokenB) external returns (address pair);

    /**
     * @notice 获取 Pair 合约的 init code hash
     * @return hash init code 的 keccak256 哈希
     * @dev 用于 CREATE2 计算交易对地址
     */
    function initCodeHash() external view returns (bytes32);
}
