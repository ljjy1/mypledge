// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title IBscPledgeOracle
 * @notice BscPledgeOracle 价格预言机接口
 * @dev 支持 Chainlink 聚合器和手动价格两种价格源
 */
interface IBscPledgeOracle {
    // ============ 事件 ============

    /// @notice 全局价格除数更新事件
    event PriceDivisorUpdated(uint256 oldDivisor, uint256 newDivisor);
    /// @notice 批量手动价格设置事件
    event ManualPricesSet(uint256[] assets, uint256[] prices);
    /// @notice 单资产手动价格设置事件
    event ManualPriceSet(uint256 asset, uint256 price);
    /// @notice 聚合器配置事件
    event AggregatorSet(uint256 asset, address aggregator, uint8 tokenDecimals);

    // ============ 管理函数 ============

    /// @notice 设置全局价格除数
    function setPriceDivisor(uint256 newDivisor) external;

    /// @notice 批量设置手动价格
    function setPrices(uint256[] calldata assets, uint256[] calldata prices) external;

    /// @notice 通过地址设置资产手动价格
    function setPrice(address asset, uint256 price) external;

    /// @notice 通过资产标识设置手动价格
    function setUnderlyingPrice(uint256 underlying, uint256 price) external;

    /// @notice 通过地址为资产配置 Chainlink 聚合器
    function setAssetsAggregator(address asset, address aggregator, uint8 tokenDecimals) external;

    /// @notice 通过资产标识配置 Chainlink 聚合器
    function setUnderlyingAggregator(uint256 underlying, address aggregator, uint8 tokenDecimals) external;

    // ============ 查询函数 ============

    /// @notice 全局价格除数
    function priceDivisor() external view returns (uint256);

    /// @notice 获取单资产价格（通过地址）
    function getPrice(address asset) external view returns (uint256);

    /// @notice 批量获取资产价格
    function getPrices(uint256[] calldata assets) external view returns (uint256[] memory);

    /// @notice 获取单资产价格（通过资产标识）
    function getUnderlyingPrice(uint256 underlying) external view returns (uint256);

    /// @notice 获取资产的聚合器配置（通过地址）
    function getAssetsAggregator(address asset) external view returns (address, uint256);

    /// @notice 获取资产的聚合器配置（通过资产标识）
    function getUnderlyingAggregator(uint256 underlying) external view returns (address, uint256);
}
