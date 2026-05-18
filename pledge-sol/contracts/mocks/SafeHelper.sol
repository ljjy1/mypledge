// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test} from "forge-std/src/Test.sol";
import {console} from "forge-std/src/Test.sol";
import {Safe} from "@safe-global/safe-smart-account/contracts/Safe.sol";
import {SafeProxyFactory} from "@safe-global/safe-smart-account/contracts/proxies/SafeProxyFactory.sol";
import {Enum} from "@safe-global/safe-smart-account/contracts/libraries/Enum.sol";

/// @title SafeHelper
/// @notice 提供 Safe 多签钱包部署和交易执行的测试辅助功能
/// @dev 继承此合约后调用 _deploySafe 初始化多签钱包，再使用 _executeViaSafe 执行交易
contract SafeHelper is Test {
    Safe internal safe;
    SafeProxyFactory internal safeFactory;
    address internal safeAddress;

    uint256 internal _ownerPrivateKey1;
    uint256 internal _ownerPrivateKey2;
    address internal _owner1;
    address internal _owner2;
    uint256 internal _threshold;

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

        safeFactory = new SafeProxyFactory();
        Safe safeSingleton = new Safe();

        address[] memory owners = new address[](2);
        owners[0] = _owner1;
        owners[1] = _owner2;

        bytes memory initializer = abi.encodeCall(
            Safe.setup,
            (owners, thresholdRequired, address(0), bytes(""), address(0), address(0), 0, payable(address(0)))
        );

        safe = Safe(payable(safeFactory.createProxyWithNonce(address(safeSingleton), initializer, 0)));
        safeAddress = address(safe);
    }

    /// @dev 通过 Safe 多签执行交易 (按阈值收集签名)
    function _executeViaSafe(address to, bytes memory data) internal {
        bytes32 txHash = safe.getTransactionHash(
            to, 0, data, Enum.Operation.Call, 0, 0, 0,
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

        safe.execTransaction(
            to, 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), signatures
        );
    }
}
