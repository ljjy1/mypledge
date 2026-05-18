// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

/**
 * @title IUniswapV2Callee 接口
 * @notice 闪电贷回调接口
 * @dev 用户合约需要实现此接口以接收闪电贷回调
 *      闪电贷允许用户在单笔交易中借款并还款，无需抵押
 *      这是 DeFi 中的一种强大金融工具
 */
interface IUniswapV2Callee {
    /**
     * @notice 闪电贷回调函数
     * @param sender 发起者地址（调用 swap 函数的地址）
     * @param amount0 借出的代币0数量
     * @param amount1 借出的代币1数量
     * @param data 用户自定义数据
     * @dev 当用户在 swap 中传入非空的 data 时触发
     *      用户合约必须在此回调中归还借款 + 手续费
     *      如果归还失败，整个交易会回滚
     * 
     * 使用示例：
     * 1. 用户调用 pair.swap(amount0Out, amount1Out, userContract, data)
     * 2. Pair 合约将代币转给 userContract
     * 3. Pair 合约调用 userContract.uniswapV2Call(...)
     * 4. userContract 在回调中使用借来的代币（如套利）
     * 5. userContract 归还借款 + 0.3% 手续费
     * 6. 交易完成
     */
    function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes calldata data) external;
}
