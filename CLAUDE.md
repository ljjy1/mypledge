# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Pledge Protocol — a DeFi lending protocol on BSC. Users deposit collateral, borrow assets, and liquidators can liquidate undercollateralized positions. The project is organized as a monorepo with three modules:

- `pledge-sol/` — Solidity smart contracts (Hardhat 3, Solidity 0.8.28)
- `pledge-fe/` — Frontend (placeholder, not yet implemented)
- `pledge-be/` — Backend (placeholder, not yet implemented)

## Smart Contract Architecture

Three core contracts with interface separation under `contracts/pledge/interfaces/`:

- **PledgePool** — Core lending pool. Manages pool creation, deposit, borrow, repay, withdraw, and liquidation. Uses OpenZeppelin ReentrancyGuard. Owner is a Safe multisig wallet.
- **DebtToken** — Debt receipt token (ERC20-like, non-transferable). Tracks lender/borrower positions per asset with double mapping `(address => DebtType => uint256)`.
- **BscPledgeOracle** — Price oracle with admin-controlled price updates. Prices are in USD with 1e8 precision.

Key design patterns:
- All admin functions restricted to Safe multisig wallet (`onlyOwner` modifier checking against stored `owner` address)
- Collateral/liquidation factors use 1e18 precision (must be < 1e18, liquidationFactor < collateralFactor)
- Interest calculation: `borrowed * interestRate * timeElapsed / (365 days * 1e18)`
- Low-level `call` for ERC20 transfers instead of IERC20 interface (gas-efficient, handles non-standard tokens)

## Development Commands

All commands run from `pledge-sol/` directory:

```bash
# Install dependencies
npm install

# Run all tests (Solidity + Mocha/TypeScript)
npx hardhat test

# Run only Solidity (Foundry-style) tests
npx hardhat test solidity

# Run only Mocha/TypeScript tests
npx hardhat test mocha

# Compile contracts
npx hardhat compile

# Deploy to local chain
npx hardhat ignition deploy ignition/modules/Counter.ts

# Deploy to Sepolia (requires SEPOLIA_PRIVATE_KEY)
npx hardhat ignition deploy --network sepolia ignition/modules/Counter.ts
```

## Configuration

Hardhat config uses Hardhat 3 Beta with ESM (`"type": "module"` in package.json). Networks defined:

- `hardhatMainnet` — EDR simulated L1
- `hardhatOp` — EDR simulated Optimism chain type
- `sepolia` — HTTP RPC, requires `SEPOLIA_RPC_URL` and `SEPOLIA_PRIVATE_KEY` env vars

Two Solidity compiler profiles: `default` (no optimizer) and `production` (optimizer enabled, 200 runs).

## Test Structure

- Solidity tests: `contracts/*.t.sol` — Foundry-style using `forge-std/Test.sol`
- TypeScript tests: `test/*.ts` — Mocha + Chai + ethers.js via `network.create()`

## Conventions

- Solidity version: `^0.8.28`
- NatSpec comments in Chinese (Chinese documentation for Chinese team)
- Custom errors over require strings for gas efficiency
- SPDX license: MIT for all pledge contracts
