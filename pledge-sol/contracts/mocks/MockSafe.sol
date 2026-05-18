// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title MockSafe
 * @notice 轻量级 Safe 多签模拟合约，专用于测试环境
 * @dev 实现 SafeHelper 所需的接口子集，避免 Safe.sol 的复杂内联汇编导致 stack-too-deep 编译错误
 *
 * 支持功能：
 * - setup()              初始化 owner 列表和签名阈值
 * - getTransactionHash() 计算 EIP-712 交易哈希
 * - execTransaction()    验证阈值签名并执行目标调用
 * - nonce()              获取当前 nonce
 * - getOwners()          获取 owner 列表
 */
contract MockSafe {
    address[] private _owners;
    uint256 private _threshold;
    uint256 private _nonce;

    event ExecutionSuccess(bytes32 txHash);

    receive() external payable {}

    /**
     * @notice 初始化多签钱包（与 Safe.setup 接口兼容）
     * @param owners owner 地址列表
     * @param threshold 所需签名阈值
     */
    function setup(
        address[] memory owners,
        uint256 threshold,
        address,    /* to — unused */
        bytes memory, /* data — unused */
        address,    /* fallbackHandler — unused */
        address,    /* paymentToken — unused */
        uint256,    /* payment — unused */
        address payable /* paymentReceiver — unused */
    ) external {
        // 对 owner 列表按地址升序排序，与签名拼接顺序一致
        for (uint256 i = 0; i < owners.length; i++) {
            for (uint256 j = i + 1; j < owners.length; j++) {
                if (owners[i] > owners[j]) {
                    address tmp = owners[i];
                    owners[i] = owners[j];
                    owners[j] = tmp;
                }
            }
        }
        _owners = owners;
        _threshold = threshold;
        _nonce = 0;
    }

    /// @notice 返回 owner 列表
    function getOwners() external view returns (address[] memory) {
        return _owners;
    }

    /// @notice 返回当前 nonce
    function nonce() external view returns (uint256) {
        return _nonce;
    }

    /**
     * @notice 计算交易哈希（EIP-712 风格，与 Safe 兼容）
     * @return bytes32 交易哈希
     */
    function getTransactionHash(
        address to,
        uint256 value,
        bytes memory data,
        uint8 operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address refundReceiver,
        uint256 _nonce_
    ) public view returns (bytes32) {
        bytes32 safeTxHash = keccak256(
            abi.encode(
                keccak256("safeTx(address to,uint256 value,bytes data,uint8 operation,uint256 safeTxGas,uint256 baseGas,uint256 gasPrice,address gasToken,address refundReceiver,uint256 nonce)"),
                to, value, keccak256(data), operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce_
            )
        );
        bytes32 domainSeparator = keccak256(abi.encode(block.chainid, address(this)));
        return keccak256(abi.encodePacked(hex"1901", domainSeparator, safeTxHash));
    }

    /**
     * @notice 执行多签交易：验证签名 → 递增 nonce → 调用目标
     * @dev 签名按 owner 地址升序排列，每签名 65 字节 (r32 + s32 + v1)
     * @return success 调用是否成功
     */
    function execTransaction(
        address to,
        uint256 value,
        bytes memory data,
        uint8 operation,
        uint256 safeTxGas,
        uint256 baseGas,
        uint256 gasPrice,
        address gasToken,
        address refundReceiver,
        bytes memory signatures
    ) external returns (bool) {
        uint256 currentNonce = _nonce;
        bytes32 txHash = getTransactionHash(to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, currentNonce);

        require(signatures.length == _threshold * 65, "MockSafe: invalid sig length");

        // 按签名顺序恢复签名者地址，与 owner 列表升序逐一比对
        // 使用内联汇编提取 bytes memory 中的数据（Solidity 不支持 memory 切片语法）
        for (uint256 i = 0; i < _threshold; i++) {
            uint256 offset = i * 65;
            bytes32 r;
            bytes32 s;
            uint8 v;
            assembly {
                let ptr := add(signatures, 0x20)
                r := mload(add(ptr, offset))
                s := mload(add(ptr, add(offset, 0x20)))
                // v 是第 offset+64 字节处的 1 字节值
                v := byte(0, mload(add(ptr, add(offset, 0x40))))
            }
            if (v < 27) v += 27; // 标准化 v 值
            address signer = ecrecover(txHash, v, r, s);
            require(signer == _owners[i], "MockSafe: invalid signer");
        }

        _nonce = currentNonce + 1;

        (bool success,) = to.call{value: value}(data);
        require(success, "GS013");

        emit ExecutionSuccess(txHash);
        return success;
    }
}
