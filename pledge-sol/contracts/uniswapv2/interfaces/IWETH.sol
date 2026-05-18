// SPDX-License-Identifier: MIT
// 开源许可证声明，使用 MIT 许可证
pragma solidity ^0.8.28; // 声明 Solidity 版本，要求 0.8.20 或更高版本

/**
 * @title IWETH 接口
 * @notice WETH 合约接口
 * @dev 定义 WETH 的存款和取款功能
 *      WETH 是 Wrapped Ether 的缩写，是以太坊的 ERC20 包装版本
 *      主要用于在 DEX 中交易 ETH，因为 ETH 不是 ERC20 代币
 */
interface IWETH {
    /**
     * @notice 存款函数：将 ETH 包装成 WETH
     * @dev 用户发送 ETH，合约铸造等量 WETH 给用户
     *      1 ETH = 1 WETH，始终保持 1:1 锚定
     */
    function deposit() external payable;

    /**
     * @notice 取款函数：将 WETH 解包成 ETH
     * @param amount 取款金额（WETH 数量）
     * @dev 用户销毁 WETH，合约返还等量 ETH
     *      1 WETH = 1 ETH，始终保持 1:1 锚定
     */
    function withdraw(uint256 amount) external;
}
