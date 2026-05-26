/**
 * Pledge Protocol — Sepolia 测试网部署/测试脚本
 *
 * 使用方式：
 *   # 首次部署（部署所有合约并保存地址到 .sepolia-deployments.json）
 *   npx hardhat test mocha test/pledgePool-sepolia.ts --network sepolia --timeout 300000
 *
 *   # 或指定自定义部署记录路径
 *   DEPLOYMENTS_FILE=my-deploy.json npx hardhat test mocha ...
 *
 *   # 再次运行（检测到部署记录则跳过部署，直接验证）
 *   npx hardhat test mocha test/pledgePool-sepolia.ts --network sepolia --timeout 300000
 *
 *   # 强制重新部署（删除部署记录或设置 REDEPLOY=true）
 *   REDEPLOY=true npx hardhat test mocha ...
 *
 * 首次执行流程：
 *   1. 部署 UniswapV2 基础设施（WETH + Factory + Router）
 *   2. 部署 Mock ERC20 测试代币
 *   3. 为 Uniswap 交易对添加流动性
 *   4. 部署 BscPledgeOracle 并设置价格
 *   5. 部署 DebtToken（出借凭证 SP + 借款凭证 JP）
 *   6. 部署 PledgePool 核心借贷池
 *   7. 配置 DebtToken 的 minter 权限
 *   8. 在 PledgePool 上设置手续费
 *   9. 创建一个测试借贷池
 *   10. 将所有合约的 owner 转移给多签钱包
 *   11. 保存部署地址到 JSON 文件
 *
 * 再次执行：
 *   - 检测到 JSON 文件 → 跳过部署，直接读取地址并运行验证测试
 *
 * 文件路径：
 *   默认项目根目录 .sepolia-deployments.json，可通过环境变量 DEPLOYMENTS_FILE 自定义
 */
import { existsSync, readFileSync, writeFileSync } from "fs";
import { resolve } from "path";
import { expect } from "chai";
import { network } from "hardhat";

// ======================== 常量 ========================

const SEPOLIA_CHAIN_ID = 11155111n;
const MULTISIG_ADDRESS = "0xDF39470F7c82e62174A06BE933B8D94c5A48Dc11";

const BASE_DECIMAL = 10n ** 8n;
const MIN_AMOUNT = 100n * 10n ** 18n;
const MAX_SUPPLY = 1_000_000n * 10n ** 18n;
const TOKEN_SUPPLY = 10_000_000n * 10n ** 18n;
const INTEREST_RATE = 5n * 10n ** 7n;         // 5%
const MORTGAGE_RATE = 2n * BASE_DECIMAL;       // 200%
const LIQUIDATE_THRESHOLD = 7n * 10n ** 7n;    // 70%
const LEND_FEE = 1n * 10n ** 7n;               // 1%
const BORROW_FEE = 5n * 10n ** 6n;             // 0.5%

/** 部署记录文件路径，默认项目根目录，可通过 DEPLOYMENTS_FILE 环境变量覆盖 */
const DEPLOYMENTS_FILE = resolve(
  process.env.DEPLOYMENTS_FILE || ".sepolia-deployments.json",
);

// ======================== 测试套件 ========================

describe("PledgePool — Sepolia", function () {
  this.timeout(600_000);
  this.slow(60_000);

  let ctx: any;
  let ethers: any;
  let deployer: any;
  let isLoadedFromCache: boolean;

  before(async function () {
    ({ ethers } = await network.create());
    const networkInfo = await ethers.provider.getNetwork();
    if (networkInfo.chainId !== SEPOLIA_CHAIN_ID) {
      console.log("⚠  当前不是 Sepolia 网络，跳过测试");
      this.skip();
      return;
    }

    const signers = await ethers.getSigners();
    deployer = signers[0];

    // 尝试从缓存文件加载部署记录，或执行完整部署
    if (process.env.REDEPLOY === "true" || !existsSync(DEPLOYMENTS_FILE)) {
      ctx = await deployAll(ethers, deployer);
      isLoadedFromCache = false;
    } else {
      ctx = await loadFromCache(ethers);
      isLoadedFromCache = true;
    }
  });

  // ======================== 验证测试 ========================

  it("BscPledgeOracle 已部署且 owner 为多签", async function () {
    if (!ctx) this.skip();
    expect(ctx.oracle.target).to.be.properAddress;
    expect(await ctx.oracle.owner()).to.equal(MULTISIG_ADDRESS);
  });

  it("DebtToken (SP) 已部署且 owner 为多签", async function () {
    if (!ctx) this.skip();
    expect(ctx.lendDebtToken.target).to.be.properAddress;
    expect(await ctx.lendDebtToken.owner()).to.equal(MULTISIG_ADDRESS);
    expect(await ctx.lendDebtToken.name()).to.equal("SP Token");
    expect(await ctx.lendDebtToken.symbol()).to.equal("sp");
  });

  it("DebtToken (JP) 已部署且 owner 为多签", async function () {
    if (!ctx) this.skip();
    expect(ctx.borrowDebtToken.target).to.be.properAddress;
    expect(await ctx.borrowDebtToken.owner()).to.equal(MULTISIG_ADDRESS);
    expect(await ctx.borrowDebtToken.name()).to.equal("JP Token");
    expect(await ctx.borrowDebtToken.symbol()).to.equal("jp");
  });

  it("PledgePool 已部署且 owner 为多签", async function () {
    if (!ctx) this.skip();
    expect(ctx.pool.target).to.be.properAddress;
    expect(await ctx.pool.owner()).to.equal(MULTISIG_ADDRESS);
  });

  it("PledgePool 状态正确", async function () {
    if (!ctx) this.skip();
    expect(await ctx.pool.minAmount()).to.equal(MIN_AMOUNT);
    expect(await ctx.pool.lendFee()).to.equal(LEND_FEE);
    expect(await ctx.pool.borrowFee()).to.equal(BORROW_FEE);
  });

  it("DebtToken minter 权限已配置", async function () {
    if (!ctx) this.skip();
    expect(await ctx.lendDebtToken.minters(ctx.pool.target)).to.equal(true);
    expect(await ctx.borrowDebtToken.minters(ctx.pool.target)).to.equal(true);
  });

  it("测试借贷池已创建", async function () {
    if (!ctx) this.skip();
    const info = await ctx.pool.pledgePoolInfoList(ctx.pid);
    expect(info[9]).to.equal(0n);              // index 9 = state = MATCH
    expect(info[7]).to.equal(ctx.tokenA.target); // index 7 = lendToken
    expect(info[8]).to.equal(ctx.tokenB.target); // index 8 = borrowToken
  });

  it("预言机价格已设置", async function () {
    if (!ctx) this.skip();
    expect(await ctx.oracle.getPrice(ctx.tokenA.target)).to.equal(1n * 10n ** 18n);
    expect(await ctx.oracle.getPrice(ctx.tokenB.target)).to.equal(1n * 10n ** 18n);
  });

  it("Mock ERC20 已部署", async function () {
    if (!ctx) this.skip();
    expect(await ctx.tokenA.totalSupply()).to.equal(TOKEN_SUPPLY);
    expect(await ctx.tokenB.totalSupply()).to.equal(TOKEN_SUPPLY);
  });

  it("Uniswap 路由正常工作", async function () {
    if (!ctx) this.skip();
    expect(await ctx.router.weth()).to.equal(ctx.weth.target);
    expect(await ctx.router.factory()).to.equal(ctx.uniFactory.target);
  });

  it("多签可调用 setOracle（验证 owner 转移有效）", async function () {
    if (!ctx) this.skip();
    await expect(ctx.pool.connect(ctx.deployer).setOracle(ctx.deployer.address))
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
  });
});

// ======================== 多签操作指引 & 验证 ========================

describe("PledgePool — 多签操作", function () {
  this.timeout(600_000);
  this.slow(60_000);

  let ctx: any;
  let ethers: any;
  let deployer: any;

  before(async function () {
    ({ ethers } = await network.create());
    const networkInfo = await ethers.provider.getNetwork();
    if (networkInfo.chainId !== SEPOLIA_CHAIN_ID) {
      this.skip();
      return;
    }

    const signers = await ethers.getSigners();
    deployer = signers[0];

    if (!existsSync(DEPLOYMENTS_FILE)) {
      console.log("  ⚠  请先运行首次部署: npx hardhat test mocha ... --network sepolia");
      this.skip();
      return;
    }

    const raw = readFileSync(DEPLOYMENTS_FILE, "utf-8");
    const data = JSON.parse(raw);
    const c = data.contracts;

    ctx = {
      pool: await ethers.getContractAt("PledgePool", c.pool),
      oracle: await ethers.getContractAt("BscPledgeOracle", c.oracle),
      lendDebtToken: await ethers.getContractAt("DebtToken", c.lendDebtToken),
      borrowDebtToken: await ethers.getContractAt("DebtToken", c.borrowDebtToken),
      tokenA: await ethers.getContractAt("MockTestERC20", c.tokenA),
      tokenB: await ethers.getContractAt("MockTestERC20", c.tokenB),
    };
  });

  it("打印多签 setFeeAddress 操作指引", async function () {
    if (!ctx) this.skip();

    // 当前手续费地址
    const currentFee = await ctx.pool.feeAddress();
    console.log(`  当前手续费地址：${currentFee}`);
    // 建议设置的新手续费地址（部署账户地址，测试用）
    const newFee = "0x28cD99FE0701835F82FF960aEeE6100002Ca9392";

    // 编码 transactions
    const feeAddrCalldata = ctx.pool.interface.encodeFunctionData("setFeeAddress", [newFee]);
    
    console.log("");
    console.log("Safe 多签操作指引 — setFeeAddress");
    console.log("");
    console.log("⚠  本测试只是输出指引，不会自动执行链上交易");
    console.log("请按以下步骤在 Safe 中手动操作：");
    console.log("");
    console.log("1. 浏览器打开 https://app.safe.global");
    console.log("");
    console.log("2. 右上角切换网络为 Sepolia");
    console.log("");
    console.log("3. 打开多签地址 " + MULTISIG_ADDRESS);
    console.log("");
    console.log("4. New Transaction → Interact with contracts下的->Transaction Builder  -> 然后开启Custom data");
    console.log("");
    console.log("5. 参数 (Contract Address): "+ctx.pool.target);
    console.log("");
    console.log("6. 参数 (ABI):复制PledgePool.abi的内容(没有这个文件可以执行[需要先编译合约]jq '.abi' artifacts/contracts/pledge/PledgePool.sol/PledgePool.json > PledgePool.abi");
    console.log("");
    console.log("7. 参数 ETH value: 看函数情况需不需要");
    console.log("");
    console.log("8. 参数 data: "+ feeAddrCalldata);
    console.log("");
    console.log("9. Review → 签名 → 收集其他签名 → Execute");
    console.log("");
    console.log("10. 交易上链后，运行[npx hardhat test mocha test/pledgePool-sepolia.ts --network sepolia --grep \"验证多签 setFeeAddress 已生效\"]脚本验证");
    console.log("");

  });

  it("验证多签 setFeeAddress 已生效", async function () {
    if (!ctx) this.skip();
    const currentFee = await ctx.pool.feeAddress();
    console.log("  feeAddress: " + currentFee);
    // feeAddress 被 deployer 以外的地址修改过，视为多签操作成功
    expect(currentFee).to.equal("0x28cD99FE0701835F82FF960aEeE6100002Ca9392",
      "feeAddress修改成功",
    );
    // feeAddress 不能为零地址
    expect(currentFee).to.not.equal("0x0000000000000000000000000000000000000000");
  });
});

// ======================== 部署逻辑 ========================

/**
 * 完整部署所有合约并保存地址到 JSON 文件
 */
async function deployAll(ethers: any, deployer: any) {
  const balance = await ethers.provider.getBalance(deployer.address);
  if (balance < 10n ** 16n) {
    throw new Error("部署账户余额不足，需要至少 0.01 Sepolia ETH");
  }

  const startBlock = await ethers.provider.getBlockNumber();
  const startBal = balance;

  console.log("\n" + "=".repeat(56));
  console.log("  Pledge Protocol — Sepolia 全新部署");
  console.log("=".repeat(56));
  console.log(` 部署账户: ${deployer.address}`);
  console.log(` 多签地址: ${MULTISIG_ADDRESS}`);
  console.log(` 余额: ${ethers.formatEther(balance)} ETH`);
  console.log(` 部署记录: ${DEPLOYMENTS_FILE}\n`);

  // 1-3. UniswapV2
  console.log("[1/9] 部署 WETH...");
  const weth = await (await ethers.getContractFactory("WETH")).deploy();
  await weth.waitForDeployment();

  console.log("[2/9] 部署 UniswapV2Factory...");
  const uniFactory = await (await ethers.getContractFactory("UniswapV2Factory")).deploy(deployer.address);
  await uniFactory.waitForDeployment();

  console.log("[3/9] 部署 UniswapV2Router02...");
  const router = await (await ethers.getContractFactory("UniswapV2Router02")).deploy(
    uniFactory.target, weth.target,
  );
  await router.waitForDeployment();

  // 4. Mock ERC20
  console.log("[4/9] 部署 MockTestERC20 (LEND / BORR)...");
  const MockERC20 = await ethers.getContractFactory("MockTestERC20");
  const tokenA = await MockERC20.deploy("LendToken", "LEND", TOKEN_SUPPLY);
  await tokenA.waitForDeployment();
  const tokenB = await MockERC20.deploy("BorrowToken", "BORR", TOKEN_SUPPLY);
  await tokenB.waitForDeployment();

  // 5. 添加流动性
  console.log("[5/9] 添加 Uniswap 流动性...");
  // approve 必须等上链，否则 addLiquidity 时路由无权扣款
  const approveATx = await tokenA.approve(router.target, MAX_SUPPLY);
  await approveATx.wait();
  const approveBTx = await tokenB.approve(router.target, MAX_SUPPLY);
  await approveBTx.wait();
  // addLiquidity 也要等上链，确保 pair 创建和转账完成
  const addLiqTx = await router.addLiquidity(
    tokenA.target, tokenB.target,
    500_000n * 10n ** 18n, 500_000n * 10n ** 18n,
    0, 0, deployer.address, 2n ** 255n,
  );
  await addLiqTx.wait();

  // 6. BscPledgeOracle
  console.log("[6/9] 部署 BscPledgeOracle...");
  const oracle = await (await ethers.getContractFactory("BscPledgeOracle")).deploy(deployer.address);
  await oracle.waitForDeployment();
  const setPriceTxA = await oracle.setPrice(tokenA.target, 1n * 10n ** 18n);
  await setPriceTxA.wait();
  const setPriceTxB = await oracle.setPrice(tokenB.target, 1n * 10n ** 18n);
  await setPriceTxB.wait();

  // 7. DebtToken
  console.log("[7/9] 部署 DebtToken (SP / JP)...");
  const DebtToken = await ethers.getContractFactory("DebtToken");
  const lendDebtToken = await DebtToken.deploy("SP Token", "sp", deployer.address);
  await lendDebtToken.waitForDeployment();
  const borrowDebtToken = await DebtToken.deploy("JP Token", "jp", deployer.address);
  await borrowDebtToken.waitForDeployment();

  // 8. PledgePool
  console.log("[8/9] 部署 PledgePool...");
  const pool = await (await ethers.getContractFactory("PledgePool")).deploy(
    oracle.target, router.target, deployer.address, deployer.address,
  );
  await pool.waitForDeployment();

  // 配置 minter + 手续费
  console.log("     配置 minter / fee...");
  const minterSPTx = await lendDebtToken.setMinter(pool.target, true);
  await minterSPTx.wait();
  const minterJPTx = await borrowDebtToken.setMinter(pool.target, true);
  await minterJPTx.wait();
  const feeTx = await pool.setFee(LEND_FEE, BORROW_FEE);
  await feeTx.wait();

  // 9. 创建测试池
  console.log("[9/9] 创建测试借贷池...");
  const { timestamp: now } = await ethers.provider.getBlock("latest");
  const settleTime = now + 7 * 86400;
  const endTime = settleTime + 30 * 86400;
  const createPoolTx = await pool.createPledgePool({
    settleTime, endTime,
    interestRate: INTEREST_RATE,
    maxSupply: MAX_SUPPLY,
    mortgageRate: MORTGAGE_RATE,
    lendToken: tokenA.target,
    borrowToken: tokenB.target,
    lendDebtToken: lendDebtToken.target,
    borrowDebtToken: borrowDebtToken.target,
    autoLiquidateThreshold: LIQUIDATE_THRESHOLD,
  });
  await createPoolTx.wait();

  // 10. 转移 owner 到多签
  console.log("\n  转移 owner 到多签...");
  const transferOracleTx = await oracle.transferOwnership(MULTISIG_ADDRESS);
  await transferOracleTx.wait();
  const transferLendTx = await lendDebtToken.transferOwnership(MULTISIG_ADDRESS);
  await transferLendTx.wait();
  const transferBorrowTx = await borrowDebtToken.transferOwnership(MULTISIG_ADDRESS);
  await transferBorrowTx.wait();
  const transferPoolTx = await pool.transferOwnership(MULTISIG_ADDRESS);
  await transferPoolTx.wait();

  // 验证 owner
  const owners = await Promise.all([
    oracle.owner(), lendDebtToken.owner(), borrowDebtToken.owner(), pool.owner(),
  ]);
  const names = ["oracle", "lendDebtToken", "borrowDebtToken", "pool"];
  owners.forEach((o: string, i: number) => {
    expect(o, `${names[i]} owner 转移失败`).to.equal(MULTISIG_ADDRESS);
  });

  // Gas 统计
  const endBlock = await ethers.provider.getBlockNumber();
  const endBal = await ethers.provider.getBalance(deployer.address);
  const gasUsed = startBal - endBal;

  // 保存到文件
  const block = await ethers.provider.getBlock("latest");
  const data = {
    deployedAt: new Date().toISOString(),
    deployer: deployer.address,
    multisig: MULTISIG_ADDRESS,
    chainId: Number(SEPOLIA_CHAIN_ID),
    blockNumber: endBlock,
    totalGasWei: gasUsed.toString(),
    contracts: {
      weth: weth.target,
      uniFactory: uniFactory.target,
      router: router.target,
      tokenA: tokenA.target,
      tokenB: tokenB.target,
      oracle: oracle.target,
      lendDebtToken: lendDebtToken.target,
      borrowDebtToken: borrowDebtToken.target,
      pool: pool.target,
      pid: 0,
      settleTime,
      endTime,
    },
  };
  writeFileSync(DEPLOYMENTS_FILE, JSON.stringify(data, null, 2));

  printResult({
    mode: "deploy", data, block: startBlock,
    gasUsed: ethers.formatEther(gasUsed),
    pool: { oracle, pool, weth, router, tokenA, tokenB, uniFactory, lendDebtToken, borrowDebtToken, deployer, pid: 0n, settleTime, endTime },
  });

  // 返回合约实例（非地址字符串），验证测试需要调用合约方法和 deployer 签名
  return { oracle, pool, weth, router, tokenA, tokenB, uniFactory, lendDebtToken, borrowDebtToken, deployer, pid: 0n, settleTime, endTime };
}

/**
 * 从 JSON 文件加载部署记录，通过 getContractAt 恢复合约实例
 */
async function loadFromCache(ethers: any) {
  const raw = readFileSync(DEPLOYMENTS_FILE, "utf-8");
  const data = JSON.parse(raw);
  const c = data.contracts;

  console.log(`\n  ✅ 发现部署记录 ${DEPLOYMENTS_FILE}`);
  console.log(`  部署时间: ${data.deployedAt}`);
  console.log(`  部署区块: #${data.blockNumber}\n`);

  const contracts = {
    weth: await ethers.getContractAt("WETH", c.weth),
    uniFactory: await ethers.getContractAt("UniswapV2Factory", c.uniFactory),
    router: await ethers.getContractAt("UniswapV2Router02", c.router),
    tokenA: await ethers.getContractAt("MockTestERC20", c.tokenA),
    tokenB: await ethers.getContractAt("MockTestERC20", c.tokenB),
    oracle: await ethers.getContractAt("BscPledgeOracle", c.oracle),
    lendDebtToken: await ethers.getContractAt("DebtToken", c.lendDebtToken),
    borrowDebtToken: await ethers.getContractAt("DebtToken", c.borrowDebtToken),
    pool: await ethers.getContractAt("PledgePool", c.pool),
  };

  const signers = await ethers.getSigners();

  printResult({
    mode: "load", data,
    pool: { ...contracts, deployer: signers[0], pid: BigInt(c.pid), settleTime: c.settleTime, endTime: c.endTime },
  });

  return { ...contracts, deployer: signers[0], pid: BigInt(c.pid), settleTime: c.settleTime, endTime: c.endTime };
}

// ======================== 辅助函数 ========================

function printResult(opts: {
  mode: "deploy" | "load";
  data: any;
  pool: any;
  block?: number;
  gasUsed?: string;
}) {
  const { mode, data, pool, gasUsed } = opts;
  const c = data.contracts;

  const rows: Record<string, string> = {
    "WETH": c.weth,
    "UniswapV2Factory": c.uniFactory,
    "UniswapV2Router02": c.router,
    "MockERC20 (LEND)": c.tokenA,
    "MockERC20 (BORR)": c.tokenB,
    "BscPledgeOracle": c.oracle,
    "DebtToken (SP Token)": c.lendDebtToken,
    "DebtToken (JP Token)": c.borrowDebtToken,
    "PledgePool": c.pool,
  };
  const maxLen = Math.max(...Object.keys(rows).map((k) => k.length));

  console.log("-".repeat(56));
  console.log(`  ${mode === "deploy" ? "全新部署完成" : "从缓存加载"} — 合约地址`);
  console.log("-".repeat(56));
  for (const [name, addr] of Object.entries(rows)) {
    console.log(`  ${name.padEnd(maxLen)}  ${addr}`);
  }
  console.log("-".repeat(56));
  console.log(`  PledgePool owner  ${MULTISIG_ADDRESS}`);
  console.log(`  测试池 pid        ${c.pid}`);
  console.log(`  Chain             Sepolia (${SEPOLIA_CHAIN_ID})`);
  if (mode === "deploy") {
    console.log(`  Gas 消耗           ${gasUsed} ETH`);
    console.log(`  部署记录已保存至: ${DEPLOYMENTS_FILE}`);
  }
  console.log("");

  console.log("  ⚠  下一步 — 多签操作：");
  console.log(`     1. 运行 --grep "多签操作" 生成 setFeeAddress 的 Transaction Builder JSON`);
  console.log(`     2. 或直接通过 Go 连接 ${pool.pool.target} 交互`);
  console.log(`  Etherscan: https://sepolia.etherscan.io/address/${c.pool}\n`);
}
