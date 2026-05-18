// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

/**
 * @title UQ112x112 库
 * @notice 定点数运算库
 * @dev 用于价格计算的定点数格式，112位小数部分
 *      格式：Q112.112，即整数部分112位，小数部分112位
 *      这种格式可以精确表示小数价格，避免浮点数精度问题
 *      例如：价格 1.5 可以表示为 1.5 * 2^112 = 1.5 * Q112
 */
library UQ112x112 {
    // Q112 格式的单位：2^112 = 5192296858534827628530496329220096
    // 这是一个常量，用于将整数转换为定点数
    // 乘以 Q112 相当于左移 112 位，将小数点放在第 112 位之后
    uint224 constant Q112 = 2 ** 112;

    /**
     * @notice 将 uint112 编码为 UQ112x112 格式
     * @param y 要编码的值（uint112 类型，最大值为 2^112 - 1）
     * @return z 编码后的定点数（uint224 类型）
     * @dev 编码过程：将 y 乘以 Q112，相当于将 y 左移 112 位
     *      这样 y 就变成了定点数，小数点在第 112 位之后
     *      例如：y = 1，编码后 z = Q112 = 2^112，表示 1.0
     */
    function encode(uint112 y) internal pure returns (uint224 z) {
        z = uint224(y) * Q112; // 将 y 转换为 uint224 并乘以 Q112，完成编码
    }

    /**
     * @notice UQ112x112 格式除法
     * @param x 被除数（定点数，UQ112x112 格式）
     * @param y 除数（普通整数，uint112 类型）
     * @return z 商（定点数，UQ112x112 格式）
     * @dev 定点数除法：x / y，结果仍然是定点数
     *      因为 x 已经是乘以 Q112 的值，所以直接除以 y 即可
     *      例如：x = 1.5 * Q112，y = 2，结果 z = 0.75 * Q112
     *      这个函数常用于计算价格比率
     */
    function uqdiv(uint224 x, uint112 y) internal pure returns (uint224 z) {
        z = x / uint224(y); // 将 y 转换为 uint224 并执行除法，结果为定点数
    }
}
