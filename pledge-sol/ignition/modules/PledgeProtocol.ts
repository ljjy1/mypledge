/**
 * Pledge Protocol 完整部署模块
 *
 * 部署顺序：
 *   1. BscPledgeOracle（预言机）
 *   2. DebtToken（出借凭证 SP Token）
 *   3. DebtToken（借款凭证 JP Token）
 *   4. PledgePool（核心借贷池）
 *   5. 配置 DebtToken 的 minter 权限
 *
 * 使用方式：
 *   npx hardhat ignition deploy ignition/modules/PledgeProtocol.ts --network <network-name>
 *
 * 可通过参数自定义配置:
 *   npx hardhat ignition deploy ignition/modules/PledgeProtocol.ts \
 *     --network sepolia \
 *     --parameters ignition/parameters/sepolia.json
 */
import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

export default buildModule("PledgeProtocol", (m) => {
  // ============ 部署参数 ============

  // 多签钱包地址（PledgePool 和 Oracle 的 owner）
  const owner = m.getParameter("owner");

  // DEX 路由地址（如 PancakeSwap Router）
  const swapRouter = m.getParameter("swapRouter");

  // 手续费收款地址
  const feeAddress = m.getParameter("feeAddress");

  // 债务凭证代币名称/符号
  const spTokenName = m.getParameter("spTokenName", "SP Token");
  const spTokenSymbol = m.getParameter("spTokenSymbol", "sp");
  const jpTokenName = m.getParameter("jpTokenName", "JP Token");
  const jpTokenSymbol = m.getParameter("jpTokenSymbol", "jp");

  // ============ 1. 部署预言机 ============

  const oracle = m.contract("BscPledgeOracle", [owner]);

  // ============ 2. 部署债务凭证代币 ============

  const lendDebtToken = m.contract("DebtToken", [spTokenName, spTokenSymbol, owner], {
    id: "LendDebtToken",
  });
  const borrowDebtToken = m.contract("DebtToken", [jpTokenName, jpTokenSymbol, owner], {
    id: "BorrowDebtToken",
  });

  // ============ 3. 部署核心借贷池 ============

  const pool = m.contract("PledgePool", [
    oracle,
    swapRouter,
    feeAddress,
    owner,
  ]);

  // ============ 4. 配置权限 ============

  // 授权 PledgePool 可以铸造/销毁出借凭证
  m.call(lendDebtToken, "setMinter", [pool, true]);

  // 授权 PledgePool 可以铸造/销毁借款凭证
  m.call(borrowDebtToken, "setMinter", [pool, true]);

  // ============ 导出合约引用 ============

  return { oracle, lendDebtToken, borrowDebtToken, pool };
});
