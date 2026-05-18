// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

import "@openzeppelin/contracts/token/ERC20/IERC20.sol"; // 导入 OpenZeppelin 的 ERC20 接口
import "./UniswapV2Pair.sol"; // 导入交易对合约
import "./libraries/Math.sol"; // 导入数学库
import "./libraries/UniswapV2Library.sol"; // 导入 UniswapV2 工具库
import "./interfaces/IWETH.sol"; // 导入 WETH 接口
import "./interfaces/IUniswapV2Router02.sol"; // 导入路由接口

/**
 * @title UniswapV2Router02 路由合约
 * @notice 用户与 UniswapV2 交互的主要入口
 * @dev 提供添加/移除流动性、代币交换等便捷功能
 *      Router 是用户与 UniswapV2 交互的主要接口
 *      它封装了与 Factory 和 Pair 的交互逻辑
 *      提供了更友好的 API，包括：
 *      - 自动计算最优添加流动性数量
 *      - 支持多跳交易
 *      - 支持 ETH 与代币的交易
 *      - 支持手续费代币
 */
contract UniswapV2Router02 is IUniswapV2Router02 {
    using Math for uint256; // 为 uint256 类型启用 Math 库的方法

    // 工厂合约地址：不可变变量，部署后不可修改
    address public immutable factory;
    // WETH 合约地址：不可变变量，部署后不可修改
    address public immutable weth;

    /**
     * @notice 构造函数
     * @param _factory 工厂合约地址
     * @param _weth WETH 合约地址
     * @dev 初始化工厂和 WETH 地址
     */
    constructor(address _factory, address _weth) {
        factory = _factory; // 设置工厂合约地址
        weth = _weth; // 设置 WETH 合约地址
    }

    /**
     * @notice 接收 ETH 的回调函数
     * @dev 只有 WETH 合约可以向此合约发送 ETH
     *      当 WETH 解包时，会将 ETH 发送到此合约
     */
    receive() external payable {
        // 确保只有 WETH 合约可以发送 ETH
        require(msg.sender == weth, "UniswapV2Router: INVALID_SENDER");
    }

    // ============ 内部辅助函数 ============

    /**
     * @notice 安全转账 ETH
     * @param to 接收地址
     * @param amount 转账金额
     * @dev 使用 call 方法发送 ETH，这是最安全的发送方式
     */
    function _safeTransferETH(address to, uint256 amount) internal {
        // 使用 call 发送 ETH
        (bool success,) = to.call{value: amount}("");
        // 确保发送成功
        require(success, "UniswapV2Router: ETH_TRANSFER_FAILED");
    }

    /**
     * @notice 确保转账成功（从调用者账户转账）
     * @param token 代币地址
     * @param to 接收地址
     * @param amount 转账金额
     * @dev 使用 transferFrom 从调用者账户转账代币
     *      0x23b872dd 是 transferFrom(address,address,uint256) 的选择器
     */
    function _safeTransfer(address token, address to, uint256 amount) internal {
        // 调用代币的 transferFrom 函数
        // 0x23b872dd = transferFrom(address,address,uint256) 的选择器
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x23b872dd, msg.sender, to, amount));
        // 检查调用是否成功，以及返回值是否为 true（或没有返回值）
        require(success && (data.length == 0 || abi.decode(data, (bool))), "UniswapV2Router: TRANSFER_FAILED");
    }

    /**
     * @notice 确保转账成功（从合约自身账户转账）
     * @param token 代币地址
     * @param to 接收地址
     * @param amount 转账金额
     * @dev 使用 transfer 从合约自身账户转账代币
     *      0xa9059cbb 是 transfer(address,uint256) 的选择器
     */
    function _safeTransferOwn(address token, address to, uint256 amount) internal {
        // 调用代币的 transfer 函数
        // 0xa9059cbb = transfer(address,uint256) 的选择器
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0xa9059cbb, to, amount));
        // 检查调用是否成功，以及返回值是否为 true（或没有返回值）
        require(success && (data.length == 0 || abi.decode(data, (bool))), "UniswapV2Router: TRANSFER_FAILED");
    }

    // ============ 流动性管理 ============

    /**
     * @notice 添加流动性
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @param amountADesired 期望添加的代币A数量
     * @param amountBDesired 期望添加的代币B数量
     * @param amountAMin 最小代币A数量（滑点保护）
     * @param amountBMin 最小代币B数量（滑点保护）
     * @param to LP 代币接收地址
     * @param deadline 截止时间（交易必须在此时间前完成）
     * @return amountA 实际添加的代币A数量
     * @return amountB 实际添加的代币B数量
     * @return liquidity 获得的流动性数量
     * @dev 自动计算最优添加数量，确保按当前价格比例添加
     */
    function addLiquidity(
        address tokenA,
        address tokenB,
        uint256 amountADesired,
        uint256 amountBDesired,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountA, uint256 amountB, uint256 liquidity) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");

        // 检查交易对是否存在，不存在则创建
        address pair = IUniswapV2Factory(factory).getPair(tokenA, tokenB);
        if (pair == address(0)) {
            // 如果交易对不存在，创建新的交易对
            pair = IUniswapV2Factory(factory).createPair(tokenA, tokenB);
        }

        // 获取当前储备量
        (uint256 reserveA, uint256 reserveB) = UniswapV2Library.getReserves(factory, tokenA, tokenB);

        if (reserveA == 0 && reserveB == 0) {
            // 如果是新交易对，使用用户期望的数量
            (amountA, amountB) = (amountADesired, amountBDesired);
        } else {
            // 如果交易对已存在，计算最优添加数量
            // 根据 amountADesired 计算需要的 amountB
            uint256 amountBOptimal = amountADesired * reserveB / reserveA;
            if (amountBOptimal <= amountBDesired) {
                // 如果计算出的 amountB 在期望范围内
                require(amountBOptimal >= amountBMin, "UniswapV2Router: INSUFFICIENT_B_AMOUNT");
                (amountA, amountB) = (amountADesired, amountBOptimal);
            } else {
                // 如果计算出的 amountB 超出期望范围，反向计算 amountA
                uint256 amountAOptimal = amountBDesired * reserveA / reserveB;
                require(amountAOptimal >= amountAMin, "UniswapV2Router: INSUFFICIENT_A_AMOUNT");
                (amountA, amountB) = (amountAOptimal, amountBDesired);
            }
        }

        // 将代币从用户账户转入交易对
        _safeTransfer(tokenA, pair, amountA);
        _safeTransfer(tokenB, pair, amountB);
        // 铸造 LP 代币给用户
        liquidity = IUniswapV2Pair(pair).mint(to);
    }

    /**
     * @notice 添加 ETH 流动性
     * @param token 代币地址
     * @param amountTokenDesired 期望添加的代币数量
     * @param amountTokenMin 最小代币数量（滑点保护）
     * @param amountETHMin 最小 ETH 数量（滑点保护）
     * @param to LP 代币接收地址
     * @param deadline 截止时间
     * @dev 用户发送 ETH，合约自动将其包装成 WETH
     */
    function addLiquidityETH(
        address token,
        uint256 amountTokenDesired,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external payable returns (uint256 amountToken, uint256 amountETH, uint256 liquidity) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");

        // 检查交易对是否存在，不存在则创建
        address pair = IUniswapV2Factory(factory).getPair(token, weth);
        if (pair == address(0)) {
            // 如果交易对不存在，创建新的交易对
            pair = IUniswapV2Factory(factory).createPair(token, weth);
        }

        // 获取当前储备量
        (uint256 reserveToken, uint256 reserveETH) = UniswapV2Library.getReserves(factory, token, weth);

        if (reserveToken == 0 && reserveETH == 0) {
            // 如果是新交易对，使用用户期望的数量
            (amountToken, amountETH) = (amountTokenDesired, msg.value);
        } else {
            // 如果交易对已存在，计算最优添加数量
            // 根据 amountTokenDesired 计算需要的 ETH
            uint256 amountETHOptimal = amountTokenDesired * reserveETH / reserveToken;
            if (amountETHOptimal <= msg.value) {
                // 如果计算出的 ETH 在发送范围内
                require(amountETHOptimal >= amountETHMin, "UniswapV2Router: INSUFFICIENT_ETH_AMOUNT");
                (amountToken, amountETH) = (amountTokenDesired, amountETHOptimal);
            } else {
                // 如果计算出的 ETH 超出发送范围，反向计算 token 数量
                uint256 amountTokenOptimal = msg.value * reserveToken / reserveETH;
                require(amountTokenOptimal >= amountTokenMin, "UniswapV2Router: INSUFFICIENT_TOKEN_AMOUNT");
                (amountToken, amountETH) = (amountTokenOptimal, msg.value);
            }
        }

        // 将代币从用户账户转入交易对
        _safeTransfer(token, pair, amountToken);
        // 将 ETH 包装成 WETH
        IWETH(weth).deposit{value: amountETH}();
        // 将 WETH 转入交易对
        _safeTransferOwn(weth, pair, amountETH);
        // 铸造 LP 代币给用户
        liquidity = IUniswapV2Pair(pair).mint(to);

        // 如果发送的 ETH 多于实际需要的，退还多余部分
        if (msg.value > amountETH) _safeTransferETH(msg.sender, msg.value - amountETH);
    }

    /**
     * @notice 移除流动性
     * @param tokenA 代币A地址
     * @param tokenB 代币B地址
     * @param liquidity 要销毁的流动性数量
     * @param amountAMin 最小获得的代币A数量（滑点保护）
     * @param amountBMin 最小获得的代币B数量（滑点保护）
     * @param to 代币接收地址
     * @param deadline 截止时间
     * @return amountA 获得的代币A数量
     * @return amountB 获得的代币B数量
     * @dev public 可见性，因为其他函数需要调用
     */
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint256 liquidity,
        uint256 amountAMin,
        uint256 amountBMin,
        address to,
        uint256 deadline
    ) public returns (uint256 amountA, uint256 amountB) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 计算交易对地址
        address pair = UniswapV2Library.pairFor(factory, tokenA, tokenB);
        // 将 LP 代币从用户账户转入交易对
        _safeTransfer(pair, pair, liquidity);
        // 销毁 LP 代币并获取代币
        (uint256 amount0, uint256 amount1) = IUniswapV2Pair(pair).burn(to);
        // 获取代币顺序
        (address token0,) = UniswapV2Library.sortTokens(tokenA, tokenB);
        // 根据代币顺序确定返回的金额
        (amountA, amountB) = tokenA == token0 ? (amount0, amount1) : (amount1, amount0);
        // 检查获得的代币数量是否满足最小值要求
        require(amountA >= amountAMin, "UniswapV2Router: INSUFFICIENT_A_AMOUNT");
        require(amountB >= amountBMin, "UniswapV2Router: INSUFFICIENT_B_AMOUNT");
    }

    /**
     * @notice 移除 ETH 流动性
     * @param token 代币地址
     * @param liquidity 要销毁的流动性数量
     * @param amountTokenMin 最小获得的代币数量（滑点保护）
     * @param amountETHMin 最小获得的 ETH 数量（滑点保护）
     * @param to 代币和 ETH 接收地址
     * @param deadline 截止时间
     * @dev 移除流动性后会自动将 WETH 解包成 ETH
     */
    function removeLiquidityETH(
        address token,
        uint256 liquidity,
        uint256 amountTokenMin,
        uint256 amountETHMin,
        address to,
        uint256 deadline
    ) external returns (uint256 amountToken, uint256 amountETH) {
        // 调用 removeLiquidity 移除流动性
        // 注意：WETH 发送到路由合约，而不是用户
        (amountToken, amountETH) = removeLiquidity(
            token,
            weth,
            liquidity,
            amountTokenMin,
            amountETHMin,
            address(this), // WETH 发送到路由合约
            deadline
        );
        // 将代币发送给用户
        _safeTransferOwn(token, to, amountToken);
        // 将 WETH 解包成 ETH
        IWETH(weth).withdraw(amountETH);
        // 将 ETH 发送给用户
        _safeTransferETH(to, amountETH);
    }

    // ============ 交易功能 ============

    /**
     * @notice 精确输入交换
     * @param amountIn 精确的输入金额
     * @param amountOutMin 最小输出金额（滑点保护）
     * @param path 交易路径（代币地址数组）
     * @param to 接收地址
     * @param deadline 截止时间
     * @return amounts 金额数组
     * @dev 用户指定精确的输入金额，获得尽可能多的输出
     */
    function swapExactTokensForTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 计算每一步的输出金额
        amounts = getAmountsOut(amountIn, path);
        // 检查最终输出是否满足最小值要求
        require(amounts[amounts.length - 1] >= amountOutMin, "UniswapV2Router: INSUFFICIENT_OUTPUT_AMOUNT");
        // 将输入代币从用户账户转入第一个交易对
        _safeTransfer(path[0], UniswapV2Library.pairFor(factory, path[0], path[1]), amounts[0]);
        // 执行交换
        _swap(amounts, path, to);
    }

    /**
     * @notice 精确输出交换
     * @param amountOut 精确的输出金额
     * @param amountInMax 最大输入金额（滑点保护）
     * @param path 交易路径（代币地址数组）
     * @param to 接收地址
     * @param deadline 截止时间
     * @return amounts 金额数组
     * @dev 用户指定精确的输出金额，支付尽可能少的输入
     */
    function swapTokensForExactTokens(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 计算每一步需要的输入金额
        amounts = getAmountsIn(amountOut, path);
        // 检查初始输入是否超过最大值
        require(amounts[0] <= amountInMax, "UniswapV2Router: EXCESSIVE_INPUT_AMOUNT");
        // 将输入代币从用户账户转入第一个交易对
        _safeTransfer(path[0], UniswapV2Library.pairFor(factory, path[0], path[1]), amounts[0]);
        // 执行交换
        _swap(amounts, path, to);
    }

    /**
     * @notice 用精确 ETH 交换代币
     * @param amountOutMin 最小输出金额（滑点保护）
     * @param path 交易路径（第一个必须是 WETH）
     * @param to 接收地址
     * @param deadline 截止时间
     * @dev 用户发送精确的 ETH，获得代币
     */
    function swapExactETHForTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable returns (uint256[] memory amounts) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 确保路径的第一个代币是 WETH
        require(path[0] == weth, "UniswapV2Router: INVALID_PATH");
        // 计算每一步的输出金额
        amounts = getAmountsOut(msg.value, path);
        // 检查最终输出是否满足最小值要求
        require(amounts[amounts.length - 1] >= amountOutMin, "UniswapV2Router: INSUFFICIENT_OUTPUT_AMOUNT");
        // 将 ETH 包装成 WETH
        IWETH(weth).deposit{value: amounts[0]}();
        // 将 WETH 转入第一个交易对
        _safeTransferOwn(weth, UniswapV2Library.pairFor(factory, path[0], path[1]), amounts[0]);
        // 执行交换
        _swap(amounts, path, to);
    }

    /**
     * @notice 用代币交换精确 ETH
     * @param amountOut 精确的 ETH 输出金额
     * @param amountInMax 最大代币输入金额（滑点保护）
     * @param path 交易路径（最后一个必须是 WETH）
     * @param to 接收地址
     * @param deadline 截止时间
     * @dev 用户指定精确的 ETH 输出，支付代币
     */
    function swapTokensForExactETH(
        uint256 amountOut,
        uint256 amountInMax,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 确保路径的最后一个代币是 WETH
        require(path[path.length - 1] == weth, "UniswapV2Router: INVALID_PATH");
        // 计算每一步需要的输入金额
        amounts = getAmountsIn(amountOut, path);
        // 检查初始输入是否超过最大值
        require(amounts[0] <= amountInMax, "UniswapV2Router: EXCESSIVE_INPUT_AMOUNT");
        // 将输入代币从用户账户转入第一个交易对
        _safeTransfer(path[0], UniswapV2Library.pairFor(factory, path[0], path[1]), amounts[0]);
        // 执行交换，WETH 发送到路由合约
        _swap(amounts, path, address(this));
        // 将 WETH 解包成 ETH
        IWETH(weth).withdraw(amounts[amounts.length - 1]);
        // 将 ETH 发送给用户
        _safeTransferETH(to, amounts[amounts.length - 1]);
    }

    /**
     * @notice 用精确代币交换 ETH
     * @param amountIn 精确的代币输入金额
     * @param amountOutMin 最小 ETH 输出金额（滑点保护）
     * @param path 交易路径（最后一个必须是 WETH）
     * @param to 接收地址
     * @param deadline 截止时间
     * @dev 用户指定精确的代币输入，获得 ETH
     */
    function swapExactTokensForETH(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external returns (uint256[] memory amounts) {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 确保路径的最后一个代币是 WETH
        require(path[path.length - 1] == weth, "UniswapV2Router: INVALID_PATH");
        // 计算每一步的输出金额
        amounts = getAmountsOut(amountIn, path);
        // 检查最终输出是否满足最小值要求
        require(amounts[amounts.length - 1] >= amountOutMin, "UniswapV2Router: INSUFFICIENT_OUTPUT_AMOUNT");
        // 将输入代币从用户账户转入第一个交易对
        _safeTransfer(path[0], UniswapV2Library.pairFor(factory, path[0], path[1]), amounts[0]);
        // 执行交换，WETH 发送到路由合约
        _swap(amounts, path, address(this));
        // 将 WETH 解包成 ETH
        IWETH(weth).withdraw(amounts[amounts.length - 1]);
        // 将 ETH 发送给用户
        _safeTransferETH(to, amounts[amounts.length - 1]);
    }

    /**
     * @notice 用精确 ETH 交换代币（支持手续费代币）
     * @param amountOutMin 最小输出金额（滑点保护）
     * @param path 交易路径（第一个必须是 WETH）
     * @param to 接收地址
     * @param deadline 截止时间
     * @dev 对于有转账手续费的代币，需要使用此函数
     *      因为普通交换函数假设输入金额等于实际到达交易对的金额
     */
    function swapExactETHForTokensSupportingFeeOnTransferTokens(
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external payable {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 确保路径的第一个代币是 WETH
        require(path[0] == weth, "UniswapV2Router: INVALID_PATH");
        // 获取输入金额
        uint256 amountIn = msg.value;
        // 将 ETH 包装成 WETH
        IWETH(weth).deposit{value: amountIn}();
        // 将 WETH 转入第一个交易对
        _safeTransferOwn(weth, UniswapV2Library.pairFor(factory, path[0], path[1]), amountIn);
        // 记录接收者交易前的代币余额
        uint256 balanceBefore = IERC20(path[path.length - 1]).balanceOf(to);
        // 执行支持手续费代币的交换
        _swapSupportingFeeOnTransferTokens(path, to);
        // 计算实际获得的代币数量
        uint256 amountOut = IERC20(path[path.length - 1]).balanceOf(to) - balanceBefore;
        // 检查输出是否满足最小值要求
        require(amountOut >= amountOutMin, "UniswapV2Router: INSUFFICIENT_OUTPUT_AMOUNT");
    }

    /**
     * @notice 用精确代币交换代币（支持手续费代币）
     * @param amountIn 精确的输入金额
     * @param amountOutMin 最小输出金额（滑点保护）
     * @param path 交易路径
     * @param to 接收地址
     * @param deadline 截止时间
     * @dev 对于有转账手续费的代币，需要使用此函数
     */
    function swapExactTokensForTokensSupportingFeeOnTransferTokens(
        uint256 amountIn,
        uint256 amountOutMin,
        address[] calldata path,
        address to,
        uint256 deadline
    ) external {
        // 检查交易是否过期
        require(deadline >= block.timestamp, "UniswapV2Router: EXPIRED");
        // 将输入代币从用户账户转入第一个交易对
        _safeTransfer(path[0], UniswapV2Library.pairFor(factory, path[0], path[1]), amountIn);
        // 记录接收者交易前的代币余额
        uint256 balanceBefore = IERC20(path[path.length - 1]).balanceOf(to);
        // 执行支持手续费代币的交换
        _swapSupportingFeeOnTransferTokens(path, to);
        // 计算实际获得的代币数量
        uint256 amountOut = IERC20(path[path.length - 1]).balanceOf(to) - balanceBefore;
        // 检查输出是否满足最小值要求
        require(amountOut >= amountOutMin, "UniswapV2Router: INSUFFICIENT_OUTPUT_AMOUNT");
    }

    // ============ 内部交换函数 ============

    /**
     * @notice 内部交换函数
     * @param amounts 金额数组
     * @param path 交易路径
     * @param _to 最终接收地址
     * @dev 执行多跳交换，沿着路径依次交换
     */
    function _swap(uint256[] memory amounts, address[] memory path, address _to) internal {
        // 遍历路径中的每一步
        for (uint256 i; i < path.length - 1; ++i) {
            // 获取输入和输出代币
            (address input, address output) = (path[i], path[i + 1]);
            // 对代币地址排序
            (address token0,) = UniswapV2Library.sortTokens(input, output);
            // 获取输出金额
            uint256 amountOut = amounts[i + 1];
            // 根据代币顺序确定 amount0Out 和 amount1Out
            (uint256 amount0Out, uint256 amount1Out) = input == token0 ? (uint256(0), amountOut) : (amountOut, uint256(0));
            // 确定接收地址：如果不是最后一步，发送到下一个交易对；否则发送到用户
            address to = i < path.length - 2 ? UniswapV2Library.pairFor(factory, output, path[i + 2]) : _to;
            // 执行交换
            IUniswapV2Pair(UniswapV2Library.pairFor(factory, input, output)).swap(
                amount0Out, amount1Out, to, new bytes(0) // 空的 bytes 表示不是闪电贷
            );
        }
    }

    /**
     * @notice 支持手续费代币的交换
     * @param path 交易路径
     * @param _to 最终接收地址
     * @dev 对于有转账手续费的代币，需要在每一步计算实际到达交易对的金额
     */
    function _swapSupportingFeeOnTransferTokens(address[] memory path, address _to) internal {
        // 遍历路径中的每一步
        for (uint256 i; i < path.length - 1; ++i) {
            // 获取输入和输出代币
            (address input, address output) = (path[i], path[i + 1]);
            // 对代币地址排序
            (address token0,) = UniswapV2Library.sortTokens(input, output);
            // 获取交易对合约
            IUniswapV2Pair pair = IUniswapV2Pair(UniswapV2Library.pairFor(factory, input, output));
            uint256 amountInput; // 实际输入金额
            uint256 amountOutput; // 计算的输出金额
            {
                // 获取当前储备量
                (uint256 reserve0, uint256 reserve1,) = pair.getReserves();
                // 根据代币顺序确定输入和输出储备量
                (uint256 reserveInput, uint256 reserveOutput) = input == token0
                    ? (reserve0, reserve1)
                    : (reserve1, reserve0);
                // 计算实际到达交易对的金额（余额 - 储备量）
                amountInput = IERC20(input).balanceOf(address(pair)) - reserveInput;
                // 根据实际输入计算输出金额
                amountOutput = UniswapV2Library.getAmountOut(amountInput, reserveInput, reserveOutput);
            }
            // 根据代币顺序确定 amount0Out 和 amount1Out
            (uint256 amount0Out, uint256 amount1Out) = input == token0
                ? (uint256(0), amountOutput)
                : (amountOutput, uint256(0));
            // 确定接收地址
            address to = i < path.length - 2 ? UniswapV2Library.pairFor(factory, output, path[i + 2]) : _to;
            // 执行交换
            pair.swap(amount0Out, amount1Out, to, new bytes(0));
        }
    }

    // ============ 价格查询函数 ============

    /**
     * @notice 根据输入计算输出金额（多跳）
     * @param amountIn 输入金额
     * @param path 交易路径
     * @return amounts 金额数组
     * @dev 计算沿着路径交换的每一步金额
     */
    function getAmountsOut(uint256 amountIn, address[] memory path) public view returns (uint256[] memory amounts) {
        // 确保路径至少有两个代币
        require(path.length >= 2, "UniswapV2Router: INVALID_PATH");
        // 创建金额数组
        amounts = new uint256[](path.length);
        // 设置初始输入金额
        amounts[0] = amountIn;
        // 遍历路径，计算每一步的输出
        for (uint256 i; i < path.length - 1; ++i) {
            // 获取当前交易对的储备量
            (uint256 reserveIn, uint256 reserveOut) = UniswapV2Library.getReserves(factory, path[i], path[i + 1]);
            // 计算输出金额
            amounts[i + 1] = UniswapV2Library.getAmountOut(amounts[i], reserveIn, reserveOut);
        }
    }

    /**
     * @notice 根据输出计算输入金额（多跳）
     * @param amountOut 输出金额
     * @param path 交易路径
     * @return amounts 金额数组
     * @dev 反向计算：从最终输出推算需要的初始输入
     */
    function getAmountsIn(uint256 amountOut, address[] memory path) public view returns (uint256[] memory amounts) {
        // 确保路径至少有两个代币
        require(path.length >= 2, "UniswapV2Router: INVALID_PATH");
        // 创建金额数组
        amounts = new uint256[](path.length);
        // 设置最终输出金额
        amounts[amounts.length - 1] = amountOut;
        // 反向遍历路径，计算每一步需要的输入
        for (uint256 i = path.length - 1; i > 0; --i) {
            // 获取当前交易对的储备量
            (uint256 reserveIn, uint256 reserveOut) = UniswapV2Library.getReserves(factory, path[i - 1], path[i]);
            // 计算需要的输入金额
            amounts[i - 1] = UniswapV2Library.getAmountIn(amounts[i], reserveIn, reserveOut);
        }
    }

    /**
     * @notice 引用：计算流动性数量
     * @param amountA 代币A数量
     * @param reserveA 代币A储备量
     * @param reserveB 代币B储备量
     * @return amountB 对应的代币B数量
     * @dev 根据储备比例计算对应的代币数量
     */
    function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) internal pure returns (uint256 amountB) {
        // 确保输入金额大于 0
        require(amountA > 0, "UniswapV2Router: INSUFFICIENT_AMOUNT");
        // 确保储备量大于 0
        require(reserveA > 0 && reserveB > 0, "UniswapV2Router: INSUFFICIENT_LIQUIDITY");
        // 按比例计算对应的代币B数量
        amountB = amountA * reserveB / reserveA;
    }
}
