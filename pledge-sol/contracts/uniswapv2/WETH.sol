// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

import "@openzeppelin/contracts/token/ERC20/ERC20.sol"; // 导入 OpenZeppelin 的 ERC20 实现
import "./interfaces/IWETH.sol"; // 导入 WETH 接口

/**
 * @title WETH - Wrapped Ether
 * @notice 以太坊包装代币合约
 * @dev 将 ETH 包装成 ERC20 代币，方便在 DEX 中交易
 *      WETH 是一种特殊的 ERC20 代币，与 ETH 1:1 锚定
 *      用户可以存入 ETH 获得 WETH，也可以销毁 WETH 取回 ETH
 *      这样 ETH 就可以在 DEX 中像其他 ERC20 代币一样交易
 */
contract WETH is ERC20, IWETH {
    // 存款事件：当用户存入 ETH 时触发
    event Deposit(address indexed account, uint256 amount);
    // 取款事件：当用户取回 ETH 时触发
    event Withdrawal(address indexed account, uint256 amount);

    /**
     * @notice 构造函数：初始化 WETH 代币
     * @dev 调用 ERC20 构造函数设置代币名称和符号
     *      WETH 初始供应量为 0，通过存款铸造
     */
    constructor() ERC20("Wrapped Ether", "WETH") {
        // WETH 初始供应量为 0，通过存款铸造
        // 不需要在这里执行任何操作
    }

    /**
     * @notice 存款函数：将 ETH 包装成 WETH
     * @dev 用户发送 ETH，合约铸造等量 WETH 给用户
     *      1 ETH = 1 WETH，始终保持 1:1 锚定
     */
    function deposit() external payable {
        // 铸造等量的 WETH 给调用者
        // msg.value 是用户发送的 ETH 数量
        _mint(msg.sender, msg.value);
        // 触发存款事件，记录存款操作
        emit Deposit(msg.sender, msg.value);
    }

    /**
     * @notice 取款函数：将 WETH 解包成 ETH
     * @param amount 取款金额（WETH 数量）
     * @dev 用户销毁 WETH，合约返还等量 ETH
     *      1 WETH = 1 ETH，始终保持 1:1 锚定
     */
    function withdraw(uint256 amount) external {
        // 检查用户 WETH 余额是否足够
        require(balanceOf(msg.sender) >= amount, "WETH: insufficient balance");
        // 销毁用户指定数量的 WETH
        _burn(msg.sender, amount);
        // 向用户发送等量的 ETH
        // 使用 call 方法发送 ETH，这是最安全的发送方式
        (bool success,) = payable(msg.sender).call{value: amount}("");
        // 确保 ETH 发送成功
        require(success, "WETH: transfer failed");
        // 触发取款事件，记录取款操作
        emit Withdrawal(msg.sender, amount);
    }

    /**
     * @notice 接收 ETH 的回调函数
     * @dev 当合约收到 ETH 时自动调用
     *      这个函数使得合约可以直接接收 ETH 转账
     *      接收到的 ETH 会自动转换为 WETH
     */
    receive() external payable {
        // 铸造等量的 WETH 给发送者
        _mint(msg.sender, msg.value);
        // 触发存款事件
        emit Deposit(msg.sender, msg.value);
    }
}
