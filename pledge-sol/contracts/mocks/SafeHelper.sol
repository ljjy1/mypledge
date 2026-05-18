// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/src/Test.sol";
import {MockSafe} from "./MockSafe.sol";

/// @title SafeHelper
/// @notice 提供 Safe 多签钱包部署和交易执行的测试辅助功能
/// @dev 继承此合约后调用 _deploySafe 初始化多签钱包，再使用 _executeViaSafe 执行交易
///      使用轻量级 MockSafe 替代真实 Safe.sol，以避免 stack-too-deep 编译错误
contract SafeHelper is Test {
    MockSafe internal safe;
    address internal safeAddress;

    uint256 internal _ownerPrivateKey1;
    uint256 internal _ownerPrivateKey2;
    address internal _owner1;
    address internal _owner2;
    uint256 internal _threshold;

    // 标记下一次 _executeViaSafe 预期 GS013 回滚（在 execTransaction 前注入 expectRevert）
    bool internal _mockSafeExpectGs013;

    /// @dev 部署 Safe 多签钱包并设置参数
    /// @param key1 第一个所有者私钥
    /// @param key2 第二个所有者私钥
    /// @param thresholdRequired 所需签名数量 (1 或 2)
    function _deploySafe(uint256 key1, uint256 key2, uint256 thresholdRequired) internal {
        _ownerPrivateKey1 = key1;
        _ownerPrivateKey2 = key2;
        _owner1 = vm.addr(key1);
        _owner2 = vm.addr(key2);
        _threshold = thresholdRequired;

        safe = new MockSafe();
        address[] memory owners = new address[](2);
        owners[0] = _owner1;
        owners[1] = _owner2;

        safe.setup(owners, thresholdRequired, address(0), bytes(""), address(0), address(0), 0, payable(address(0)));
        safeAddress = address(safe);
    }

    /// @dev 标记下一次 _executeViaSafe 预期抛出 GS013 错误（Safe 交易执行失败）
    function _expectGSO13() internal {
        _mockSafeExpectGs013 = true;
    }

    /// @dev 通过 Safe 多签执行交易 (按阈值收集签名)
    function _executeViaSafe(address to, bytes memory data) internal {
        bytes32 txHash = safe.getTransactionHash(
            to, 0, data, 0, 0, 0, 0,
            address(0), payable(address(0)), safe.nonce()
        );

        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(_ownerPrivateKey1, txHash);
        (uint8 v2, bytes32 r2, bytes32 s2) = vm.sign(_ownerPrivateKey2, txHash);

        // 按地址排序拼接签名
        bytes memory signatures;
        if (_owner1 < _owner2) {
            signatures = abi.encodePacked(r1, s1, v1, r2, s2, v2);
        } else {
            signatures = abi.encodePacked(r2, s2, v2, r1, s1, v1);
        }

        // 在 execTransaction 前注入 expectRevert，避免被 getTransactionHash/nonce 提前消耗
        if (_mockSafeExpectGs013) {
            _mockSafeExpectGs013 = false;
            vm.expectRevert("GS013");
        }

        safe.execTransaction(
            to, 0, data, 0,
            0, 0, 0, address(0), payable(address(0)), signatures
        );
    }
}
