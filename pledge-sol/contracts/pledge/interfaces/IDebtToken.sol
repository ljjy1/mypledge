// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title IDebtToken
 * @notice 债务凭证代币接口
 */
interface IDebtToken {
    /**
     * @notice 设置铸造者权限
     * @param minter 铸造者地址
     * @param status 授权状态
     */
    function setMinter(address minter, bool status) external;

    /**
     * @notice 铸造债务凭证
     * @param account 目标账户
     * @param amount 铸造数量
     */
    function mint(address account, uint256 amount) external;

    /**
     * @notice 销毁债务凭证
     * @param account 目标账户
     * @param amount 销毁数量
     */
    function burn(address account, uint256 amount) external;

    /**
     * @notice 获取指定账户的债务余额
     * @param account 目标账户
     * @return 余额
     */
    function balanceOf(address account) external view returns (uint256);
}
