// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/IDebtToken.sol";

/**
 * @title DebtToken
 * @notice 债务凭证代币合约
 * @dev 实现出借和借款凭证的铸造与销毁
 *      owner 为 Safe 多签钱包，关键操作需要多签确认
 * @author Pledge Protocol
 */
contract DebtToken is IDebtToken, Ownable{
    /// @notice 代币名称
    string public name;
    /// @notice 代币符号
    string public symbol;
    /// @notice 代币精度
    uint8 public constant decimals = 18;


    /// @notice 授权的铸造者列表
    mapping(address => bool) public minters;

    /// @notice 用户 -> 余额
    mapping(address => uint256) private _balances;

    /// @notice 总供应量
    uint256 public totalSupply;

    /// @notice 授权的铸造者变更事件
    event MinterAdded(address indexed minter, bool status);

    /// @notice 债务凭证铸造事件
    event DebtMinted(address indexed account, uint256 amount);

    /// @notice 债务凭证销毁事件
    event DebtBurned(address indexed account, uint256 amount);

    /// @notice 未授权铸造错误
    error NotAuthorizedMinter();

    /// @notice 余额不足错误
    error InsufficientBalance();

    /// @notice 无效参数错误
    error InvalidParameter();


    constructor(string memory _name, string memory _symbol, address _owner) Ownable(_owner) {
        name = _name;
        symbol = _symbol;
    }

    /// @notice 设置铸造者权限
    function setMinter(address minter, bool status) external onlyOwner {
        minters[minter] = status;
        emit MinterAdded(minter, status);
    }

    /// @notice 铸造债务凭证
    function mint(address account, uint256 amount) external {
        if (!minters[msg.sender]) revert NotAuthorizedMinter();
        if (account == address(0)) revert InvalidParameter();
        if (amount == 0) revert InvalidParameter();

        _balances[account] += amount;
        totalSupply += amount;
        emit DebtMinted(account, amount);
    }

    /// @notice 销毁债务凭证
    function burn(address account, uint256 amount) external {
        if (!minters[msg.sender]) revert NotAuthorizedMinter();
        if (_balances[account] < amount) revert InsufficientBalance();

        _balances[account] -= amount;
        totalSupply -= amount;
        emit DebtBurned(account, amount);
    }

    /// @notice 获取指定账户的债务余额
    function balanceOf(address account) external view returns (uint256) {
        return _balances[account];
    }
}
