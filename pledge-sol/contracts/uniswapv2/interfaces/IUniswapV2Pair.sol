// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

/**
 * @title IUniswapV2Pair 接口
 * @notice 定义交易对合约的核心功能
 * @dev 交易对合约是 UniswapV2 的核心，实现了 AMM（自动做市商）逻辑
 *      每个交易对包含两种代币的流动性池
 */
interface IUniswapV2Pair {
    /**
     * @notice 获取工厂合约地址
     * @return 工厂合约地址
     */
    function factory() external view returns (address);

    /**
     * @notice 获取代币0地址
     * @return 代币0地址（地址较小的代币）
     */
    function token0() external view returns (address);

    /**
     * @notice 获取代币1地址
     * @return 代币1地址（地址较大的代币）
     */
    function token1() external view returns (address);

    /**
     * @notice 获取储备量
     * @return reserve0 代币0储备量
     * @return reserve1 代币1储备量
     * @return blockTimestampLast 最后更新时间戳
     * @dev 储备量用于计算交易价格和流动性
     */
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);

    /**
     * @notice 获取代币0的累计价格
     * @return 代币0的累计价格值
     * @dev 用于计算时间加权平均价格（TWAP）
     */
    function price0CumulativeLast() external view returns (uint256);

    /**
     * @notice 获取代币1的累计价格
     * @return 代币1的累计价格值
     * @dev 用于计算时间加权平均价格（TWAP）
     */
    function price1CumulativeLast() external view returns (uint256);

    /**
     * @notice 获取上次 K 值
     * @return 上次记录的 K 值（reserve0 * reserve1）
     * @dev 用于协议手续费计算
     */
    function kLast() external view returns (uint256);

    /**
     * @notice 铸造 LP 代币
     * @param to 接收地址
     * @return liquidity 铸造的流动性数量
     * @dev 当用户向交易对添加流动性时调用
     *      用户需要先将代币转入交易对合约，然后调用此函数
     */
    function mint(address to) external returns (uint256 liquidity);

    /**
     * @notice 销毁 LP 代币并取回代币
     * @param to 接收地址
     * @return amount0 取回的代币0数量
     * @return amount1 取回的代币1数量
     * @dev 用户将 LP 代币转回交易对，然后调用此函数取回代币
     */
    function burn(address to) external returns (uint256 amount0, uint256 amount1);

    /**
     * @notice 执行代币交换
     * @param amount0Out 输出的代币0数量
     * @param amount1Out 输出的代币1数量
     * @param to 接收地址
     * @param data 闪电贷回调数据（如果非空则触发闪电贷）
     * @dev 执行代币交换，支持普通交换和闪电贷
     *      闪电贷：先借出代币，在回调中归还，无需抵押
     */
    function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes calldata data) external;

    /**
     * @notice 强制同步储备量
     * @dev 将合约余额同步到储备量
     *      用于处理直接转账到合约的情况
     */
    function skim(address to) external;

    /**
     * @notice 强制使储备量与余额匹配
     * @dev 更新储备量以匹配当前合约余额
     *      用于处理储备量与实际余额不一致的情况
     */
    function sync() external;

    /**
     * @notice 初始化交易对
     * @param _token0 代币0地址
     * @param _token1 代币1地址
     * @dev 由工厂合约在创建交易对时调用一次
     */
    function initialize(address _token0, address _token1) external;
}
