// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

/**
 * @title Math 库
 * @notice 提供数学运算辅助函数
 * @dev 这是一个库合约，提供常用的数学运算功能
 */
library Math {
    /**
     * @notice 返回两个数中的最大值
     * @param a 第一个数
     * @param b 第二个数
     * @return 较大的数
     */
    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a >= b ? a : b; // 如果 a 大于等于 b，返回 a，否则返回 b
    }

    /**
     * @notice 返回两个数中的最小值
     * @param a 第一个数
     * @param b 第二个数
     * @return 较小的数
     */
    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b; // 如果 a 小于 b，返回 a，否则返回 b
    }

    /**
     * @notice 计算平方根（Babylonian 方法）
     * @param y 输入值
     * @return z 平方根结果
     * @dev 使用 Babylonian 方法迭代计算平方根，这是一种古老的数值计算方法
     *      通过不断逼近的方式计算平方根，精度足够用于智能合约
     */
    function sqrt(uint256 y) internal pure returns (uint256 z) {
        if (y > 3) { // 如果 y 大于 3，使用迭代方法计算平方根
            z = y; // 初始化 z 为 y
            uint256 x = y / 2 + 1; // 初始化 x 为 y/2 + 1，这是平方根的初始估计值
            while (x < z) { // 当 x 小于 z 时继续迭代，逐步逼近真实的平方根
                z = x; // 将 x 的值赋给 z，保存当前估计值
                x = (y / x + x) / 2; // 使用 Babylonian 公式计算新的估计值：(y/x + x) / 2
            }
            // 循环结束后，z 就是 y 的整数平方根
        } else if (y != 0) { // 如果 y 不等于 0 且小于等于 3
            z = 1; // 平方根为 1（因为 1*1=1, 2*2=4>3, 3*3=9>3）
        }
        // 如果 y == 0，z 默认为 0，不需要特殊处理
    }
}
