// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

import "forge-std/src/Test.sol"; // 导入 Foundry 测试库
import "./UniswapV2Factory.sol"; // 导入工厂合约
import "./UniswapV2Router02.sol"; // 导入路由合约
import "./UniswapV2Pair.sol"; // 导入交易对合约
import "./WETH.sol"; // 导入 WETH 合约
import "@openzeppelin/contracts/token/ERC20/ERC20.sol"; // 导入 OpenZeppelin 的 ERC20 实现
import "@openzeppelin/contracts/token/ERC20/IERC20.sol"; // 导入 OpenZeppelin 的 ERC20 接口
import "../mocks/MockTestERC20.sol";


/**
 * @title FeeToken 带手续费的测试代币
 * @notice 用于测试手续费代币的交换
 * @dev 这个代币在每次转账时收取 1% 的手续费
 *      用于测试 UniswapV2 对手续费代币的支持
 */
contract FeeToken is ERC20 {
    uint256 public feeRate = 10; // 手续费率：10/1000 = 1%

    /**
     * @notice 构造函数
     * @param name 代币名称
     * @param symbol 代币符号
     * @param initialSupply 初始供应量
     */
    constructor(string memory name, string memory symbol, uint256 initialSupply) ERC20(name, symbol) {
        _mint(msg.sender, initialSupply); // 铸造初始供应量给部署者
    }

    /**
     * @notice 重写 transfer 函数，添加手续费逻辑
     * @param to 接收地址
     * @param amount 转账金额
     * @return 是否成功
     */
    function transfer(address to, uint256 amount) public override returns (bool) {
        uint256 fee = amount * feeRate / 1000; // 计算手续费：金额 * 1%
        _transfer(_msgSender(), address(0xdead), fee); // 将手续费发送到销毁地址
        _transfer(_msgSender(), to, amount - fee); // 将剩余金额发送给接收者
        return true;
    }

    /**
     * @notice 重写 transferFrom 函数，添加手续费逻辑
     * @param from 发送地址
     * @param to 接收地址
     * @param amount 转账金额
     * @return 是否成功
     */
    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        uint256 fee = amount * feeRate / 100; // 计算手续费：金额 * 1%
        _transfer(from, address(0xdead), fee); // 将手续费发送到销毁地址
        _transfer(from, to, amount - fee); // 将剩余金额发送给接收者
        return true;
    }
}

/**
 * @title UniswapV2 测试合约
 * @notice 测试 UniswapV2 核心功能
 * @dev 这个测试合约覆盖了 UniswapV2 的主要功能：
 *      - Factory：创建交易对、设置费用接收地址
 *      - 流动性管理：添加/移除流动性
 *      - 代币交换：各种交换方式
 *      - 手续费代币支持
 *      - 价格计算
 *      - Pair 操作
 *      - WETH 功能
 */
contract UniswapV2Test is Test {
    UniswapV2Factory public factory; // 工厂合约实例
    UniswapV2Router02 public router; // 路由合约实例
    WETH public weth; // WETH 合约实例
    MockTestERC20 public tokenA; // 测试代币 A
    MockTestERC20 public tokenB; // 测试代币 B
    FeeToken public feeToken; // 带手续费的测试代币

    address public alice = makeAddr("alice"); // 测试账户 alice
    address public bob = makeAddr("bob"); // 测试账户 bob
    address public charlie = makeAddr("charlie"); // 测试账户 charlie

    uint256 constant ONE_TOKEN = 10 ** 18; // 一个代币的单位（18 位精度）

    /**
     * @notice 测试初始化函数
     * @dev 在每个测试前运行，部署合约并设置初始状态
     */
    function setUp() public {
        // 部署 WETH 合约
        weth = new WETH();
        // 部署 Factory 合约，设置当前地址为管理员
        factory = new UniswapV2Factory(address(this));
        // 部署 Router 合约，传入 Factory 和 WETH 地址
        router = new UniswapV2Router02(address(factory), address(weth));

        // 部署测试代币，初始供应量各 1,000,000
        tokenA = new MockTestERC20("Token A", "TKA", 1000000 * ONE_TOKEN);
        tokenB = new MockTestERC20("Token B", "TKB", 1000000 * ONE_TOKEN);
        feeToken = new FeeToken("Fee Token", "FEE", 1000000 * ONE_TOKEN);

        // 给测试账户分配代币
        tokenA.transfer(alice, 10000 * ONE_TOKEN); // 给 alice 分配 10000 个 tokenA
        tokenB.transfer(alice, 10000 * ONE_TOKEN); // 给 alice 分配 10000 个 tokenB
        tokenA.transfer(bob, 10000 * ONE_TOKEN); // 给 bob 分配 10000 个 tokenA
        tokenB.transfer(bob, 10000 * ONE_TOKEN); // 给 bob 分配 10000 个 tokenB
        tokenA.transfer(charlie, 10000 * ONE_TOKEN); // 给 charlie 分配 10000 个 tokenA
        feeToken.transfer(charlie, 10000 * ONE_TOKEN); // 给 charlie 分配 10000 个 feeToken

        // 给测试账户分配 ETH（用于 WETH 相关测试）
        vm.deal(alice, 100 ether); // 给 alice 分配 100 ETH
        vm.deal(bob, 100 ether); // 给 bob 分配 100 ETH
        vm.deal(charlie, 100 ether); // 给 charlie 分配 100 ETH
    }

    // ============ Factory 测试 ============

    /**
     * @notice 测试创建交易对
     * @dev 验证交易对创建成功，地址正确记录
     */
    function testCreatePair() public {
        // 创建 tokenA 和 tokenB 的交易对
        address pair = factory.createPair(address(tokenA), address(tokenB));
        // 验证交易对地址不为零
        assertTrue(pair != address(0));
        // 验证交易对数量为 1
        assertEq(factory.allPairsLength(), 1);
        // 验证可以通过 getPair 查询到交易对
        assertEq(factory.getPair(address(tokenA), address(tokenB)), pair);
    }

    /**
     * @notice 测试创建交易对时使用相同地址
     * @dev 应该失败，因为两个代币地址不能相同
     */
    function testCreatePairIdenticalAddresses() public {
        // 期望失败，错误信息为 "IDENTICAL_ADDRESSES"
        vm.expectRevert("UniswapV2: IDENTICAL_ADDRESSES");
        // 尝试用相同的地址创建交易对
        factory.createPair(address(tokenA), address(tokenA));
    }

    /**
     * @notice 测试创建交易对时使用零地址
     * @dev 应该失败，因为代币地址不能为零
     */
    function testCreatePairZeroAddress() public {
        // 期望失败，错误信息为 "ZERO_ADDRESS"
        vm.expectRevert("UniswapV2: ZERO_ADDRESS");
        // 尝试用零地址创建交易对
        factory.createPair(address(0), address(tokenB));
    }

    /**
     * @notice 测试创建已存在的交易对
     * @dev 应该失败，因为交易对已存在
     */
    function testCreatePairAlreadyExists() public {
        // 先创建一个交易对
        factory.createPair(address(tokenA), address(tokenB));
        // 期望失败，错误信息为 "PAIR_EXISTS"
        vm.expectRevert("UniswapV2: PAIR_EXISTS");
        // 尝试再次创建相同的交易对
        factory.createPair(address(tokenA), address(tokenB));
    }

    /**
     * @notice 测试设置费用接收地址
     * @dev 验证管理员可以设置费用接收地址
     */
    function testSetFeeTo() public {
        // 创建新的费用接收地址
        address newFeeTo = makeAddr("feeTo");
        // 设置费用接收地址
        factory.setFeeTo(newFeeTo);
        // 验证设置成功
        assertEq(factory.feeTo(), newFeeTo);
    }

    /**
     * @notice 测试非管理员设置费用接收地址
     * @dev 应该失败，因为只有管理员可以设置
     */
    function testSetFeeToUnauthorized() public {
        // 切换到 alice（非管理员）
        vm.prank(alice);
        // 期望失败，错误信息为 "FORBIDDEN"
        vm.expectRevert("UniswapV2: FORBIDDEN");
        // 尝试设置费用接收地址
        factory.setFeeTo(alice);
    }

    // ============ 流动性测试 ============

    /**
     * @notice 测试添加流动性
     * @dev 验证添加流动性成功，LP 代币正确铸造
     */
    function testAddLiquidity() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);
        tokenB.approve(address(router), type(uint256).max);

        // 添加流动性：各 1000 个代币
        (uint256 amountA, uint256 amountB, uint256 liquidity) = router.addLiquidity(
            address(tokenA), // 代币 A 地址
            address(tokenB), // 代币 B 地址
            1000 * ONE_TOKEN, // 期望添加的代币 A 数量
            1000 * ONE_TOKEN, // 期望添加的代币 B 数量
            0, // 最小代币 A 数量
            0, // 最小代币 B 数量
            alice, // LP 代币接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证添加的代币数量正确
        assertEq(amountA, 1000 * ONE_TOKEN);
        assertEq(amountB, 1000 * ONE_TOKEN);
        // 验证获得了 LP 代币
        assertTrue(liquidity > 0);

        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));
        // 验证 alice 持有 LP 代币
        assertEq(IERC20(pair).balanceOf(alice), liquidity);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试添加流动性时自动创建交易对
     * @dev 验证如果交易对不存在，会自动创建
     */
    function testAddLiquidityCreatesPair() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 验证当前没有交易对
        assertEq(factory.allPairsLength(), 0);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);
        tokenB.approve(address(router), type(uint256).max);

        // 添加流动性（会自动创建交易对）
        router.addLiquidity(
            address(tokenA),
            address(tokenB),
            1000 * ONE_TOKEN,
            1000 * ONE_TOKEN,
            0,
            0,
            alice,
            block.timestamp + 1 hours
        );

        // 验证交易对已创建
        assertEq(factory.allPairsLength(), 1);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试添加流动性时交易已过期
     * @dev 应该失败，因为截止时间已过
     */
    function testAddLiquidityExpired() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);
        tokenB.approve(address(router), type(uint256).max);

        // 期望失败，错误信息为 "EXPIRED"
        vm.expectRevert("UniswapV2Router: EXPIRED");
        // 尝试添加流动性，截止时间为过去的时间
        router.addLiquidity(
            address(tokenA),
            address(tokenB),
            1000 * ONE_TOKEN,
            1000 * ONE_TOKEN,
            0,
            0,
            alice,
            block.timestamp - 1 // 截止时间为过去
        );

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试添加流动性时代币 B 数量不足
     * @dev 应该失败，因为计算出的代币数量不满足最小值要求
     */
    function testAddLiquidityInsufficientBAmount() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);
        tokenB.approve(address(router), type(uint256).max);

        // 先添加一次流动性，建立价格比例
        router.addLiquidity(
            address(tokenA),
            address(tokenB),
            1000 * ONE_TOKEN,
            1000 * ONE_TOKEN,
            0,
            0,
            alice,
            block.timestamp + 1 hours
        );

        // 期望失败，错误信息为 "INSUFFICIENT_A_AMOUNT"
        // 因为价格比例是 1:1，但期望添加 1000 tokenA 只提供 100 tokenB
        vm.expectRevert("UniswapV2Router: INSUFFICIENT_A_AMOUNT");
        // 尝试添加流动性，但代币数量不符合价格比例
        router.addLiquidity(
            address(tokenA),
            address(tokenB),
            1000 * ONE_TOKEN, // 期望添加 1000 tokenA
            100 * ONE_TOKEN, // 但只愿意提供 100 tokenB
            900 * ONE_TOKEN, // 最小需要 900 tokenA
            0,
            alice,
            block.timestamp + 1 hours
        );

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试添加 ETH 流动性
     * @dev 验证可以添加 ETH 和代币的流动性
     */
    function testAddLiquidityETH() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 添加 ETH 流动性：1000 tokenA + 10 ETH
        (uint256 amountToken, uint256 amountETH, uint256 liquidity) = router.addLiquidityETH{value: 10 ether}(
            address(tokenA), // 代币地址
            1000 * ONE_TOKEN, // 期望添加的代币数量
            0, // 最小代币数量
            0, // 最小 ETH 数量
            alice, // LP 代币接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证添加的数量正确
        assertEq(amountToken, 1000 * ONE_TOKEN);
        assertEq(amountETH, 10 ether);
        // 验证获得了 LP 代币
        assertTrue(liquidity > 0);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试移除流动性
     * @dev 验证可以移除流动性并取回代币
     */
    function testRemoveLiquidity() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);
        tokenB.approve(address(router), type(uint256).max);

        // 先添加流动性
        (,, uint256 liquidity) = router.addLiquidity(
            address(tokenA),
            address(tokenB),
            1000 * ONE_TOKEN,
            1000 * ONE_TOKEN,
            0,
            0,
            alice,
            block.timestamp + 1 hours
        );

        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));
        // 授权 router 可以使用 alice 的 LP 代币
        IERC20(pair).approve(address(router), liquidity);

        // 记录移除前的代币余额
        uint256 aliceABefore = tokenA.balanceOf(alice);
        uint256 aliceBBefore = tokenB.balanceOf(alice);

        // 移除流动性
        router.removeLiquidity(
            address(tokenA),
            address(tokenB),
            liquidity, // 要销毁的 LP 代币数量
            0, // 最小获得的代币 A 数量
            0, // 最小获得的代币 B 数量
            alice, // 代币接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了代币
        assertTrue(tokenA.balanceOf(alice) > aliceABefore);
        assertTrue(tokenB.balanceOf(alice) > aliceBBefore);
        // 验证 LP 代币已销毁
        assertEq(IERC20(pair).balanceOf(alice), 0);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试移除 ETH 流动性
     * @dev 验证可以移除 ETH 流动性并取回 ETH
     */
    function testRemoveLiquidityETH() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 授权 router 可以使用 alice 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 先添加 ETH 流动性
        (,, uint256 liquidity) = router.addLiquidityETH{value: 10 ether}(
            address(tokenA),
            1000 * ONE_TOKEN,
            0,
            0,
            alice,
            block.timestamp + 1 hours
        );

        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(weth));
        // 授权 router 可以使用 alice 的 LP 代币
        IERC20(pair).approve(address(router), liquidity);

        // 记录移除前的余额
        uint256 aliceEthBefore = alice.balance;
        uint256 aliceTokenBefore = tokenA.balanceOf(alice);

        // 移除 ETH 流动性
        router.removeLiquidityETH(
            address(tokenA),
            liquidity,
            0,
            0,
            alice,
            block.timestamp + 1 hours
        );

        // 验证获得了 ETH 和代币
        assertTrue(alice.balance > aliceEthBefore);
        assertTrue(tokenA.balanceOf(alice) > aliceTokenBefore);

        // 停止模拟 alice
        vm.stopPrank();
    }

    // ============ 交换测试 ============

    /**
     * @notice 测试精确输入交换
     * @dev 验证可以指定精确的输入金额进行交换
     */
    function testSwapExactTokensForTokens() public {
        // 先添加流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);

        // 切换到 bob 账户
        vm.startPrank(bob);
        // 授权 router 可以使用 bob 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 记录交换前的代币余额
        uint256 bobBBefore = tokenB.balanceOf(bob);

        // 设置交换路径：tokenA -> tokenB
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 执行交换：用 100 tokenA 换 tokenB
        uint256[] memory amounts = router.swapExactTokensForTokens(
            100 * ONE_TOKEN, // 输入金额
            0, // 最小输出金额
            path, // 交换路径
            bob, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了 tokenB
        assertTrue(tokenB.balanceOf(bob) > bobBBefore);
        // 验证返回的金额数组长度正确
        assertEq(amounts.length, 2);

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试精确输出交换
     * @dev 验证可以指定精确的输出金额进行交换
     */
    function testSwapTokensForExactTokens() public {
        // 先添加流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);

        // 切换到 bob 账户
        vm.startPrank(bob);
        // 授权 router 可以使用 bob 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 记录交换前的代币余额
        uint256 bobBBefore = tokenB.balanceOf(bob);
        // 设置期望获得的 tokenB 数量
        uint256 desiredOut = 50 * ONE_TOKEN;

        // 设置交换路径：tokenA -> tokenB
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 执行交换：用 tokenA 换精确的 50 tokenB
        uint256[] memory amounts = router.swapTokensForExactTokens(
            desiredOut, // 期望的输出金额
            type(uint256).max, // 最大输入金额
            path, // 交换路径
            bob, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了精确的 tokenB 数量
        assertEq(tokenB.balanceOf(bob) - bobBBefore, desiredOut);
        // 验证输入金额大于 0
        assertTrue(amounts[0] > 0);

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试用精确 ETH 交换代币
     * @dev 验证可以用 ETH 交换代币
     */
    function testSwapExactETHForTokens() public {
        // 先添加 ETH 流动性
        _addLiquidityETH(alice, 1000 * ONE_TOKEN, 10 ether);

        // 切换到 bob 账户
        vm.startPrank(bob);

        // 记录交换前的代币余额
        uint256 bobTokenBefore = tokenA.balanceOf(bob);

        // 设置交换路径：WETH -> tokenA
        address[] memory path = new address[](2);
        path[0] = address(weth);
        path[1] = address(tokenA);

        // 执行交换：用 1 ETH 换 tokenA
        router.swapExactETHForTokens{value: 1 ether}(
            0, // 最小输出金额
            path, // 交换路径
            bob, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了 tokenA
        assertTrue(tokenA.balanceOf(bob) > bobTokenBefore);

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试用代币交换精确 ETH
     * @dev 验证可以用代币交换精确的 ETH
     */
    function testSwapTokensForExactETH() public {
        // 先添加 ETH 流动性
        _addLiquidityETH(alice, 1000 * ONE_TOKEN, 10 ether);

        // 切换到 bob 账户
        vm.startPrank(bob);
        // 授权 router 可以使用 bob 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 记录交换前的 ETH 余额
        uint256 bobEthBefore = bob.balance;
        // 设置期望获得的 ETH 数量
        uint256 desiredEthOut = 1 ether;

        // 设置交换路径：tokenA -> WETH
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(weth);

        // 执行交换：用 tokenA 换精确的 1 ETH
        router.swapTokensForExactETH(
            desiredEthOut, // 期望的 ETH 输出金额
            type(uint256).max, // 最大代币输入金额
            path, // 交换路径
            bob, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了精确的 ETH 数量
        assertEq(bob.balance - bobEthBefore, desiredEthOut);

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试用精确代币交换 ETH
     * @dev 验证可以用精确的代币交换 ETH
     */
    function testSwapExactTokensForETH() public {
        // 先添加 ETH 流动性
        _addLiquidityETH(alice, 1000 * ONE_TOKEN, 10 ether);

        // 切换到 bob 账户
        vm.startPrank(bob);
        // 授权 router 可以使用 bob 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 记录交换前的 ETH 余额
        uint256 bobEthBefore = bob.balance;

        // 设置交换路径：tokenA -> WETH
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(weth);

        // 执行交换：用 100 tokenA 换 ETH
        router.swapExactTokensForETH(
            100 * ONE_TOKEN, // 输入代币金额
            0, // 最小 ETH 输出金额
            path, // 交换路径
            bob, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了 ETH
        assertTrue(bob.balance > bobEthBefore);

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试交换时交易已过期
     * @dev 应该失败，因为截止时间已过
     */
    function testSwapExpired() public {
        // 先添加流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);

        // 切换到 bob 账户
        vm.startPrank(bob);
        // 授权 router 可以使用 bob 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 设置交换路径
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 期望失败，错误信息为 "EXPIRED"
        vm.expectRevert("UniswapV2Router: EXPIRED");
        // 尝试交换，截止时间为过去的时间
        router.swapExactTokensForTokens(
            100 * ONE_TOKEN,
            0,
            path,
            bob,
            block.timestamp - 1 // 截止时间为过去
        );

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试交换时输出不足
     * @dev 应该失败，因为输出金额不满足最小值要求
     */
    function testSwapInsufficientOutput() public {
        // 先添加流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);

        // 切换到 bob 账户
        vm.startPrank(bob);
        // 授权 router 可以使用 bob 的代币
        tokenA.approve(address(router), type(uint256).max);

        // 设置交换路径
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 期望失败，错误信息为 "INSUFFICIENT_OUTPUT_AMOUNT"
        vm.expectRevert("UniswapV2Router: INSUFFICIENT_OUTPUT_AMOUNT");
        // 尝试交换，但最小输出金额设置过高
        router.swapExactTokensForTokens(
            100 * ONE_TOKEN, // 输入 100 tokenA
            1000 * ONE_TOKEN, // 要求至少输出 1000 tokenB（不可能满足）
            path,
            bob,
            block.timestamp + 1 hours
        );

        // 停止模拟 bob
        vm.stopPrank();
    }

    /**
     * @notice 测试交换时路径无效
     * @dev 应该失败，因为 ETH 交换路径的第一个代币必须是 WETH
     */
    function testSwapInvalidPath() public {
        // 切换到 bob 账户
        vm.startPrank(bob);

        // 设置交换路径：tokenA -> tokenB（不是 WETH 开头）
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 期望失败，错误信息为 "INVALID_PATH"
        vm.expectRevert("UniswapV2Router: INVALID_PATH");
        // 尝试用 ETH 交换，但路径不是以 WETH 开头
        router.swapExactETHForTokens{value: 1 ether}(
            0,
            path,
            bob,
            block.timestamp + 1 hours
        );

        // 停止模拟 bob
        vm.stopPrank();
    }

    // ============ 手续费代币测试 ============

    /**
     * @notice 测试手续费代币交换
     * @dev 验证支持手续费代币的交换函数可以正确处理有转账手续费的代币
     */
    function testSwapWithFeeToken() public {
        // 切换到 charlie 账户
        vm.startPrank(charlie);

        // 授权 router 可以使用 charlie 的代币
        feeToken.approve(address(router), type(uint256).max);
        tokenA.approve(address(router), type(uint256).max);

        // 添加流动性：feeToken + tokenA
        router.addLiquidity(
            address(feeToken),
            address(tokenA),
            1000 * ONE_TOKEN,
            1000 * ONE_TOKEN,
            0,
            0,
            charlie,
            block.timestamp + 1 hours
        );

        // 停止模拟 charlie
        vm.stopPrank();

        // 再次切换到 charlie 账户
        vm.startPrank(charlie);
        // 授权 router 可以使用 charlie 的 feeToken
        feeToken.approve(address(router), type(uint256).max);

        // 记录交换前的 tokenA 余额
        uint256 charlieTokenABefore = tokenA.balanceOf(charlie);

        // 设置交换路径：feeToken -> tokenA
        address[] memory path = new address[](2);
        path[0] = address(feeToken);
        path[1] = address(tokenA);

        // 使用支持手续费代币的交换函数
        router.swapExactTokensForTokensSupportingFeeOnTransferTokens(
            100 * ONE_TOKEN, // 输入金额
            0, // 最小输出金额
            path, // 交换路径
            charlie, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );

        // 验证获得了 tokenA
        assertTrue(tokenA.balanceOf(charlie) > charlieTokenABefore);

        // 停止模拟 charlie
        vm.stopPrank();
    }

    // ============ 价格计算测试 ============

    /**
     * @notice 测试计算输出金额
     * @dev 验证可以根据输入金额计算输出金额
     */
    function testGetAmountsOut() public {
        // 先添加流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);

        // 设置交换路径
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 计算输出金额
        uint256[] memory amounts = router.getAmountsOut(100 * ONE_TOKEN, path);

        // 验证返回的金额数组长度正确
        assertEq(amounts.length, 2);
        // 验证输入金额正确
        assertEq(amounts[0], 100 * ONE_TOKEN);
        // 验证输出金额大于 0
        assertTrue(amounts[1] > 0);
        // 验证输出金额小于输入金额（因为有手续费和价格影响）
        assertTrue(amounts[1] < 100 * ONE_TOKEN);
    }

    /**
     * @notice 测试计算输入金额
     * @dev 验证可以根据输出金额计算需要的输入金额
     */
    function testGetAmountsIn() public {
        // 先添加流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);

        // 设置交换路径
        address[] memory path = new address[](2);
        path[0] = address(tokenA);
        path[1] = address(tokenB);

        // 计算需要的输入金额
        uint256[] memory amounts = router.getAmountsIn(50 * ONE_TOKEN, path);

        // 验证返回的金额数组长度正确
        assertEq(amounts.length, 2);
        // 验证输出金额正确
        assertEq(amounts[1], 50 * ONE_TOKEN);
        // 验证输入金额大于 0
        assertTrue(amounts[0] > 0);
        // 验证输入金额小于输出金额的两倍（因为有手续费）
        assertTrue(amounts[0] < 100 * ONE_TOKEN);
    }

    /**
     * @notice 测试计算输出金额时路径无效
     * @dev 应该失败，因为路径长度不足
     */
    function testGetAmountsOutInvalidPath() public {
        // 设置只有 1 个代币的路径
        address[] memory path = new address[](1);
        path[0] = address(tokenA);

        // 期望失败，错误信息为 "INVALID_PATH"
        vm.expectRevert("UniswapV2Router: INVALID_PATH");
        // 尝试计算输出金额
        router.getAmountsOut(100 * ONE_TOKEN, path);
    }

    // ============ Pair 测试 ============

    /**
     * @notice 测试 Pair 铸造 LP 代币
     * @dev 验证可以直接调用 Pair 的 mint 函数
     */
    function testPairMint() public {
        // 创建交易对
        factory.createPair(address(tokenA), address(tokenB));
        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));

        // 切换到 alice 账户
        vm.startPrank(alice);

        // 将代币转入交易对
        tokenA.transfer(pair, 1000 * ONE_TOKEN);
        tokenB.transfer(pair, 1000 * ONE_TOKEN);

        // 调用 mint 函数铸造 LP 代币
        uint256 liquidity = UniswapV2Pair(pair).mint(alice);

        // 验证获得了 LP 代币
        assertTrue(liquidity > 0);
        // 验证 alice 持有 LP 代币
        assertEq(IERC20(pair).balanceOf(alice), liquidity);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试 Pair 销毁 LP 代币
     * @dev 验证可以直接调用 Pair 的 burn 函数
     */
    function testPairBurn() public {
        // 创建交易对
        factory.createPair(address(tokenA), address(tokenB));
        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));

        // 切换到 alice 账户
        vm.startPrank(alice);

        // 将代币转入交易对
        tokenA.transfer(pair, 1000 * ONE_TOKEN);
        tokenB.transfer(pair, 1000 * ONE_TOKEN);

        // 铸造 LP 代币
        uint256 liquidity = UniswapV2Pair(pair).mint(alice);

        // 将 LP 代币转回交易对
        IERC20(pair).transfer(pair, liquidity);

        // 调用 burn 函数销毁 LP 代币并取回代币
        (uint256 amount0, uint256 amount1) = UniswapV2Pair(pair).burn(alice);

        // 验证取回了代币
        assertTrue(amount0 > 0);
        assertTrue(amount1 > 0);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试 Pair 交换
     * @dev 验证可以直接调用 Pair 的 swap 函数
     */
    function testPairSwap() public {
        // 创建交易对
        factory.createPair(address(tokenA), address(tokenB));
        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));

        // 切换到 alice 账户
        vm.startPrank(alice);

        // 将代币转入交易对并铸造 LP 代币
        tokenA.transfer(pair, 1000 * ONE_TOKEN);
        tokenB.transfer(pair, 1000 * ONE_TOKEN);
        UniswapV2Pair(pair).mint(alice);

        // 记录 bob 的 tokenB 余额
        uint256 bobBBefore = tokenB.balanceOf(bob);

        // 将 tokenA 转入交易对（作为交换输入）
        tokenA.transfer(pair, 100 * ONE_TOKEN);

        // 调用 swap 函数执行交换
        UniswapV2Pair(pair).swap(0, 90 * ONE_TOKEN, bob, "");

        // 验证 bob 获得了 tokenB
        assertTrue(tokenB.balanceOf(bob) > bobBBefore);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试 Pair 同步
     * @dev 验证 sync 函数可以更新储备量
     */
    function testPairSync() public {
        // 创建交易对
        factory.createPair(address(tokenA), address(tokenB));
        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));

        // 切换到 alice 账户
        vm.startPrank(alice);

        // 将代币转入交易对并铸造 LP 代币
        tokenA.transfer(pair, 1000 * ONE_TOKEN);
        tokenB.transfer(pair, 1000 * ONE_TOKEN);
        UniswapV2Pair(pair).mint(alice);

        // 额外转入 tokenA（不通过正常的添加流动性方式）
        tokenA.transfer(pair, 100 * ONE_TOKEN);

        // 调用 sync 函数同步储备量
        UniswapV2Pair(pair).sync();

        // 验证储备量已更新
        (uint112 reserve0,,) = UniswapV2Pair(pair).getReserves();
        assertEq(reserve0, 1100 * ONE_TOKEN);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试 Pair 清理
     * @dev 验证 skim 函数可以取回多余的代币
     */
    function testPairSkim() public {
        // 创建交易对
        factory.createPair(address(tokenA), address(tokenB));
        // 获取交易对地址
        address pair = factory.getPair(address(tokenA), address(tokenB));

        // 切换到 alice 账户
        vm.startPrank(alice);

        // 将代币转入交易对并铸造 LP 代币
        tokenA.transfer(pair, 1000 * ONE_TOKEN);
        tokenB.transfer(pair, 1000 * ONE_TOKEN);
        UniswapV2Pair(pair).mint(alice);

        // 额外转入 tokenA
        tokenA.transfer(pair, 100 * ONE_TOKEN);

        // 记录 alice 的 tokenA 余额
        uint256 aliceABefore = tokenA.balanceOf(alice);

        // 调用 skim 函数取回多余的代币
        UniswapV2Pair(pair).skim(alice);

        // 验证 alice 获得了多余的 tokenA
        assertTrue(tokenA.balanceOf(alice) > aliceABefore);

        // 停止模拟 alice
        vm.stopPrank();
    }

    // ============ WETH 测试 ============

    /**
     * @notice 测试 WETH 存款
     * @dev 验证可以将 ETH 包装成 WETH
     */
    function testWETHDeposit() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 存款金额
        uint256 amount = 1 ether;

        // 调用 deposit 函数将 ETH 包装成 WETH
        weth.deposit{value: amount}();

        // 验证 alice 持有 WETH
        assertEq(weth.balanceOf(alice), amount);
        // 验证 WETH 合约持有 ETH
        assertEq(address(weth).balance, amount);

        // 停止模拟 alice
        vm.stopPrank();
    }

    /**
     * @notice 测试 WETH 取款
     * @dev 验证可以将 WETH 解包成 ETH
     */
    function testWETHWithdraw() public {
        // 切换到 alice 账户
        vm.startPrank(alice);

        // 存款金额
        uint256 amount = 1 ether;

        // 先存款
        weth.deposit{value: amount}();

        // 记录取款前的 ETH 余额
        uint256 ethBefore = alice.balance;

        // 调用 withdraw 函数将 WETH 解包成 ETH
        weth.withdraw(amount);

        // 验证 alice 不再持有 WETH
        assertEq(weth.balanceOf(alice), 0);
        // 验证 alice 获得了 ETH
        assertEq(alice.balance, ethBefore + amount);

        // 停止模拟 alice
        vm.stopPrank();
    }

    // ============ 多跳交换测试 ============

    /**
     * @notice 测试多跳交换
     * @dev 验证可以通过多个交易对进行交换
     */
    function testMultiHopSwap() public {
        // 部署第三个测试代币用于多跳交换
        MockTestERC20 tokenC = new MockTestERC20("Token C", "TKC", 1000000 * ONE_TOKEN);
        
        // 给 alice 分配 tokenC
        tokenC.transfer(alice, 10000 * ONE_TOKEN);
        
        // 首先添加 tokenA 和 tokenB 的流动性
        _addLiquidity(alice, 1000 * ONE_TOKEN, 1000 * ONE_TOKEN);
        
        // 添加 tokenB 和 tokenC 的流动性
        vm.startPrank(alice);
        tokenB.approve(address(router), type(uint256).max);
        tokenC.approve(address(router), type(uint256).max);
        router.addLiquidity(
            address(tokenB),
            address(tokenC),
            1000 * ONE_TOKEN,
            1000 * ONE_TOKEN,
            0,
            0,
            alice,
            block.timestamp + 1 hours
        );
        vm.stopPrank();
        
        // 切换到 bob 账户进行多跳交换
        vm.startPrank(bob);
        tokenA.approve(address(router), type(uint256).max);
        
        // 记录交换前的 tokenC 余额
        uint256 bobTokenCBefore = tokenC.balanceOf(bob);
        
        // 设置多跳交换路径：tokenA -> tokenB -> tokenC
        address[] memory path = new address[](3);
        path[0] = address(tokenA);
        path[1] = address(tokenB);
        path[2] = address(tokenC);
        
        // 执行多跳交换：用 100 tokenA 换 tokenC
        uint256[] memory amounts = router.swapExactTokensForTokens(
            100 * ONE_TOKEN, // 输入金额
            0, // 最小输出金额
            path, // 多跳路径
            bob, // 接收地址
            block.timestamp + 1 hours // 截止时间
        );
        
        // 验证获得了 tokenC
        assertTrue(tokenC.balanceOf(bob) > bobTokenCBefore);
        // 验证返回的金额数组长度正确
        assertEq(amounts.length, 3);
        // 验证输入金额正确
        assertEq(amounts[0], 100 * ONE_TOKEN);
        // 验证中间金额和输出金额都大于 0
        assertTrue(amounts[1] > 0);
        assertTrue(amounts[2] > 0);
        
        // 停止模拟 bob
        vm.stopPrank();
    }

    // ============ 辅助函数 ============

    /**
     * @notice 辅助函数：添加流动性
     * @param user 添加流动性的账户
     * @param amountA 代币 A 数量
     * @param amountB 代币 B 数量
     */
    function _addLiquidity(address user, uint256 amountA, uint256 amountB) internal {
        // 切换到指定账户
        vm.startPrank(user);
        // 授权 router 可以使用代币
        tokenA.approve(address(router), type(uint256).max);
        tokenB.approve(address(router), type(uint256).max);
        // 添加流动性
        router.addLiquidity(
            address(tokenA),
            address(tokenB),
            amountA,
            amountB,
            0,
            0,
            user,
            block.timestamp + 1 hours
        );
        // 停止模拟
        vm.stopPrank();
    }

    /**
     * @notice 辅助函数：添加 ETH 流动性
     * @param user 添加流动性的账户
     * @param amountToken 代币数量
     * @param amountETH ETH 数量
     */
    function _addLiquidityETH(address user, uint256 amountToken, uint256 amountETH) internal {
        // 切换到指定账户
        vm.startPrank(user);
        // 授权 router 可以使用代币
        tokenA.approve(address(router), type(uint256).max);
        // 添加 ETH 流动性
        router.addLiquidityETH{value: amountETH}(
            address(tokenA),
            amountToken,
            0,
            0,
            user,
            block.timestamp + 1 hours
        );
        // 停止模拟
        vm.stopPrank();
    }
}
