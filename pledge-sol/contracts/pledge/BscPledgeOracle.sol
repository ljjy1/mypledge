// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "./interfaces/IBscPledgeOracle.sol";


contract BscPledgeOracle is Ownable, IBscPledgeOracle {

    // ============ 状态变量 ============
    // 资产标识 => Chainlink 聚合器
    mapping(uint256 => AggregatorV3Interface) internal assetsMap;
    // 资产标识 => 代币自身小数位数 (1~30)
    mapping(uint256 => uint8) internal decimalsMap;
    // 资产标识 => 手动价格（已标准化至 18 位精度）
    mapping(uint256 => uint256) internal priceMap;

    uint256 internal constant TARGET_DECIMALS = 18;
    // 全局价格除数（默认 1）
    uint256 public priceDivisor = 1;


    /**
     * 构造函数
     * _owner 多签钱包地址
     */
    constructor(address _owner) Ownable(_owner) {

    }

    // ============ 管理函数 ============

    /**
     * @notice 设置全局价格除数
     * @param newDivisor 新除数，必须 > 0
     */
    function setPriceDivisor(uint256 newDivisor) external onlyOwner {
        require(newDivisor > 0, "Divisor cannot be zero");
        uint256 old = priceDivisor;
        priceDivisor = newDivisor;
        emit PriceDivisorUpdated(old, newDivisor);
    }

    /**
     * @notice 批量设置手动价格（适用于无预言机资产）
     * @param assets 资产标识数组
     * @param prices 对应的标准化价格（18 位精度）
     */
    function setPrices(uint256[] calldata assets, uint256[] calldata prices) external onlyOwner {
        require(assets.length == prices.length, "Length mismatch");
        for (uint256 i = 0; i < assets.length; i++) {
            priceMap[assets[i]] = prices[i];
        }
        emit ManualPricesSet(assets, prices);
    }

    /**
     * @notice 设置单个资产的手动价格
     * @param asset 资产地址（会转为 uint256 标识）
     * @param price 标准化价格（18 位精度）
     */
    function setPrice(address asset, uint256 price) external onlyOwner {
        uint256 underlying = uint256(uint160(asset));
        priceMap[underlying] = price;
        emit ManualPriceSet(underlying, price);
    }

    /**
     * @notice 设置底层资产的手动价格
     * @param underlying 资产标识（0 代表原生代币）
     * @param price 标准化价格（18 位精度）
     */
    function setUnderlyingPrice(uint256 underlying, uint256 price) external onlyOwner {
        priceMap[underlying] = price;
        emit ManualPriceSet(underlying, price);
    }

    /**
     * @notice 为资产配置 Chainlink 聚合器
     * @param asset 资产地址
     * @param aggregator Chainlink 聚合器地址
     * @param tokenDecimals 代币自身小数位数 (1~30)
     */
    function setAssetsAggregator(address asset, address aggregator, uint8 tokenDecimals) external onlyOwner {
        _setAggregator(uint256(uint160(asset)), aggregator, tokenDecimals);
    }

    /**
     * @notice 为底层资产配置 Chainlink 聚合器
     * @param underlying 资产标识（0 代表原生代币）
     * @param aggregator Chainlink 聚合器地址
     * @param tokenDecimals 代币自身小数位数 (1~30)
     */
    function setUnderlyingAggregator(uint256 underlying, address aggregator, uint8 tokenDecimals) external onlyOwner {
        _setAggregator(underlying, aggregator, tokenDecimals);
    }

    // ============ 内部函数 ============

    function _setAggregator(uint256 assetId, address aggregator, uint8 tokenDecimals) internal {
        require(aggregator != address(0), "Invalid aggregator");
        require(tokenDecimals > 0 && tokenDecimals <= 30, "Invalid token decimals");
        assetsMap[assetId] = AggregatorV3Interface(aggregator);
        decimalsMap[assetId] = tokenDecimals;
        emit AggregatorSet(assetId, aggregator, tokenDecimals);

    }

    // ============ 价格查询 ============

    /**
     * @notice 获取单个资产价格（标准化为 18 位精度）
     * @param asset 资产地址
     */
    function getPrice(address asset) public view returns (uint256) {
        return getUnderlyingPrice(uint256(uint160(asset)));
    }

    /**
     * @notice 批量获取资产价格
     */
    function getPrices(uint256[] memory assets) public view returns (uint256[] memory) {
        uint256[] memory prices = new uint256[](assets.length);
        for (uint256 i = 0; i < assets.length; i++) {
            prices[i] = getUnderlyingPrice(assets[i]);
        }
        return prices;
    }

    /**
     * @dev 获取标准化价格的核心逻辑
     * @param underlying 资产标识
     * @return 标准化价格（18 位精度）
     */
    function getUnderlyingPrice(uint256 underlying) public view returns (uint256) {
        AggregatorV3Interface aggregator = assetsMap[underlying];
        if (address(aggregator) != address(0)) {
            return _getChainlinkPrice(aggregator, decimalsMap[underlying]);
        }
        return priceMap[underlying];
    }

    /**
     * @dev 从 Chainlink 聚合器获取价格并标准化
     * @notice 所有算术运算均受 Solidity 0.8 原生溢出检查保护
     */
    function _getChainlinkPrice(AggregatorV3Interface aggregator, uint8 tokenDecimals)
    internal view returns (uint256)
    {
        (
            uint80 roundId,
            int256 price,
            ,
            uint256 updatedAt,
            uint80 answeredInRound
        ) = aggregator.latestRoundData();

        require(price > 0, "Oracle price <= 0");
        require(answeredInRound >= roundId, "Stale round");
        require(updatedAt != 0, "Incomplete round");
        // 已移除价格有效期限检查

        uint8 oracleDecimals = aggregator.decimals();
        uint256 rawPrice = uint256(price);
        uint256 divisor = priceDivisor;

        // 第一步：将预言机价格精度统一到 TARGET_DECIMALS
        if (oracleDecimals < TARGET_DECIMALS) {
            rawPrice *= 10 ** (TARGET_DECIMALS - oracleDecimals);
        } else if (oracleDecimals > TARGET_DECIMALS) {
            rawPrice /= 10 ** (oracleDecimals - TARGET_DECIMALS);
        }

        // 第二步：根据代币精度调整，并除以全局除数
        // 为减少精度损失，需要放大时优先乘法，需要缩小时优先除法
        if (tokenDecimals < TARGET_DECIMALS) {
            uint256 scale = 10 ** (TARGET_DECIMALS - tokenDecimals);
            return (rawPrice * scale) / divisor;
        } else if (tokenDecimals > TARGET_DECIMALS) {
            uint256 scale = 10 ** (tokenDecimals - TARGET_DECIMALS);
            return (rawPrice / scale) / divisor;
        } else {
            return rawPrice / divisor;
        }
    }

    // ============ 向后兼容的查询 ============

    function getAssetsAggregator(address asset) public view returns (address, uint256) {
        uint256 id = uint256(uint160(asset));
        return (address(assetsMap[id]), decimalsMap[id]);
    }

    function getUnderlyingAggregator(uint256 underlying) public view returns (address, uint256) {
        return (address(assetsMap[underlying]), decimalsMap[underlying]);
    }
}
