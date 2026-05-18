// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// 导入 Safe 多签辅助库
import {SafeHelper} from "../mocks/SafeHelper.sol";
// 导入被测试的预言机合约
import {BscPledgeOracle} from "./BscPledgeOracle.sol";
// 导入预言机接口（用于类型验证）
import {IBscPledgeOracle} from "./interfaces/IBscPledgeOracle.sol";
// 导入 Ownable 的 error 定义
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
// 导入 Hardhat 控制台日志
import {console} from "forge-std/src/console.sol";
import {Enum} from "@safe-global/safe-smart-account/contracts/libraries/Enum.sol";

// ============ Mock Chainlink 聚合器 ============

/**
 * @title MockAggregator
 * @notice 模拟 Chainlink 聚合器，用于测试预言机的 Chainlink 价格路径
 */
contract MockAggregator {
    // 当前轮次 ID
    uint80 private _roundId;
    // 当前价格（可能为负）
    int256 private _price;
    // 价格更新时间戳
    uint256 private _updatedAt;
    // 价格被确认的轮次
    uint80 private _answeredInRound;
    // 预言机精度（如 8 表示 8 位小数）
    uint8 private _decimals;

    /**
     * @notice 设置聚合器的轮次数据
     * @param roundId 轮次 ID
     * @param price 价格（支持负值用于异常测试）
     * @param updatedAt 更新时间戳，设为 0 可模拟不完整轮次
     * @param answeredInRound 价格被确认的轮次，小于 roundId 可模拟过期数据
     */
    function setLatestRoundData(uint80 roundId, int256 price, uint256 updatedAt, uint80 answeredInRound) external {
        _roundId = roundId;
        _price = price;
        _updatedAt = updatedAt;
        _answeredInRound = answeredInRound;
    }

    /**
     * @notice 设置聚合器的精度
     * @param d 精度值（如 6、8、18）
     */
    function setDecimals(uint8 d) external {
        _decimals = d;
    }

    /**
     * @notice 返回最新的轮次数据（与 Chainlink AggregatorV3Interface 一致）
     * @return roundId 轮次 ID
     * @return price 价格
     * @return startedAt 本轮起始时间（此处固定返回 0）
     * @return updatedAt 更新时间
     * @return answeredInRound 确认轮次
     */
    function latestRoundData() external view returns (uint80, int256, uint256, uint256, uint80) {
        return (_roundId, _price, 0, _updatedAt, _answeredInRound);
    }

    /**
     * @notice 返回聚合器精度
     * @return 精度值
     */
    function decimals() external view returns (uint8) {
        return _decimals;
    }
}

// ============ BscPledgeOracle 测试合约 ============

/**
 * @title BscPledgeOracleTest
 * @notice BscPledgeOracle 合约的完整测试套件
 * @dev 继承 SafeHelper，使用 Safe 多签钱包作为合约 owner，模拟生产环境
 */
contract BscPledgeOracleTest is SafeHelper {
    // 被测试的预言机实例
    BscPledgeOracle oracle;
    // 模拟的 Chainlink 聚合器
    MockAggregator aggregator;

    // 测试用资产地址 1
    address constant ASSET_1 = address(0x1111);
    // 测试用资产地址 2
    address constant ASSET_2 = address(0x2222);
    // 测试用资产地址 3
    address constant ASSET_3 = address(0x3333);
    // ASSET_1 对应的 uint256 标识
    uint256 constant UNDERLYING_1 = uint256(uint160(ASSET_1));
    // ASSET_2 对应的 uint256 标识
    uint256 constant UNDERLYING_2 = uint256(uint160(ASSET_2));
    // ASSET_3 对应的 uint256 标识
    uint256 constant UNDERLYING_3 = uint256(uint160(ASSET_3));
    // 随机地址，用于测试非 owner 权限检查
    address random = makeAddr("random");

    // BTC 模拟价格 ~$50000（18 位精度）
    uint256 constant PRICE_1 = 50000e18;
    // ETH 模拟价格 ~$3000（18 位精度）
    uint256 constant PRICE_2 = 3000e18;

    // 事件声明（用于 expectEmit 断言）
    event PriceDivisorUpdated(uint256 oldDivisor, uint256 newDivisor);
    event ManualPricesSet(uint256[] assets, uint256[] prices);
    event ManualPriceSet(uint256 asset, uint256 price);
    event AggregatorSet(uint256 asset, address aggregator, uint8 tokenDecimals);

    /**
     * @notice 每个测试用例执行前的初始化函数
     * @dev 部署 Safe 多签、预言机和模拟聚合器
     */
    function setUp() public {
        // 部署 Safe 多签钱包（双签名阈值）
        _deploySafe(0xA11CE, 0xB0B, 2);
        // 部署预言机，owner 为 Safe 地址
        oracle = new BscPledgeOracle(safeAddress);
        // 部署模拟聚合器
        aggregator = new MockAggregator();
    }

    // ==================== 构造函数测试 ====================

    /**
     * @notice 测试构造函数正确设置多签地址为 owner
     */
    function test_constructor_setsOwner() public view {
        assertEq(oracle.owner(), safeAddress);
    }

    /**
     * @notice 测试 priceDivisor 默认值为 1
     */
    function test_constructor_defaultPriceDivisor() public view {
        assertEq(oracle.priceDivisor(), 1);
    }

    /**
     * @notice 测试向零地址传参构造函数会 revert
     */
    function test_constructor_revertsZeroOwner() public {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector, address(0)));
        new BscPledgeOracle(address(0));
    }

    // ==================== setPriceDivisor 测试 ====================

    /**
     * @notice 测试通过 Safe 多签成功设置价格除数
     */
    function test_setPriceDivisor_updatesDivisor() public {
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPriceDivisor, (10)));
        assertEq(oracle.priceDivisor(), 10);
    }

    /**
     * @notice 测试 setPriceDivisor 正确触发 PriceDivisorUpdated 事件
     */
    function test_setPriceDivisor_emitsEvent() public {
        vm.expectEmit(true, false, false, true);
        emit PriceDivisorUpdated(1, 100);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPriceDivisor, (100)));
    }

    /**
     * @notice 测试将除数设为 0 会 revert
     */
    function test_setPriceDivisor_revertsZero() public {
        vm.expectRevert("Divisor cannot be zero");
        vm.prank(safeAddress);
        oracle.setPriceDivisor(0);
    }

    /**
     * @notice 测试非 owner 调用 setPriceDivisor 会 revert
     */
    function test_setPriceDivisor_revertsNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        oracle.setPriceDivisor(10);
    }

    // ==================== setPrices（批量）测试 ====================

    /**
     * @notice 测试批量设置资产价格成功且可正确查询
     */
    function test_setPrices_batchSetsPrices() public {
        // 准备资产标识数组
        uint256[] memory assets = new uint256[](2);
        assets[0] = UNDERLYING_1;
        assets[1] = UNDERLYING_2;
        // 准备对应的价格数组
        uint256[] memory prices = new uint256[](2);
        prices[0] = PRICE_1;
        prices[1] = PRICE_2;
        // 通过 Safe 多签执行批量设置
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPrices, (assets, prices)));
        // 验证两个资产的价格均正确设置
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_1), PRICE_1);
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_2), PRICE_2);
    }

    /**
     * @notice 测试 setPrices 正确触发 ManualPricesSet 事件
     */
    function test_setPrices_emitsEvent() public {
        // 准备单元素数组
        uint256[] memory assets = new uint256[](1);
        assets[0] = UNDERLYING_1;
        uint256[] memory prices = new uint256[](1);
        prices[0] = PRICE_1;
        // 断言事件触发
        vm.expectEmit(true, false, false, true);
        emit ManualPricesSet(assets, prices);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPrices, (assets, prices)));
    }

    /**
     * @notice 测试资产和价格数组长度不一致会 revert
     */
    function test_setPrices_revertsLengthMismatch() public {
        // 资产数组 2 个元素，价格数组 1 个元素
        uint256[] memory assets = new uint256[](2);
        assets[0] = UNDERLYING_1;
        assets[1] = UNDERLYING_2;
        uint256[] memory prices = new uint256[](1);
        prices[0] = PRICE_1;
        vm.expectRevert("Length mismatch");
        vm.prank(safeAddress);
        oracle.setPrices(assets, prices);
    }

    /**
     * @notice 测试非 owner 调用 setPrices 会 revert
     */
    function test_setPrices_revertsNonOwner() public {
        uint256[] memory assets = new uint256[](1);
        assets[0] = UNDERLYING_1;
        uint256[] memory prices = new uint256[](1);
        prices[0] = PRICE_1;
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        oracle.setPrices(assets, prices);
    }

    // ==================== setPrice（地址）测试 ====================

    /**
     * @notice 测试通过资产地址设置价格，验证地址到 underlying 的转换和查询一致性
     */
    function test_setPrice_setsPriceByAddress() public {
        // 通过 Safe 设置资产地址的价格
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPrice, (ASSET_1, PRICE_1)));
        // getPrice 应返回正确价格
        assertEq(oracle.getPrice(ASSET_1), PRICE_1);
        // 地址转 uint256 后应与 underlying 查询结果一致
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_1), PRICE_1);
    }

    /**
     * @notice 测试 setPrice 正确触发 ManualPriceSet 事件（asset 参数转为 underlying）
     */
    function test_setPrice_emitsEvent() public {
        vm.expectEmit(true, false, false, true);
        emit ManualPriceSet(UNDERLYING_1, PRICE_1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPrice, (ASSET_1, PRICE_1)));
    }

    /**
     * @notice 测试非 owner 调用 setPrice 会 revert
     */
    function test_setPrice_revertsNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        oracle.setPrice(ASSET_1, PRICE_1);
    }

    // ==================== setUnderlyingPrice 测试 ====================

    /**
     * @notice 测试通过 underlying 标识设置价格，验证 getPrice 和 getUnderlyingPrice 均正确返回
     */
    function test_setUnderlyingPrice_setsPrice() public {
        // 通过 Safe 设置 underlying 的价格
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, PRICE_1)));
        // getUnderlyingPrice 应返回正确价格
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_1), PRICE_1);
        // getPrice 通过地址查询也应返回相同价格
        assertEq(oracle.getPrice(ASSET_1), PRICE_1);
    }

    /**
     * @notice 测试 setUnderlyingPrice 正确触发 ManualPriceSet 事件
     */
    function test_setUnderlyingPrice_emitsEvent() public {
        vm.expectEmit(true, false, false, true);
        emit ManualPriceSet(UNDERLYING_1, PRICE_1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, PRICE_1)));
    }

    /**
     * @notice 测试非 owner 调用 setUnderlyingPrice 会 revert
     */
    function test_setUnderlyingPrice_revertsNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        oracle.setUnderlyingPrice(UNDERLYING_1, PRICE_1);
    }

    // ==================== setAssetsAggregator 测试 ====================

    /**
     * @notice 测试通过资产地址配置聚合器成功，验证查询返回正确地址和精度
     */
    function test_setAssetsAggregator_setsAggregator() public {
        // 通过 Safe 为 ASSET_1 配置聚合器和代币精度
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        // 查询聚合器配置
        (address agg, uint256 decimals) = oracle.getAssetsAggregator(ASSET_1);
        // 验证聚合器地址匹配
        assertEq(agg, address(aggregator));
        // 验证代币精度匹配
        assertEq(decimals, 18);
    }

    /**
     * @notice 测试 setAssetsAggregator 正确触发 AggregatorSet 事件
     */
    function test_setAssetsAggregator_emitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit AggregatorSet(UNDERLYING_1, address(aggregator), 18);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
    }

    /**
     * @notice 测试设置零地址聚合器会 revert
     */
    function test_setAssetsAggregator_revertsZeroAggregator() public {
        vm.expectRevert("Invalid aggregator");
        vm.prank(safeAddress);
        oracle.setAssetsAggregator(ASSET_1, address(0), 18);
    }

    /**
     * @notice 测试代币精度设为 0 会 revert
     */
    function test_setAssetsAggregator_revertsZeroDecimals() public {
        vm.expectRevert("Invalid token decimals");
        vm.prank(safeAddress);
        oracle.setAssetsAggregator(ASSET_1, address(aggregator), 0);
    }

    /**
     * @notice 测试代币精度超过 30 会 revert
     */
    function test_setAssetsAggregator_revertsExcessiveDecimals() public {
        vm.expectRevert("Invalid token decimals");
        vm.prank(safeAddress);
        oracle.setAssetsAggregator(ASSET_1, address(aggregator), 31);
    }

    /**
     * @notice 测试非 owner 调用 setAssetsAggregator 会 revert
     */
    function test_setAssetsAggregator_revertsNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        oracle.setAssetsAggregator(ASSET_1, address(aggregator), 18);
    }

    // ==================== setUnderlyingAggregator 测试 ====================
    // 注意：BscPledgeOracle.sol:101 存在拼写错误（onlyOwnerde），
    // 以下测试假设修复为 onlyOwner 后正常运行

    /**
     * @notice 测试通过 underlying 标识配置聚合器成功
     */
    function test_setUnderlyingAggregator_setsAggregator() public {
        // 通过 Safe 为 UNDERLYING_1 配置聚合器
        _executeViaSafe(
            address(oracle),
            abi.encodeCall(oracle.setUnderlyingAggregator, (UNDERLYING_1, address(aggregator), 18))
        );
        // 查询聚合器配置
        (address agg, uint256 decimals) = oracle.getUnderlyingAggregator(UNDERLYING_1);
        // 验证聚合器地址匹配
        assertEq(agg, address(aggregator));
        // 验证代币精度匹配
        assertEq(decimals, 18);
    }

    /**
     * @notice 测试 setUnderlyingAggregator 正确触发 AggregatorSet 事件
     */
    function test_setUnderlyingAggregator_emitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit AggregatorSet(UNDERLYING_1, address(aggregator), 18);
        _executeViaSafe(
            address(oracle),
            abi.encodeCall(oracle.setUnderlyingAggregator, (UNDERLYING_1, address(aggregator), 18))
        );
    }

    /**
     * @notice 测试非 owner 调用 setUnderlyingAggregator 会 revert
     */
    function test_setUnderlyingAggregator_revertsNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        oracle.setUnderlyingAggregator(UNDERLYING_1, address(aggregator), 18);
    }

    // ==================== 手动价格查询测试 ====================

    /**
     * @notice 设置手动价格的 modifier，减少重复代码
     */
    modifier withManualPrice() {
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, PRICE_1)));
        _;
    }

    /**
     * @notice 测试 getPrice 返回手动设置的价格
     */
    function test_getPrice_returnsManualPrice() public withManualPrice {
        assertEq(oracle.getPrice(ASSET_1), PRICE_1);
    }

    /**
     * @notice 测试 getUnderlyingPrice 返回手动设置的价格
     */
    function test_getUnderlyingPrice_returnsManualPrice() public withManualPrice {
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_1), PRICE_1);
    }

    /**
     * @notice 测试 getPrice 和 getUnderlyingPrice 对同一资产返回一致的价格
     */
    function test_getPrice_assetAndUnderlyingMatch() public withManualPrice {
        // getPrice(address) 内部调用 getUnderlyingPrice(uint256(address))，结果应一致
        assertEq(oracle.getPrice(ASSET_1), oracle.getUnderlyingPrice(UNDERLYING_1));
    }

    /**
     * @notice 测试 getPrices 批量查询函数正确返回多个手动价格
     */
    function test_getPrices_batchReturnsManualPrices() public {
        // 准备查询资产列表
        uint256[] memory assets = new uint256[](2);
        assets[0] = UNDERLYING_1;
        assets[1] = UNDERLYING_2;
        // 准备对应的价格列表
        uint256[] memory prices = new uint256[](2);
        prices[0] = PRICE_1;
        prices[1] = PRICE_2;
        // 通过 Safe 批量设置价格
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPrices, (assets, prices)));
        // 批量查询价格
        uint256[] memory results = oracle.getPrices(assets);
        // 验证结果数量和顺序
        assertEq(results.length, 2);
        assertEq(results[0], PRICE_1);
        assertEq(results[1], PRICE_2);
    }

    /**
     * @notice 测试未设置价格时返回 0
     */
    function test_getPrice_returnsZeroForUnsetPrice() public {
        assertEq(oracle.getPrice(ASSET_1), 0);
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_1), 0);
    }

    /**
     * @notice 测试覆盖已存在的手动价格
     */
    function test_getPrice_overwritesExistingPrice() public {
        // 先设置初始价格
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, PRICE_1)));
        assertEq(oracle.getPrice(ASSET_1), PRICE_1);
        // 覆盖为新价格
        uint256 newPrice = 60000e18;
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, newPrice)));
        assertEq(oracle.getPrice(ASSET_1), newPrice);
    }

    // ==================== Chainlink 价格查询测试 ====================

    /**
     * @notice 设置聚合器的 modifier，模拟 BTC/USD 预言机（8 位精度）
     */
    modifier withAggregator() {
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, int256(50000e8), block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        _;
    }

    /**
     * @notice 测试 getPrice 通过 Chainlink 聚合器正确获取并标准化价格
     */
    function test_getPrice_chainlinkPath() public withAggregator {
        // 预言机 8 位精度 → 标准化到 18 位：rawPrice * 10^(18-8) = 50000e8 * 1e10 = 50000e18
        uint256 expected = uint256(50000e8) * 1e10;
        assertEq(oracle.getPrice(ASSET_1), expected);
    }

    /**
     * @notice 测试聚合器价格优先于手动价格（双源情况下聚合器优先）
     */
    function test_getPrice_aggregatorTakesPriority() public withAggregator {
        // 同时设置手动价格和聚合器，聚合器应优先（getUnderlyingPrice 先检查 assetsMap）
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, 999e18)));
        // 期望返回聚合器价格而非手动价格
        uint256 expected = uint256(50000e8) * 1e10;
        assertEq(oracle.getPrice(ASSET_1), expected);
    }

    /**
     * @notice 测试预言机精度低于 18 时正确向上归一化（如 EUR/USD 6 位精度）
     */
    function test_getPrice_normalizesOracleDecimalsUp() public {
        // oracleDecimals=6（如 EUR/USD 聚合器）, tokenDecimals=18
        aggregator.setDecimals(6);
        aggregator.setLatestRoundData(1, 1e6, block.timestamp, 1); // $1
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        // Step1: rawPrice * 10^(18-6) = 1e6 * 1e12 = 1e18
        // Step2: tokenDecimals(18) == TARGET_DECIMALS, return rawPrice / divisor = 1e18
        assertEq(oracle.getPrice(ASSET_1), 1e18);
    }

    /**
     * @notice 测试预言机精度与目标精度一致时不做归一化
     */
    function test_getPrice_normalizesOracleDecimalsDown() public {
        // oracleDecimals=18（如某些 L2 预言机）, tokenDecimals=18
        aggregator.setDecimals(18);
        aggregator.setLatestRoundData(1, int256(1000e18), block.timestamp, 1); // $1000
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        // oracleDecimals == TARGET_DECIMALS, 无需精度转换
        // tokenDecimals(18) == TARGET_DECIMALS, return rawPrice / divisor = 1000e18
        assertEq(oracle.getPrice(ASSET_1), 1000e18);
    }

    /**
     * @notice 测试低精度代币（如 USDT 6 位）的价格缩放
     */
    function test_getPrice_appliesTokenDecimalsScale() public {
        // USDT 场景: tokenDecimals=6, 预言机 8 位精度
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, 1e8, block.timestamp, 1); // $1
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 6)));
        // Step1: oracle(8) < 18 → rawPrice * 10^10 = 1e8 * 1e10 = 1e18
        // Step2: tokenDecimals(6) < 18 → (1e18 * 10^12) / 1 = 1e30
        // 使 PledgePool 中 rawAmount(6位精度) * price / 1e18 = USD价值(18位)
        assertEq(oracle.getPrice(ASSET_1), 1e30);
    }

    /**
     * @notice 测试双资产等值计价：1 个 ASSET_1 等价于 1e10 个 ASSET_2
     * @dev ASSET_2 Chainlink 聚合器返回 1e8（$1），ASSET_1 返回 1e18（$1e10）
     *      两者 tokenDecimals=6，oracleDecimals=8
     *      PledgePool 中价值计算公式：value = rawAmount * getPrice / 1e18
     *      1 个 ASSET_1(6位精度) = 1 * 1e6 * getPrice(ASSET_1) / 1e18
     *                         = 1e6 * 1e40 / 1e18 = 1e28
     *      1e10 个 ASSET_2(6位精度) = 1e10 * 1e6 * getPrice(ASSET_2) / 1e18
     *                            = 1e16 * 1e30 / 1e18 = 1e28
     */
    function test_assetEquivalence_asset1Eq1e10Asset2() public {
        // 配置 ASSET_2 聚合器：价格 1e8（$1），tokenDecimals=6，oracleDecimals=8
        //表示1个ASSET_2  也就是 10**6 =1000000
        MockAggregator aggregator2 = new MockAggregator();
        aggregator2.setDecimals(8);
        aggregator2.setLatestRoundData(1, 1e8, block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(
            oracle.setAssetsAggregator, (ASSET_2, address(aggregator2), 6)
        ));

        // 配置 ASSET_1 聚合器：价格 1e18（$1e10），tokenDecimals=6，oracleDecimals=8
        MockAggregator aggregator1 = new MockAggregator();
        aggregator1.setDecimals(8);
        aggregator1.setLatestRoundData(1, 1e18, block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(
            oracle.setAssetsAggregator, (ASSET_1, address(aggregator1), 6)
        ));

        // 验证 getPrice
        // ASSET_2: Step1 1e8 * 1e10 = 1e18, Step2 1e18 * 1e12 / 1 = 1e30
        assertEq(oracle.getPrice(ASSET_2), 1e30);
        // ASSET_1: Step1 1e18 * 1e10 = 1e28, Step2 1e28 * 1e12 / 1 = 1e40
        assertEq(oracle.getPrice(ASSET_1), 1e40);

        // 计算 1 个 ASSET_1 的价值（1 whole token = 1e6 atomic units）
        // value = amount * price / 1e18 = 1e6 * 1e40 / 1e18 = 1e28
        uint256 valueOfOneAsset1 = 1e6 * oracle.getPrice(ASSET_1) / 1e18;
        assertEq(valueOfOneAsset1, 1e28);

        // 计算 1e10 个 ASSET_2 的价值（1e10 whole tokens = 1e16 atomic units）
        // value = 1e16 * 1e30 / 1e18 = 1e28
        uint256 valueOf1e10Asset2 = (1e10 * 1e6) * oracle.getPrice(ASSET_2) / 1e18;
        assertEq(valueOf1e10Asset2, 1e28);

        // 两者价值相等：1 个 ASSET_1 = 1e10 个 ASSET_2
        assertEq(valueOfOneAsset1, valueOf1e10Asset2);
    }

    /**
     * @notice 测试高精度代币（如 30 位）的价格缩放
     */
    function test_getPrice_applyLessTokenDecimalsScale() public {
        // 假设代币 30 位精度（极端情况）
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, 1e8, block.timestamp, 1); // $1
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 30)));
        // Step1: rawPrice * 10^10 = 1e18
        // Step2: tokenDecimals(30) > 18 → (1e18 / 10^12) / 1 = 1e6
        assertEq(oracle.getPrice(ASSET_1), 1e6);
    }

    /**
     * @notice 测试通过 setUnderlyingAggregator 配置的聚合器正常返回价格
     */
    function test_getPrice_underlyingAggregator() public {
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, int256(2000e8), block.timestamp, 1); // ETH $2000
        // 通过 underlying 标识配置聚合器
        _executeViaSafe(
            address(oracle),
            abi.encodeCall(oracle.setUnderlyingAggregator, (UNDERLYING_1, address(aggregator), 18))
        );
        // 预期价格：2000e8 * 1e10 = 2000e18
        uint256 expected = uint256(2000e8) * 1e10;
        assertEq(oracle.getUnderlyingPrice(UNDERLYING_1), expected);
    }

    // ==================== 价格除数影响测试 ====================

    /**
     * @notice 测试 priceDivisor 正确影响 Chainlink 路径的价格计算
     */
    function test_priceDivisor_affectsChainlinkPrice() public {
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, int256(50000e8), block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        // 除数设为 100
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPriceDivisor, (100)));
        // 原价 50000e18 / 100 = 500e18
        assertEq(oracle.getPrice(ASSET_1), 500e18);
    }

    /**
     * @notice 测试 priceDivisor 仅影响 Chainlink 路径，不影响手动价格
     */
    function test_priceDivisor_onlyAffectsChainlinkPath() public {
        // priceDivisor 仅影响 _getChainlinkPrice 路径，不影响 priceMap 手动价格
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_1, PRICE_1)));
        // 设置除数
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setPriceDivisor, (100)));
        // 手动价格不受除数影响，仍为原始值
        assertEq(oracle.getPrice(ASSET_1), PRICE_1);
    }

    // ==================== Chainlink 异常处理测试 ====================

    /**
     * @notice 测试预言机返回负价格时正确 revert
     */
    function test_getPrice_revertsNegativePrice() public {
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, -1, block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        // 期望 "Oracle price <= 0" 错误
        vm.expectRevert("Oracle price <= 0");
        oracle.getPrice(ASSET_1);
    }

    /**
     * @notice 测试预言机返回零价格时正确 revert
     */
    function test_getPrice_revertsZeroPrice() public {
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, 0, block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        vm.expectRevert("Oracle price <= 0");
        oracle.getPrice(ASSET_1);
    }

    /**
     * @notice 测试预言机返回过期轮次数据时正确 revert
     */
    function test_getPrice_revertsStaleRound() public {
        aggregator.setDecimals(8);
        // answeredInRound(1) < roundId(2) ⇒ 数据过时
        aggregator.setLatestRoundData(2, int256(50000e8), block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        vm.expectRevert("Stale round");
        oracle.getPrice(ASSET_1);
    }

    /**
     * @notice 测试 answeredInRound 等于 roundId 时不触发过时检查
     */
    function test_getPrice_staleRoundEqRoundPasses() public {
        // answeredInRound == roundId 应通过检查
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(5, int256(50000e8), block.timestamp, 5);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        // 不应 revert
        oracle.getPrice(ASSET_1);
    }

    /**
     * @notice 测试预言机返回不完整轮次数据时正确 revert
     */
    function test_getPrice_revertsIncompleteRound() public {
        aggregator.setDecimals(8);
        // updatedAt == 0 ⇒ 轮次未完成
        aggregator.setLatestRoundData(1, int256(50000e8), 0, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        vm.expectRevert("Incomplete round");
        oracle.getPrice(ASSET_1);
    }

    // ==================== 聚合器查询测试 ====================

    /**
     * @notice 测试未设置聚合器时查询返回空值
     */
    function test_getAssetsAggregator_returnsEmptyBeforeSet() public {
        (address agg, uint256 decimals) = oracle.getAssetsAggregator(ASSET_1);
        assertEq(agg, address(0));
        assertEq(decimals, 0);
    }

    /**
     * @notice 测试设置后正确查询聚合器信息
     */
    function test_getAssetsAggregator_returnsAfterSet() public {
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 8)));
        (address agg, uint256 decimals) = oracle.getAssetsAggregator(ASSET_1);
        assertEq(agg, address(aggregator));
        assertEq(decimals, 8);
    }

    /**
     * @notice 测试未设置前查询 underlying 聚合器返回空值
     */
    function test_getUnderlyingAggregator_returnsEmptyBeforeSet() public {
        (address agg, uint256 decimals) = oracle.getUnderlyingAggregator(UNDERLYING_1);
        assertEq(agg, address(0));
        assertEq(decimals, 0);
    }

    // ==================== getPrices 批量查询混合场景 ====================

    /**
     * @notice 测试批量查询混合价格源（一个聚合器 + 一个手动价格）
     */
    function test_getPrices_batchMixedSources() public {
        // ASSET_1: 聚合器, ASSET_2: 手动价格
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, int256(50000e8), block.timestamp, 1);
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setAssetsAggregator, (ASSET_1, address(aggregator), 18)));
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (UNDERLYING_2, PRICE_2)));
        // 批量查询两个不同价格源的价格
        uint256[] memory assets = new uint256[](2);
        assets[0] = UNDERLYING_1;
        assets[1] = UNDERLYING_2;
        uint256[] memory results = oracle.getPrices(assets);
        // 验证结果
        assertEq(results.length, 2);
        assertEq(results[0], uint256(50000e8) * 1e10);
        assertEq(results[1], PRICE_2);
    }

    // ==================== 原生代币（underlying = 0）测试 ====================

    /**
     * @notice 测试 underlying=0 代表原生代币（如 BNB）的手动价格设置和查询
     */
    function test_nativeTokenPrice() public {
        // underlying = 0 代表 BNB 等原生代币
        _executeViaSafe(address(oracle), abi.encodeCall(oracle.setUnderlyingPrice, (0, PRICE_1)));
        assertEq(oracle.getUnderlyingPrice(0), PRICE_1);
    }

    /**
     * @notice 测试原生代币通过聚合器查询价格
     */
    function test_nativeTokenAggregator() public {
        aggregator.setDecimals(8);
        aggregator.setLatestRoundData(1, int256(300e8), block.timestamp, 1); // BNB $300
        // 为原生代币配置聚合器
        _executeViaSafe(
            address(oracle),
            abi.encodeCall(oracle.setUnderlyingAggregator, (0, address(aggregator), 18))
        );
        // 预期价格：300e8 * 1e10 = 300e18
        uint256 expected = uint256(300e8) * 1e10;
        assertEq(oracle.getUnderlyingPrice(0), expected);
    }

    // ==================== Safe 多签安全测试 ====================

    /**
     * @notice 测试单签名无法执行管理操作（确保多签阈值 2 生效）
     * @dev 手动构造单签名交易并通过 Safe 执行，预期 revert
     */
    function test_safeRevertsWithSingleSignature() public {
        // 构造 setPriceDivisor 调用数据
        bytes memory data = abi.encodeCall(oracle.setPriceDivisor, (10));
        // 获取 Safe 交易哈希
        bytes32 txHash = safe.getTransactionHash(
            address(oracle), 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), safe.nonce()
        );
        // 仅使用第一个 owner 的签名（不满足阈值 2）
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(_ownerPrivateKey1, txHash);
        bytes memory sig = abi.encodePacked(r, s, v);
        // 预期 Safe 执行失败
        vm.expectRevert();
        safe.execTransaction(
            address(oracle), 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), sig
        );
        // priceDivisor 应保持默认值 1，未被修改
        assertEq(oracle.priceDivisor(), 1);
    }
}
