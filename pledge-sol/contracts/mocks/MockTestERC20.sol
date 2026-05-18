// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol"; // 导入 OpenZeppelin 的 ERC20 实现

/**
 * @title TestToken 测试代币
 * @notice 用于测试的 ERC20 代币
 * @dev 这是一个简单的 ERC20 代币合约，用于测试 UniswapV2 功能
 */
contract MockTestERC20 is ERC20 {

    /**
     * @notice 构造函数
     * @param name 代币名称
     * @param symbol 代币符号
     * @param initialSupply 初始供应量
     */
    constructor(string memory name, string memory symbol, uint256 initialSupply) ERC20(name, symbol) {
        _mint(msg.sender, initialSupply); // 铸造初始供应量给部署者
    }
}
