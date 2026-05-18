// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {SafeHelper} from "../mocks/SafeHelper.sol";
import {DebtToken} from "./DebtToken.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {console} from "forge-std/src/console.sol";
import {Enum} from "@safe-global/safe-smart-account/contracts/libraries/Enum.sol";

/// @title DebtTokenTest
/// @notice DebtToken 合约测试 — 使用 Safe 多签钱包作为 owner
contract DebtTokenTest is SafeHelper {
    DebtToken debtToken;

    address minter = makeAddr("minter");
    address userA = makeAddr("userA");
    address userB = makeAddr("userB");
    address random = makeAddr("random");

    event MinterAdded(address indexed minter, bool status);
    event DebtMinted(address indexed account, uint256 amount);
    event DebtBurned(address indexed account, uint256 amount);

    function setUp() public {
        _deploySafe(0xA11CE, 0xB0B, 2);

        debtToken = new DebtToken("Pledge Debt Token", "PDT", safeAddress);

        _executeViaSafe(
            address(debtToken),
            abi.encodeCall(debtToken.setMinter, (minter, true))
        );
    }

    // ============ 构造函数测试 ============

    function test_constructor_setsOwner() public view {
        assertEq(debtToken.owner(), safeAddress);
    }

    function test_constructor_revertsOnZeroAddress() public {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableInvalidOwner.selector,address(0)));
        new DebtToken("name", "sym", address(0));
    }

    function test_constructor_setsMetadata() public view {
        assertEq(debtToken.name(), "Pledge Debt Token");
        assertEq(debtToken.symbol(), "PDT");
        assertEq(debtToken.decimals(), 18);
    }

    // ============ setMinter 测试 ============

    function test_setMinter_authorizesMinter() public {
        address newMinter = makeAddr("newMinter");

        vm.expectEmit(true, false, false, true);
        emit MinterAdded(newMinter, true);

        _executeViaSafe(
            address(debtToken),
            abi.encodeCall(debtToken.setMinter, (newMinter, true))
        );

        assertTrue(debtToken.minters(newMinter));
    }

    function test_setMinter_deauthorizesMinter() public {
        _executeViaSafe(
            address(debtToken),
            abi.encodeCall(debtToken.setMinter, (minter, false))
        );
        assertFalse(debtToken.minters(minter));
    }

    function test_setMinter_revertsForNonOwner() public {
        vm.prank(random);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, random));
        debtToken.setMinter(minter, true);
    }

    // ============ mint 测试 ============

    function test_mint_success() public {
        vm.expectEmit(true, false, false, true);
        emit DebtMinted(userA, 100e18);

        vm.prank(minter);
        debtToken.mint(userA, 100e18);

        assertEq(debtToken.balanceOf(userA), 100e18);
    }

    function test_mint_revertsForNonMinter() public {
        vm.prank(random);
        vm.expectRevert(DebtToken.NotAuthorizedMinter.selector);
        debtToken.mint(userA, 100e18);
    }

    function test_mint_revertsForZeroAddress() public {
        vm.prank(minter);
        vm.expectRevert(DebtToken.InvalidParameter.selector);
        debtToken.mint(address(0), 100e18);
    }

    function test_mint_revertsForZeroAmount() public {
        vm.prank(minter);
        vm.expectRevert(DebtToken.InvalidParameter.selector);
        debtToken.mint(userA, 0);
    }

    function test_mint_cumulative() public {
        vm.startPrank(minter);
        debtToken.mint(userA, 100e18);
        debtToken.mint(userA, 50e18);
        vm.stopPrank();

        assertEq(debtToken.balanceOf(userA), 150e18);
    }

    // ============ burn 测试 ============

    function test_burn_success() public {
        vm.prank(minter);
        debtToken.mint(userA, 100e18);

        vm.expectEmit(true, false, false, true);
        emit DebtBurned(userA, 40e18);

        vm.prank(minter);
        debtToken.burn(userA, 40e18);

        assertEq(debtToken.balanceOf(userA), 60e18);
    }

    function test_burn_revertsForNonMinter() public {
        vm.prank(random);
        vm.expectRevert(DebtToken.NotAuthorizedMinter.selector);
        debtToken.burn(userA, 100e18);
    }

    function test_burn_revertsForInsufficientBalance() public {
        vm.prank(minter);
        debtToken.mint(userA, 50e18);

        vm.prank(minter);
        vm.expectRevert(DebtToken.InsufficientBalance.selector);
        debtToken.burn(userA, 100e18);
    }

    function test_burn_fullAmount() public {
        vm.prank(minter);
        debtToken.mint(userA, 100e18);

        vm.prank(minter);
        debtToken.burn(userA, 100e18);

        assertEq(debtToken.balanceOf(userA), 0);
    }

    // ============ 多用户测试 ============

    function test_multipleUsers_independentBalances() public {
        vm.startPrank(minter);
        debtToken.mint(userA, 100e18);
        debtToken.mint(userB, 200e18);
        vm.stopPrank();

        assertEq(debtToken.balanceOf(userA), 100e18);
        assertEq(debtToken.balanceOf(userB), 200e18);
    }

    // ============ Safe 多签安全测试 ============

    function test_safeRevertsWithSingleSignature() public {
        address newMinter = makeAddr("singleSigMinter");

        bytes memory data = abi.encodeCall(debtToken.setMinter, (newMinter, true));
        bytes32 txHash = safe.getTransactionHash(
            address(debtToken), 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), safe.nonce()
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(_ownerPrivateKey1, txHash);
        bytes memory sig = abi.encodePacked(r, s, v);

        vm.expectRevert();
        safe.execTransaction(
            address(debtToken), 0, data, Enum.Operation.Call,
            0, 0, 0, address(0), payable(address(0)), sig
        );

        assertFalse(debtToken.minters(newMinter));
    }
}
