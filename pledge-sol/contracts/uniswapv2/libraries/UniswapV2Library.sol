// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

import "../interfaces/IUniswapV2Pair.sol"; // 导入交易对接口
import "../interfaces/IUniswapV2Factory.sol"; // 导入工厂接口

/**
 * @title UniswapV2Library 库
 * @notice 提供 UniswapV2 常用辅助函数
 * @dev 这个库提供了 UniswapV2 协议的核心计算功能，包括：
 *      - 计算交易对地址（使用 CREATE2）
 *      - 代币地址排序
 *      - 获取储备量
 *      - 计算交换金额
 */
library UniswapV2Library {
    /**
     * @notice 计算交易对创建地址（CREATE2）
     * @param factory 工厂合约地址
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @return pair 交易对地址
     * @dev 使用 CREATE2 操作码确定性计算交易对地址
     *      CREATE2 允许在合约部署前就知道其地址
     *      公式：address = keccak256(0xff ++ factory_address ++ salt ++ keccak256(init_code))[12:]
     */
    function pairFor(address factory, address tokenA, address tokenB)
        internal
        view
        returns (address pair)
    {
        // 首先对代币地址排序，确保计算结果一致
        (address token0, address token1) = sortTokens(tokenA, tokenB);
        // 从工厂合约获取 init code hash，避免硬编码
        bytes32 hash = IUniswapV2Factory(factory).initCodeHash();
        // 使用 CREATE2 计算交易对地址
        pair = address(
            uint160( // 将 uint256 转换为 address（20 字节）
                uint256(
                    keccak256( // 计算 keccak256 哈希
                        abi.encodePacked( // 打包编码以下数据
                            hex"ff", // CREATE2 前缀，固定值 0xff
                            factory, // 工厂合约地址
                            keccak256(abi.encodePacked(token0, token1)), // salt：两个代币地址的哈希
                            hash // init code hash：从工厂合约动态获取
                        )
                    )
                )
            )
        );
    }

    /**
     * @notice 对两个代币地址排序
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @return token0 较小的地址
     * @return token1 较大的地址
     * @dev 排序确保无论用户以何种顺序提供代币地址，结果始终一致
     *      token0 总是较小的地址，token1 总是较大的地址
     */
    function sortTokens(address tokenA, address tokenB)
        internal
        pure
        returns (address token0, address token1)
    {
        // 确保两个地址不相同
        require(tokenA != tokenB, "UniswapV2Library: IDENTICAL_ADDRESSES");
        // 排序：较小的地址为 token0，较大的地址为 token1
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        // 确保 token0 不是零地址（如果 token0 不是零地址，token1 也不可能是零地址）
        require(token0 != address(0), "UniswapV2Library: ZERO_ADDRESS");
    }

    /**
     * @notice 获取储备量和代币顺序
     * @param factory 工厂合约地址
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @return reserveA 代币A的储备量
     * @return reserveB 代币B的储备量
     * @dev 交易对内部按 token0/token1 顺序存储储备量
     *      此函数根据用户提供的代币顺序返回对应的储备量
     */
    function getReserves(address factory, address tokenA, address tokenB)
        internal
        view
        returns (uint256 reserveA, uint256 reserveB)
    {
        // 获取排序后的 token0
        (address token0,) = sortTokens(tokenA, tokenB);
        // 从交易对合约获取储备量
        (uint256 reserve0, uint256 reserve1,) = IUniswapV2Pair(pairFor(factory, tokenA, tokenB)).getReserves();
        // 根据代币顺序返回对应的储备量
        // 如果 tokenA 是 token0，返回 (reserve0, reserve1)
        // 如果 tokenA 是 token1，返回 (reserve1, reserve0)
        (reserveA, reserveB) = tokenA == token0 ? (reserve0, reserve1) : (reserve1, reserve0);
    }

    /**
     * @notice 根据输入计算输出金额
     * @param amountIn 输入金额
     * @param reserveIn 输入代币储备量
     * @param reserveOut 输出代币储备量
     * @return amountOut 输出金额
     * @dev 使用恒定乘积公式 x * y = k 计算输出金额
     *      考虑 0.3% 的交易手续费
     *
     *      Δy = (y * Δx) / (x + Δx)  Δy输出的token数量 Δx输入的token数量 x当前的输入token的储备 y当前的输出token的储备
     *      不带手续费公式：amountOut = (amountIn * reserveOut) / (reserveIn + amountIn )
     *      带手续费 Δy = (y * (Δx * (1-0.003))) / (x + Δx * (1-0.003))    = Δy = (y * (Δx * 0.997 )) / (x + Δx * 0.997 )
     *      为了好计算都乘以1000 转化为 Δy = (y * Δx * 0.997 ) * 1000 / (x + Δx * 0.997) * 1000
     *      输出 Δy = (y * Δx * 997) / (x * 1000 + Δx * 997)
     *      最终公式：amountOut = (reserveOut * amountIn * 997) / (reserveIn * 1000 + amountIn * 997)
     */
    function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut)
        internal
        pure
        returns (uint256 amountOut)
    {
        // 确保输入金额大于 0
        require(amountIn > 0, "UniswapV2Library: INSUFFICIENT_INPUT_AMOUNT");
        // 确保两个储备量都大于 0
        require(reserveIn > 0 && reserveOut > 0, "UniswapV2Library: INSUFFICIENT_LIQUIDITY");
        // 计算扣除手续费后的输入金额（997/1000 = 99.7%，即扣除 0.3% 手续费）
        uint256 amountInWithFee = amountIn * 997;
        // 计算分子：扣除手续费后的输入金额 * 输出代币储备量
        uint256 numerator = reserveOut * amountInWithFee;
        // 计算分母：输入代币储备量 * 1000 + 扣除手续费后的输入金额
        uint256 denominator = reserveIn * 1000 + amountInWithFee;
        // 计算输出金额
        amountOut = numerator / denominator;
    }

    /**
     * @notice 根据输出计算输入金额
     * @param amountOut 输出金额
     * @param reserveIn 输入代币储备量
     * @param reserveOut 输出代币储备量
     * @return amountIn 输入金额
     * @dev 反向计算：已知期望输出，计算需要的输入金额
     *      公式：amountIn = (reserveIn * amountOut * 1000) / ((reserveOut - amountOut) * 997) + 1
     */
    function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut)
        internal
        pure
        returns (uint256 amountIn)
    {
        // 确保输出金额大于 0
        require(amountOut > 0, "UniswapV2Library: INSUFFICIENT_OUTPUT_AMOUNT");
        // 确保两个储备量都大于 0
        require(reserveIn > 0 && reserveOut > 0, "UniswapV2Library: INSUFFICIENT_LIQUIDITY");
        // 计算分子：输入代币储备量 * 输出金额 * 1000
        uint256 numerator = reserveIn * amountOut * 1000;
        // 计算分母：(输出代币储备量 - 输出金额) * 997
        uint256 denominator = (reserveOut - amountOut) * 997;
        // 计算输入金额，加 1 向上取整
        amountIn = numerator / denominator + 1;
    }
}
