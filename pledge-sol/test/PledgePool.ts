// 导入 Chai 断言库，用于编写测试断言（如 expect、assert 等）
import { expect } from "chai";
// 导入 Hardhat 的 network 对象，用于与区块链节点交互
import { network } from "hardhat";
// ESM 兼容的 require 方式，用于读取 Safe 合约的预编译 JSON 产物
import { createRequire } from "module";
const _require = createRequire(import.meta.url);

// 通过 Hardhat 网络创建 ethers.js 实例，用于部署合约和发送交易
const { ethers } = await network.create();

// ======================== Safe 合约加载函数 ========================

/**
 * 从 @safe-global/safe-contracts 包中读取 Safe 合约的 ABI 和 bytecode
 * 使用 npm 包自带的预编译产物，不依赖 Hardhat 本地编译
 */
function loadSafeBuildInfo() {
  const safeArtifact = _require("@safe-global/safe-contracts/build/artifacts/contracts/Safe.sol/Safe.json");
  const factoryArtifact = _require("@safe-global/safe-contracts/build/artifacts/contracts/proxies/SafeProxyFactory.sol/SafeProxyFactory.json");

  return {
    safeAbi: safeArtifact.abi,
    safeBytecode: safeArtifact.bytecode,
    factoryAbi: factoryArtifact.abi,
    factoryBytecode: factoryArtifact.bytecode,
  };
}

// ======================== 测试常量定义 ========================

// 基础小数位数：1e8，用于抵押率、利率等百分比参数的精度
const BASE_DECIMAL = 10n ** 8n;
// 一年的秒数：365天 * 24小时 * 3600秒，用于时间相关的计算
const BASE_YEAR = 365n * 24n * 3600n;
// 最小存入金额：100 * 1e18，低于此金额的交易将被拒绝
const MIN_AMOUNT = 100n * 10n ** 18n;
// 最大供应量：1,000,000 * 1e18，池子募资上限
const MAX_SUPPLY = 1_000_000n * 10n ** 18n;
// 测试代币的总供应量：10,000,000 * 1e18
const TOKEN_SUPPLY = 10_000_000n * 10n ** 18n;
// 利率：5% （精度 1e8，所以 5 * 10^7 = 5%）
const INTEREST_RATE = 5n * 10n ** 7n;
// 抵押率：200% （2 * 1e8 = 200%，即借出的 lendToken 需要 2 倍价值的抵押品）
const MORTGAGE_RATE = 2n * BASE_DECIMAL;
// 自动清算阈值：70% （7 * 10^7 = 70%），当抵押率低于此值时触发清算
const LIQUIDATE_THRESHOLD = 7n * 10n ** 7n;

// ======================== 测试上下文接口 ========================

/**
 * PoolContext 接口定义了测试过程中需要共享的全部合约实例和账号信息
 * 每个 describe 块的 before 钩子通过 deployBase() 构建一个 PoolContext
 * 后续的 it 测试用例通过 ctx 访问这些实例
 */
interface PoolContext {
  signers: any[];      // 签名者账号列表 —— 需要时可用其他账号
  safe: any;           // Safe 多签钱包合约 —— 作为 PledgePool 的 owner（所有管理员操作需 2/3 多签）
  safeOwners: any[];   // Safe 多签的 owner EOA 账号列表 —— 用于收集多签签名
  pool: any;           // PledgePool 合约实例 —— 被测试的核心借贷池
  oracle: any;         // BscPledgeOracle 预言机合约 —— 提供资产价格
  weth: any;           // WETH 合约 —— Uniswap 路由所需的封装 ETH
  router: any;         // UniswapV2Router02 路由合约 —— 用于 DEX 兑换
  tokenA: any;         // 测试代币 A（lendToken）—— 出借人存入的资产
  tokenB: any;         // 测试代币 B（borrowToken）—— 借款人质押的资产
  lendDebt: any;       // 出借债务代币（SP Token）—— 代表出借人的债权凭证
  borrowDebt: any;     // 借款债务代币（JP Token）—— 代表借款人的债务凭证
  owner: any;          // 部署合约用的 EOA 账号（非 PledgePool 的 owner，真正的 owner 是 Safe 多签）
  lender: any;         // 出借人账号 —— 向池子存入 lendToken
  lender2: any;        // 第二个出借人 —— 用于多用户场景测试
  borrower: any;       // 借款人账号 —— 向池子质押 borrowToken
  borrower2: any;      // 第二个借款人 —— 用于多用户场景测试
  feeCollector: any;   // 手续费收款地址 —— 收取借贷手续费
  liquidator: any;     // 清算人账号 —— 执行清算操作
  pid: bigint;         // 池子 ID（默认创建的池子 ID 为 0）
  settleTime: number;  // 结算时间戳 —— 到达此时间后池子可以进入结算阶段
  endTime: number;     // 结束时间戳 —— 到达此时间后池子可以执行 finish
}

// ======================== Safe 多签执行辅助函数 ========================

/**
 * safeExec 通过 Safe 多签钱包执行一笔交易
 * 工作流程：
 *   1. 构造 SafeTx 结构体并用 EIP-712 签名
 *   2. 收集 threshold 个多签 owner 的签名（按地址升序排列）
 *   3. 调用 Safe.execTransaction 执行交易
 *   4. safeTxGas=0 时，如果内部交易失败 Safe 会回滚（GS013），适合验证成功场景
 *   5. safeTxGas>0 时，内部交易失败 Safe 仍会成功但发出 ExecutionFailure 事件
 *
 * @param safe       Safe 多签合约实例
 * @param contract   被调用的合约实例（用于编码函数调用数据）
 * @param method     方法名
 * @param args       方法参数数组
 * @param safeOwners Safe 多签的 owner EOA 列表（前 threshold 个用于签名）
 * @param chainId    链 ID（用于 EIP-712 域分隔符）
 * @param overrides  可选参数，如 { safeTxGas: bigint }
 * @returns          交易收据
 */
async function safeExec(
  safe: any,
  contract: any,
  method: string,
  args: any[],
  safeOwners: any[],
  chainId?: bigint,
  overrides?: { safeTxGas?: bigint }
): Promise<any> {
  // 如果未提供 chainId，从网络中获取
  if (!chainId) {
    chainId = (await ethers.provider.getNetwork()).chainId;
  }
  const data = contract.interface.encodeFunctionData(method, args);
  // 获取 Safe 当前 nonce（防止重放攻击）
  const nonce: bigint = await safe.nonce();
  // 获取多签阈值（需要的签名数）
  const threshold = Number(await safe.getThreshold());
  // Safe 地址（作为 EIP-712 域分隔符的 verifyingContract）
  const verifyingContract = await safe.getAddress();
  // 构建 EIP-712 域分隔符
  const domain = { chainId, verifyingContract };
  // 定义 SafeTx 的 EIP-712 类型
  const types = {
    SafeTx: [
      { name: "to", type: "address" },
      { name: "value", type: "uint256" },
      { name: "data", type: "bytes" },
      { name: "operation", type: "uint8" },
      { name: "safeTxGas", type: "uint256" },
      { name: "baseGas", type: "uint256" },
      { name: "gasPrice", type: "uint256" },
      { name: "gasToken", type: "address" },
      { name: "refundReceiver", type: "address" },
      { name: "nonce", type: "uint256" },
    ],
  };
  // 构造 SafeTx 消息
  const safeTxGas = overrides?.safeTxGas ?? 0n;
  const message = {
    to: contract.target, value: 0n, data,
    operation: 0, // 0 = Call（不是 delegatecall）
    safeTxGas,
    baseGas: 0n, gasPrice: 0n,
    gasToken: ethers.ZeroAddress,
    refundReceiver: ethers.ZeroAddress,
    nonce,
  };
  // 对 Safe 的 owner 按地址升序排序（Safe 验证签名时要求严格递增）
  const sorted = [...safeOwners].sort((a: any, b: any) =>
    a.address.toLowerCase().localeCompare(b.address.toLowerCase())
  );
  // 取前 threshold 个 owner 收集签名
  let sigBytes = "0x";
  for (let i = 0; i < threshold && i < sorted.length; i++) {
    const owner = sorted[i];
    // 使用 EIP-712 签名 SafeTx
    const sig = await owner.signTypedData(domain, types, message);
    const { r, s, v } = ethers.Signature.from(sig);
    // 将 r（32 字节）+ s（32 字节）+ v（1 字节）拼接为 65 字节签名
    sigBytes += ethers.solidityPacked(
      ["bytes32", "bytes32", "uint8"],
      [r, s, v]
    ).slice(2);
  }
  // 如果签名列表为空（不应发生），提供一个空签名字节
  if (sigBytes === "0x") sigBytes = "0x00";
  // 执行 Safe 多签交易
  return safe.execTransaction(
    contract.target, 0n, data, 0, // to, value, data, operation
    safeTxGas, 0n, 0n,            // safeTxGas, baseGas, gasPrice
    ethers.ZeroAddress,            // gasToken
    ethers.ZeroAddress,            // refundReceiver
    sigBytes                       // 拼接的签名数据
  );
}

// ======================== 基础部署函数 ========================

/**
 * deployBase 是测试套件中最核心的工具函数
 * 它完整部署一套测试环境，包括：
 *   1. Uniswap 基础设施（WETH、Factory、Router）
 *   2. 两种测试代币（tokenA、tokenB）并为 Uniswap 添加流动性
 *   3. 预言机合约并设置价格
 *   4. 两种债务代币（lendDebt、borrowDebt）
 *   5. PledgePool 合约并创建一个默认借贷池
 *   6. 为各个测试账号分发代币并授权
 *
 * 这样每个测试用例都可以从干净的状态开始，避免测试间相互影响
 */
async function deployBase(): Promise<PoolContext> {
  // 获取当前最新的区块，用于确定链上时间
  const block = (await ethers.provider.getBlock("latest"))!;
  // 从区块中提取时间戳，作为所有时间计算的基准
  const now = block.timestamp;
  // 获取 Hardhat 自动生成的测试账号列表
  const signers = await ethers.getSigners();
  // 从签名者列表中解构出各个角色的账号
  // 注意：Hardhat 默认提供 20 个账号，足够分配所有角色
  const [owner, lender, lender2, borrower, borrower2, feeCollector, liquidator] = signers;
  // 选择 3 个未使用的账号作为 Safe 多签钱包的 owner（阈值 2/3，即任意 2 人签名即可执行）
  const safeOwner1 = signers[7];
  const safeOwner2 = signers[8];
  const safeOwner3 = signers[9];
  const safeOwners = [safeOwner1, safeOwner2, safeOwner3];
  const safeThreshold = 2;

  // --- 部署 Uniswap V2 基础设施 ---

  // 部署 WETH（Wrapped Ether）合约，Uniswap 路由需要 WETH 地址
  const WETH = await ethers.getContractFactory("WETH");
  const weth = await WETH.deploy();
  // 部署 Uniswap V2 工厂合约，用于创建交易对
  const UniswapV2Factory = await ethers.getContractFactory("UniswapV2Factory");
  const uniFactory = await UniswapV2Factory.deploy(owner.address);
  // 部署 Uniswap V2 Router02 路由合约，用于添加流动性和交易
  const UniswapV2Router02 = await ethers.getContractFactory("UniswapV2Router02");
  const router = await UniswapV2Router02.deploy(uniFactory.target, weth.target);

  // --- 部署测试代币 ---

  // 部署模拟 ERC20 代币 A（lendToken）
  // 构造函数参数：名称、符号、总供应量
  const MockERC20 = await ethers.getContractFactory("MockTestERC20");
  const tokenA = await MockERC20.deploy("LendToken", "LEND", TOKEN_SUPPLY);
  // 部署模拟 ERC20 代币 B（borrowToken）
  const tokenB = await MockERC20.deploy("BorrowToken", "BORR", TOKEN_SUPPLY);

  // 授权 Uniswap 路由器从 owner 账户转移代币，用于添加流动性
  // 授权额度为最大供应量
  await tokenA.approve(router.target, MAX_SUPPLY);
  await tokenB.approve(router.target, MAX_SUPPLY);
  // 在 Uniswap 上为 tokenA/tokenB 交易对添加流动性
  // 各注入 500,000 * 1e18，滑点容忍度为 0，接收方为 owner
  // deadline 设置为 2^255（一个极大值），避免时间戳检查导致交易过期
  await router.addLiquidity(
    tokenA.target, tokenB.target,
    500_000n * 10n ** 18n, 500_000n * 10n ** 18n,
    0, 0, owner.address, 2n ** 255n
  );

  // --- 部署预言机 ---

  // 部署 BscPledgeOracle 预言机合约，owner 为多签钱包地址
  const BscPledgeOracle = await ethers.getContractFactory("BscPledgeOracle");
  const oracle = await BscPledgeOracle.deploy(owner.address);
  // 设置 tokenA 的价格为 1 USD（精度 1e18）
  await oracle.connect(owner).setPrice(tokenA.target, 1n * 10n ** 18n);
  // 设置 tokenB 的价格为 1 USD（精度 1e18）
  await oracle.connect(owner).setPrice(tokenB.target, 1n * 10n ** 18n);

  // --- 部署债务代币 ---

  // 部署出借债务代币 SP Token（代表出借人的收益权）
  const DebtToken = await ethers.getContractFactory("DebtToken");
  const lendDebt = await DebtToken.deploy("SP Token", "sp", owner.address);
  // 部署借款债务代币 JP Token（代表借款人的债务凭证）
  const borrowDebt = await DebtToken.deploy("JP Token", "jp", owner.address);

  // --- 部署 Safe 多签钱包（作为 PledgePool 的 owner） ---

  // 从 Hardhat build-info 加载 Safe 合约 ABI 和 bytecode
  const { safeAbi, safeBytecode, factoryAbi, factoryBytecode } = loadSafeBuildInfo();
  // 部署 Safe 单例合约（master copy），所有 Safe 代理共用此实现
  const Safe = new ethers.ContractFactory(safeAbi, safeBytecode, owner);
  const safeSingleton = await Safe.deploy();
  // 部署 Safe 代理工厂，用于创建各个 Safe 多签实例
  const SafeProxyFactory = new ethers.ContractFactory(factoryAbi, factoryBytecode, owner);
  const proxyFactory = await SafeProxyFactory.deploy() as any;
  // 编码 Safe.setup() 调用数据：配置 owner 列表和签名阈值
  const setupData = safeSingleton.interface.encodeFunctionData("setup", [
    safeOwners.map((o: any) => o.address),  // _owners：多签 owner 地址列表
    safeThreshold,                          // _threshold：最少签名数（2）
    ethers.ZeroAddress,                     // to：委托调用目标（不需要）
    "0x",                                   // data：委托调用数据（空）
    ethers.ZeroAddress,                     // fallbackHandler：回退处理器（不需要）
    ethers.ZeroAddress,                     // paymentToken：支付代币（不需要）
    0n,                                     // payment：支付金额
    ethers.ZeroAddress,                     // paymentReceiver：支付接收方
  ]);
  // 创建 Safe 代理实例（代理合约的地址取决于工厂地址和 nonce，使用随机 nonce 确保唯一性）
  const createTx = await proxyFactory.createProxyWithNonce(
    safeSingleton.target, setupData, BigInt(Math.floor(Math.random() * 1e12))
  );
  const createReceipt = await createTx.wait();
  // 从 ProxyCreation 事件中解析 Safe 代理地址（topics[1] 是 indexed 的 proxy 地址）
  const proxyCreatedLog = createReceipt.logs.find(
    (log: any) => log.address.toLowerCase() === proxyFactory.target.toLowerCase()
  );
  // topics[1] 是 bytes32，取后 20 字节转换为地址
  const safeAddress = ethers.getAddress("0x" + proxyCreatedLog.topics[1].slice(26));
  // 创建 Safe 合约实例，直接使用从 build-info 加载的 ABI
  const safe = new ethers.Contract(safeAddress, safeAbi, owner);

  // --- 部署 PledgePool 合约 ---

  // 部署核心借贷池合约，构造函数传入：预言机、路由、手续费地址、Safe 多签地址（作为 owner）
  const PledgePool = await ethers.getContractFactory("PledgePool");
  const pool = await PledgePool.deploy(oracle.target, router.target, feeCollector.address, safeAddress);
  // 授权 PledgePool 合约可以铸造和销毁 lendDebt 代币
  // 注意：DebtToken 的 owner 是 EOA 部署者，不是 Safe 多签，所以直接用 owner 账号调用
  await lendDebt.connect(owner).setMinter(pool.target, true);
  // 授权 PledgePool 合约可以铸造和销毁 borrowDebt 代币
  await borrowDebt.connect(owner).setMinter(pool.target, true);

  // --- 创建默认借贷池 ---

  // 结算时间设置为 2 年后，给测试用例足够的时间在结算前进行操作
  const settleTime = now + 730 * 86400;
  // 结束时间设置为结算时间后 30 天
  const endTime = settleTime + 30 * 86400;

  // 通过 Safe 多签调用 createPledgePool 创建第一个借贷池（pid = 0）
  // 参数包括：结算时间、结束时间、利率、最大供应量、抵押率、代币地址等
  await safeExec(safe, pool, "createPledgePool", [{
    settleTime, endTime,
    interestRate: INTEREST_RATE,
    maxSupply: MAX_SUPPLY,
    mortgageRate: MORTGAGE_RATE,
    lendToken: tokenA.target,
    borrowToken: tokenB.target,
    lendDebtToken: lendDebt.target,
    borrowDebtToken: borrowDebt.target,
    autoLiquidateThreshold: LIQUIDATE_THRESHOLD,
  }], safeOwners);

  // --- 为测试账号分发代币 ---

  // 给出借人 1 转账 500,000 * 1e18 的 tokenA
  await tokenA.transfer(lender.address, 500_000n * 10n ** 18n);
  // 给出借人 2 转账 500,000 * 1e18 的 tokenA
  await tokenA.transfer(lender2.address, 500_000n * 10n ** 18n);
  // 给借款人 1 转账 500,000 * 1e18 的 tokenB
  await tokenB.transfer(borrower.address, 500_000n * 10n ** 18n);
  // 给借款人 2 转账 500,000 * 1e18 的 tokenB
  await tokenB.transfer(borrower2.address, 500_000n * 10n ** 18n);

  // --- 授权 PledgePool 转移代币 ---

  // 出借人 1 授权 PledgePool 无限额使用 tokenA
  await tokenA.connect(lender).approve(pool.target, ethers.MaxUint256);
  // 出借人 2 授权 PledgePool 无限额使用 tokenA
  await tokenA.connect(lender2).approve(pool.target, ethers.MaxUint256);
  // 借款人 1 授权 PledgePool 无限额使用 tokenB
  await tokenB.connect(borrower).approve(pool.target, ethers.MaxUint256);
  // 借款人 2 授权 PledgePool 无限额使用 tokenB
  await tokenB.connect(borrower2).approve(pool.target, ethers.MaxUint256);

  // 返回构建好的完整上下文，包含所有合约实例、账号信息和池子 ID
  return {
    signers, pool, oracle, weth, router, tokenA, tokenB,
    lendDebt, borrowDebt,
    safe, safeOwners,
    owner, lender, lender2, borrower, borrower2, feeCollector, liquidator,
    pid: 0n, settleTime, endTime,
  };
}

// ============================================================
// 1. 部署 & 初始状态
// 测试部署后合约的基本配置是否正确
// ============================================================

// "Deployment" 测试块：验证 PledgePool 合约部署后的初始状态
describe("PledgePool — Deployment", function () {
  // 声明上下文变量，在 before 钩子中初始化
  let ctx: PoolContext;
  // before 钩子：在所有测试用例之前执行一次部署
  before(async () => { ctx = await deployBase(); });

  // 验证部署后 owner、oracle、swapRouter、feeAddress 等关键字段是否正确设置
  it("should set owner, oracle, swapRouter, feeAddress correctly", async function () {
    // 验证 owner 地址为 Safe 多签合约地址（非 EOA）
    expect(await ctx.pool.owner()).to.equal(await ctx.safe.getAddress());
    // 验证预言机地址是否正确
    expect(await ctx.pool.oracle()).to.equal(ctx.oracle.target);
    // 验证 swap 路由地址是否正确
    expect(await ctx.pool.swapRouter()).to.equal(ctx.router.target);
    // 验证手续费收集地址是否正确
    expect(await ctx.pool.feeAddress()).to.equal(ctx.feeCollector.address);
    // 验证出借手续费初始值为 0
    expect(await ctx.pool.lendFee()).to.equal(0n);
    // 验证借款手续费初始值为 0
    expect(await ctx.pool.borrowFee()).to.equal(0n);
    // 验证最小存入金额初始值是否正确
    expect(await ctx.pool.minAmount()).to.equal(MIN_AMOUNT);
    // 验证全局暂停状态初始为 false（未暂停）
    expect(await ctx.pool.globalPaused()).to.equal(false);
  });
});

// ============================================================
// 2. 管理函数
// 测试 onlyOwner 权限控制、参数校验和状态变更
// ============================================================

// "Admin" 测试块：验证管理员函数的权限控制和功能正确性
describe("PledgePool — Admin", function () {
  let ctx: PoolContext;
  before(async () => { ctx = await deployBase(); });

  // 测试 setFee 函数的权限控制和取值范围校验
  it("setFee — only owner, validates range", async function () {
    // 非 owner 调用 setFee 应该被 OwnableUnauthorizedAccount 错误拒绝
    await expect(ctx.pool.connect(ctx.lender).setFee(1n, 1n))
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
    // 通过 Safe 多签设置出借手续费为 0.5%（5 * 10^6 / 1e8），借款手续费为 0.3%（3 * 10^6 / 1e8）
    await safeExec(ctx.safe, ctx.pool, "setFee", [5n * 10n ** 6n, 3n * 10n ** 6n], ctx.safeOwners);
    // 验证出借手续费已更新
    expect(await ctx.pool.lendFee()).to.equal(5n * 10n ** 6n);
    // 验证借款手续费已更新
    expect(await ctx.pool.borrowFee()).to.equal(3n * 10n ** 6n);
    // 通过 Safe 多签设置出借手续费为 0 应该被拒绝（Safe 内部交易回滚，触发 GS013）
    await expect(safeExec(ctx.safe, ctx.pool, "setFee", [0, 1n], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
    // 通过 Safe 多签设置借款手续费超过 1e8 应该被拒绝
    await expect(safeExec(ctx.safe, ctx.pool, "setFee", [1n, BASE_DECIMAL + 1n], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
  });

  // 测试 setOracle 函数的权限和事件
  it("setOracle", async function () {
    // 通过 Safe 多签设置为零地址应该被拒绝
    await expect(safeExec(ctx.safe, ctx.pool, "setOracle", [ethers.ZeroAddress], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
    // 通过 Safe 多签设置为 lender 地址，验证事件是否正确发出
    const tx = await safeExec(ctx.safe, ctx.pool, "setOracle", [ctx.lender.address], ctx.safeOwners);
    // 验证 SetOracle 事件：old 和 new 地址（通过 Safe 调用，事件仍由 PledgePool 发出）
    await expect(tx).to.emit(ctx.pool, "SetOracle").withArgs(ctx.oracle.target, ctx.lender.address);
    // 验证 oracle 地址已更新
    expect(await ctx.pool.oracle()).to.equal(ctx.lender.address);
  });

  // 测试 setSwapRouter 函数
  it("setSwapRouter", async function () {
    // 通过 Safe 多签将 swap 路由地址设置为 lender 地址
    await safeExec(ctx.safe, ctx.pool, "setSwapRouter", [ctx.lender.address], ctx.safeOwners);
    // 验证路由地址已更新
    expect(await ctx.pool.swapRouter()).to.equal(ctx.lender.address);
  });

  // 测试 setFeeAddress 函数
  it("setFeeAddress", async function () {
    // 通过 Safe 多签将手续费收集地址设置为 lender 地址
    await safeExec(ctx.safe, ctx.pool, "setFeeAddress", [ctx.lender.address], ctx.safeOwners);
    // 验证手续费地址已更新
    expect(await ctx.pool.feeAddress()).to.equal(ctx.lender.address);
  });

  // 测试 setMinAmount 函数
  it("setMinAmount", async function () {
    // 通过 Safe 多签将最小存入金额设置为 200 * 1e18
    await safeExec(ctx.safe, ctx.pool, "setMinAmount", [200n * 10n ** 18n], ctx.safeOwners);
    // 验证最小金额已更新
    expect(await ctx.pool.minAmount()).to.equal(200n * 10n ** 18n);
  });

  // 测试 setGlobalPaused 函数的开关切换
  it("setGlobalPaused toggles state", async function () {
    // 通过 Safe 多签第一次调用：开启暂停
    await safeExec(ctx.safe, ctx.pool, "setGlobalPaused", [], ctx.safeOwners);
    // 验证暂停状态为 true
    expect(await ctx.pool.globalPaused()).to.equal(true);
    // 通过 Safe 多签第二次调用：关闭暂停
    await safeExec(ctx.safe, ctx.pool, "setGlobalPaused", [], ctx.safeOwners);
    // 验证暂停状态恢复为 false
    expect(await ctx.pool.globalPaused()).to.equal(false);
  });
});

// ============================================================
// 3. 池创建
// 测试 createPledgePool 的参数校验和字段初始化
// ============================================================

// "Create Pool" 测试块：验证借贷池创建的参数校验和状态初始化
describe("PledgePool — Create Pool", function () {
  let ctx: PoolContext;
  before(async () => { ctx = await deployBase(); });

  // 测试创建一个新的借贷池并验证字段正确初始化
  it("creates pool with correct fields", async function () {
    // 获取当前区块时间作为创建池子的时间基准
    const { timestamp: now } = (await ethers.provider.getBlock("latest"))!;
    // 结算时间：当前时间 + 1 天（用于快速测试）
    const settleTime = now + 86400;
    // 结束时间：结算时间 + 30 天
    const endTime = settleTime + 30 * 86400;
    // 通过 Safe 多签调用 createPledgePool 创建第二个池子（pid = 1）
    const tx = await safeExec(ctx.safe, ctx.pool, "createPledgePool", [{
      settleTime, endTime,
      interestRate: INTEREST_RATE, maxSupply: MAX_SUPPLY,
      mortgageRate: MORTGAGE_RATE,
      lendToken: ctx.tokenA.target, borrowToken: ctx.tokenB.target,
      lendDebtToken: ctx.lendDebt.target, borrowDebtToken: ctx.borrowDebt.target,
      autoLiquidateThreshold: LIQUIDATE_THRESHOLD,
    }], ctx.safeOwners);
    // 验证 CreatePledgePool 事件正确发出，pid 为 1
    await expect(tx).to.emit(ctx.pool, "CreatePledgePool").withArgs(1n);
    // 读取新创建的池子信息
    const info = await ctx.pool.pledgePoolInfoList(1);
    // 验证结算时间设置正确（索引 0）
    expect(info[0]).to.equal(settleTime);
    // 验证结束时间设置正确（索引 1）
    expect(info[1]).to.equal(endTime);
    // 验证池子状态为 MATCH（0），即初始匹配状态（索引 9）
    expect(info[9]).to.equal(0n);
  });

  // 测试 createPledgePool 的参数验证
  it("reverts invalid params", async function () {
    // 获取当前时间作为时间基准
    const { timestamp: now } = (await ethers.provider.getBlock("latest"))!;
    // 构造一个有效的参数模板，后续测试将分别修改不同字段触发验证
    const base = {
      settleTime: now + 86400,
      endTime: now + 86400 + 30 * 86400,
      interestRate: INTEREST_RATE, maxSupply: MAX_SUPPLY,
      mortgageRate: MORTGAGE_RATE,
      lendToken: ctx.tokenA.target, borrowToken: ctx.tokenB.target,
      lendDebtToken: ctx.lendDebt.target, borrowDebtToken: ctx.borrowDebt.target,
      autoLiquidateThreshold: LIQUIDATE_THRESHOLD,
    };
    // 通过 Safe 多签设置 settleTime 为 0 应该被拒绝（内部回滚，Safe 抛 GS013）
    await expect(safeExec(ctx.safe, ctx.pool, "createPledgePool", [{ ...base, settleTime: 0 }], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
    // 设置 endTime 为 0 应该被拒绝
    await expect(safeExec(ctx.safe, ctx.pool, "createPledgePool", [{ ...base, endTime: 0 }], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
    // 设置 endTime 等于 settleTime 应该被拒绝（结束时间必须在结算时间之后）
    await expect(safeExec(ctx.safe, ctx.pool, "createPledgePool", [{ ...base, endTime: base.settleTime }], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
  });
});

// ============================================================
// 4. 存入
// 测试 lend 函数的正常流程、金额校验和时间限制
// ============================================================

// "Lend" 测试块：验证出借人存入 lendToken 的功能
describe("PledgePool — Lend", function () {
  let ctx: PoolContext;
  before(async () => { ctx = await deployBase(); });

  // 测试正常的 lend 操作：状态更新和事件发送
  it("lend ERC20 — updates state and emits event", async function () {
    // 存入金额：1000 * 1e18 个 tokenA
    const amount = 1000n * 10n ** 18n;
    // 出借人调用 lend 存入 tokenA
    const tx = await ctx.pool.connect(ctx.lender).lend(ctx.pid, amount);
    // 验证 Lend 事件正确发出，包含池子 ID、代币地址、出借人地址和金额
    await expect(tx).to.emit(ctx.pool, "Lend").withArgs(ctx.pid, ctx.tokenA.target, ctx.lender.address, amount);
    // 读取池子信息，验证 lendSupply（索引 4）已更新
    const info = await ctx.pool.pledgePoolInfoList(ctx.pid);
    expect(info[4]).to.equal(amount);
    // 读取出借人信息，验证 lendAmount（索引 0）已更新
    const lendInfo = await ctx.pool.lendInfoMap(ctx.lender.address, ctx.pid);
    expect(lendInfo[0]).to.equal(amount);
  });

  // 测试 lend 低于最小金额时被拒绝
  it("lend ERC20 — reverts below minAmount", async function () {
    // 存入金额 = MIN_AMOUNT - 1，应被拒绝
    await expect(ctx.pool.connect(ctx.lender).lend(ctx.pid, MIN_AMOUNT - 1n))
      .to.be.revertedWith("ERC20 must be greater than minAmount");
  });

  // 测试 lend 超过最大供应量时被拒绝
  it("lend ERC20 — reverts exceeds maxSupply", async function () {
    // 存入金额 = MAX_SUPPLY + 1，超过池子上限，应被拒绝
    await expect(ctx.pool.connect(ctx.lender).lend(ctx.pid, MAX_SUPPLY + 1n))
      .to.be.revertedWith("Exceeds the maximum limit");
  });

  // 测试 lend 在结算时间之后被拒绝
  it("lend ERC20 — reverts after settleTime", async function () {
    // 使用 evm_setNextBlockTimestamp 将下一个区块的时间戳设置为 settleTime + 1
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 需要先手动挖一个块来消费掉上面设置的时间戳，防止影响后续测试
    await ethers.provider.send("evm_mine");
    // 在结算时间后调用 lend 应该被 "Less than this time" 拒绝
    await expect(ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n))
      .to.be.revertedWith("Less than this time");
  });
});

// ============================================================
// 5. 质押
// 测试 borrow 函数的正常流程和时间限制
// ============================================================

// "Borrow" 测试块：验证借款人质押 borrowToken 的功能
describe("PledgePool — Borrow", function () {
  // 测试正常的 borrow 操作：状态更新和事件发送
  it("borrow ERC20 — updates state", async function () {
    // 每个测试用例独立部署，避免测试间状态干扰
    const ctx = await deployBase();
    // 质押金额：2000 * 1e18 个 tokenB
    const amount = 2000n * 10n ** 18n;
    // 借款人调用 borrow 质押 tokenB
    const tx = await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, amount);
    // 验证 Borrow 事件正确发出，包含池子 ID、代币地址、借款人地址和金额
    await expect(tx).to.emit(ctx.pool, "Borrow").withArgs(ctx.pid, ctx.tokenB.target, ctx.borrower.address, amount);
    // 读取池子信息，验证 borrowSupply（索引 5）已更新
    const poolInfo = await ctx.pool.pledgePoolInfoList(ctx.pid);
    expect(poolInfo[5]).to.equal(amount);
  });

  // 测试 borrow 在结算时间之后被拒绝
  it("borrow ERC20 — reverts after settleTime", async function () {
    const ctx = await deployBase();
    // 将下一个区块的时间戳设置为 settleTime + 1
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 手动挖矿消费时间戳设置
    await ethers.provider.send("evm_mine");
    // 结算时间后调用 borrow 应该被 "Less than this time" 拒绝
    await expect(ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 1000n * 10n ** 18n))
      .to.be.revertedWith("Less than this time");
  });
});

// ============================================================
// 6. 结算
// 测试 settlePool 函数的状态转换和金额计算
// ============================================================

// "Settle" 测试块：验证结算流程和状态转换
describe("PledgePool — Settle", function () {
  // 测试正常结算：有 lend 和 borrow 时池子进入 EXECUTION 状态
  it("settlePool transitions to EXECUTION", async function () {
    const ctx = await deployBase();
    // 先让出借人存入 1000 * 1e18 个 tokenA
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 再让借款人质押 2000 * 1e18 个 tokenB
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 将时间推进到结算时间之后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 通过 Safe 多签调用 settlePool 进行结算
    const tx = await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    // 验证 SettlePool 事件正确发出
    await expect(tx).to.emit(ctx.pool, "SettlePool");
    // 验证池子状态变为 EXECUTION（1）
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(1n);
    // 读取结算数据信息
    const dataInfo = await ctx.pool.poolDataInfoList(ctx.pid);
    // 验证结算时的 lend 金额正确（索引 0）
    expect(dataInfo[0]).to.equal(1000n * 10n ** 18n);
    // 验证结算时的 borrow 金额正确（索引 1）
    expect(dataInfo[1]).to.equal(2000n * 10n ** 18n);
  });

  // 测试零供应量结算：没有 lend 或 borrow 时池子进入 UNDONE 状态
  it("settlePool — zero supply → UNDONE", async function () {
    const ctx = await deployBase();
    // 将时间推进到结算时间之后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 通过 Safe 多签直接结算（没有 lend 和 borrow，供应量为 0）
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    // 验证池子状态变为 UNDONE（4）
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(4n);
  });

  // 测试在结算时间之前调用 settlePool 被拒绝
  it("settlePool reverts before settleTime", async function () {
    const ctx = await deployBase();
    // 当前时间还在结算时间之前，通过 Safe 调用会被内部拒绝（Safe 抛 GS013）
    await expect(safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
  });
});

// ============================================================
// 7. 领取凭证
// 测试出借人和借款人领取债务代币凭证
// ============================================================

// "Claim" 测试块：验证用户领取债务代币的能力
describe("PledgePool — Claim", function () {
  let ctx: PoolContext;
  // before 钩子：部署、存入、质押、推进时间、结算，准备领取条件
  before(async () => {
    ctx = await deployBase();
    // 出借人存入 1000 * 1e18
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 借款人质押 2000 * 1e18
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 推进时间到结算后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 通过 Safe 多签执行结算，使池子进入 EXECUTION 状态
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
  });

  // 出借人领取 lendDebtToken（SP Token）
  it("claimLendDebtToken", async function () {
    // 出借人调用 claimLendDebtToken 领取债权凭证
    await ctx.pool.connect(ctx.lender).claimLendDebtToken(ctx.pid);
    // 验证出借人收到的 SP Token 数量等于存入金额 1000 * 1e18
    expect(await ctx.lendDebt.balanceOf(ctx.lender.address)).to.equal(1000n * 10n ** 18n);
  });

  // 借款人领取 borrowDebtToken（JP Token）
  it("claimBorrow", async function () {
    // 借款人调用 claimBorrow 领取债务凭证
    await ctx.pool.connect(ctx.borrower).claimBorrow(ctx.pid);
    // 验证借款人收到的 JP Token 数量等于结算时确认的金额 2000 * 1e18
    expect(await ctx.borrowDebt.balanceOf(ctx.borrower.address)).to.equal(2000n * 10n ** 18n);
  });
});

// ============================================================
// 8. finishPool — same token
// 测试 lendToken 和 borrowToken 为同一代币时的 finish 流程
// ============================================================

// "FinishPool (same token)" 测试块：验证同币种池子的 finish 流程
describe("PledgePool — FinishPool (same token)", function () {
  let ctx: PoolContext;
  let pid: bigint, settleTime: number, endTime: number;
  // before 钩子：创建同币种池子（lendToken = borrowToken = tokenA）
  before(async () => {
    ctx = await deployBase();
    // 获取当前时间用于设置结算和结束时间
    const { timestamp: now } = (await ethers.provider.getBlock("latest"))!;
    // 结算时间：1 天后（用于快速测试）
    settleTime = now + 86400;
    // 结束时间：结算后 30 天
    endTime = settleTime + 30 * 86400;
    // 创建同币种池子（lend 和 borrow 都是 tokenA）
    await safeExec(ctx.safe, ctx.pool, "createPledgePool", [{
      settleTime, endTime,
      interestRate: INTEREST_RATE, maxSupply: MAX_SUPPLY,
      mortgageRate: MORTGAGE_RATE,
      lendToken: ctx.tokenA.target, borrowToken: ctx.tokenA.target,
      lendDebtToken: ctx.lendDebt.target, borrowDebtToken: ctx.borrowDebt.target,
      autoLiquidateThreshold: LIQUIDATE_THRESHOLD,
    }], ctx.safeOwners);
    // 新创建的池子 ID 为 1（因为默认部署已创建了 pid=0 的池子）
    pid = 1n;
    // 由于同币种池子使用 tokenA 作为抵押品，需要给借款人转账 tokenA
    await ctx.tokenA.connect(ctx.owner).transfer(ctx.borrower.address, 500_000n * 10n ** 18n);
    // 出借人授权池子使用 tokenA
    await ctx.tokenA.connect(ctx.lender).approve(ctx.pool.target, ethers.MaxUint256);
    // 借款人授权池子使用 tokenA
    await ctx.tokenA.connect(ctx.borrower).approve(ctx.pool.target, ethers.MaxUint256);
    // 出借人存入 1000 * 1e18 个 tokenA
    await ctx.pool.connect(ctx.lender).lend(pid, 1000n * 10n ** 18n);
    // 借款人质押 2000 * 1e18 个 tokenA
    await ctx.pool.connect(ctx.borrower).borrow(pid, 2000n * 10n ** 18n);
    // 推进时间到结算后并结算
    await ethers.provider.send("evm_setNextBlockTimestamp", [settleTime + 1]);
    // 通过 Safe 多签执行结算
    await safeExec(ctx.safe, ctx.pool, "settlePool", [pid], ctx.safeOwners);
  });

  // 测试同币种池子的 finish 操作
  it("finishPool same-token sets FINISH state", async function () {
    // 推进时间到结束时间后
    await ethers.provider.send("evm_setNextBlockTimestamp", [endTime + 1]);
    // 通过 Safe 多签调用 finishPool 完成池子
    await safeExec(ctx.safe, ctx.pool, "finishPool", [pid], ctx.safeOwners);
    // 验证池子状态变为 FINISH（2）
    expect(await ctx.pool.getPoolState(pid)).to.equal(2n);
    // 读取结算数据信息
    const dataInfo = await ctx.pool.poolDataInfoList(pid);
    // 验证 finishAmountLend（索引 2）大于本金，说明有收益
    expect(dataInfo[2]).to.be.gt(1000n * 10n ** 18n);
  });
});

// ============================================================
// 9. finishPool — different tokens
// 测试 lendToken 和 borrowToken 为不同代币时的 finish 流程
// ============================================================

// "FinishPool (different tokens)" 测试块：验证不同币种池子的 finish 流程（含 DEX 兑换）
describe("PledgePool — FinishPool (different tokens)", function () {
  let ctx: PoolContext;
  // before 钩子：设置手续费、存入、质押、结算
  before(async () => {
    ctx = await deployBase();
    // 设置出借手续费 1%（1 * 10^7 / 1e8）和借款手续费 0.5%
    await safeExec(ctx.safe, ctx.pool, "setFee", [1n * 10n ** 7n, 5n * 10n ** 6n], ctx.safeOwners);
    // 出借人存入 1000 * 1e18
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 借款人质押 2000 * 1e18
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 推进时间到结算后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 通过 Safe 多签执行结算
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
  });

  // 测试在结束时间之前调用 finishPool 被拒绝
  // 注意：这个测试必须在 "finishPool with swap" 之前运行
  // 因为 finish 操作会改变池子状态，导致状态检查失败而非时间检查失败
  it("finishPool reverts before endTime", async function () {
    // 当前时间还在结束时间之前，通过 Safe 调用会被内部拒绝（Safe 抛 GS013）
    await expect(safeExec(ctx.safe, ctx.pool, "finishPool", [ctx.pid], ctx.safeOwners, undefined, { safeTxGas: 0n }))
      .to.be.revertedWith("GS013");
  });

  // 测试不同币种池子的 finish 操作（涉及 DEX 兑换）
  it("finishPool with swap — emits event and calculates correctly", async function () {
    // 推进时间到结束时间后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.endTime + 1]);
    // 通过 Safe 多签调用 finishPool，内部会将 borrowToken 通过 Uniswap 兑换为 lendToken
    const tx = await safeExec(ctx.safe, ctx.pool, "finishPool", [ctx.pid], ctx.safeOwners);
    // 验证 FinishPool 事件正确发出
    await expect(tx).to.emit(ctx.pool, "FinishPool");
    // 验证池子状态变为 FINISH（2）
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(2n);
    // 读取结算数据信息
    const dataInfo = await ctx.pool.poolDataInfoList(ctx.pid);
    // 验证 finishAmountLend（索引 2）大于本金 1000 * 1e18，说明有利息收益
    expect(dataInfo[2]).to.be.gt(1000n * 10n ** 18n);
    // 验证 finishAmountBorrow（索引 3）小于总质押 2000 * 1e18，说明部分用于兑换和手续费
    expect(dataInfo[3]).to.be.lt(2000n * 10n ** 18n);
  });
});

// ============================================================
// 10. 销毁凭证取回资产
// 测试销毁债务代币并取回对应资产
// ============================================================

// "Destroy tokens" 测试块：验证销毁 SP/JP Token 并取回资产
describe("PledgePool — Destroy tokens", function () {
  let ctx: PoolContext;
  // before 钩子：完成完整的生命周期（存入、质押、结算、领取、finish）
  before(async () => {
    ctx = await deployBase();
    // 设置手续费
    await safeExec(ctx.safe, ctx.pool, "setFee", [1n * 10n ** 7n, 5n * 10n ** 6n], ctx.safeOwners);
    // 出借人存入 1000 * 1e18
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 借款人质押 2000 * 1e18
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 推进时间并结算（通过 Safe 多签）
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    // 出借人领取 SP Token
    await ctx.pool.connect(ctx.lender).claimLendDebtToken(ctx.pid);
    // 借款人领取 JP Token
    await ctx.pool.connect(ctx.borrower).claimBorrow(ctx.pid);
    // 推进时间到结束时间后并 finish（通过 Safe 多签）
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.endTime + 1]);
    await safeExec(ctx.safe, ctx.pool, "finishPool", [ctx.pid], ctx.safeOwners);
  });

  // 测试出借人销毁 SP Token 取回本金 + 利息
  it("destroyLendDebtToken — lender receives principal + interest", async function () {
    // 记录销毁前的 tokenA 余额
    const balanceBefore = await ctx.tokenA.balanceOf(ctx.lender.address);
    // 获取出借人持有的 SP Token 数量
    const spBalance = await ctx.lendDebt.balanceOf(ctx.lender.address);
    // 销毁所有 SP Token，取回对应 lendToken
    await ctx.pool.connect(ctx.lender).destroyLendDebtToken(ctx.pid, spBalance);
    // 计算销毁后的余额变化
    const balanceAfter = await ctx.tokenA.balanceOf(ctx.lender.address);
    // 验证取回金额大于 0（有利息收益）
    expect(balanceAfter - balanceBefore).to.be.gt(0n);
    // 验证 SP Token 已全部销毁
    expect(await ctx.lendDebt.balanceOf(ctx.lender.address)).to.equal(0n);
  });

  // 测试借款人销毁 JP Token 取回剩余抵押物
  it("destroyBorrowDebtToken — borrower recovers collateral", async function () {
    // 获取借款人持有的 JP Token 数量
    const jpBalance = await ctx.borrowDebt.balanceOf(ctx.borrower.address);
    // 销毁所有 JP Token，取回剩余抵押物
    await ctx.pool.connect(ctx.borrower).destroyBorrowDebtToken(ctx.pid, jpBalance);
    // 验证 JP Token 已全部销毁
    expect(await ctx.borrowDebt.balanceOf(ctx.borrower.address)).to.equal(0n);
  });
});

// ============================================================
// 11. 清算
// 测试资产价格下跌时的清算流程
// ============================================================

// "Liquidation" 测试块：验证清算机制
describe("PledgePool — Liquidation", function () {
  let ctx: PoolContext;
  // before 钩子：存入、质押、结算，然后人为压低 tokenB 价格触发清算条件
  before(async () => {
    ctx = await deployBase();
    // 设置手续费
    await safeExec(ctx.safe, ctx.pool, "setFee", [1n * 10n ** 7n, 5n * 10n ** 6n], ctx.safeOwners);
    // 出借人存入
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 借款人质押
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 推进时间并结算（通过 Safe 多签）
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    // 通过预言机将 tokenB 的价格从 1 USD 压低到 0.05 USD（5 * 10^16 / 1e18）
    // 这样借款人质押的 tokenB 价值大幅缩水，触发清算条件
    await ctx.oracle.connect(ctx.owner).setPrice(ctx.tokenB.target, 5n * 10n ** 16n);
  });

  // 验证 checkCanLiquidate 返回 true
  it("checkCanLiquidate returns true when collateral drops", async function () {
    // 当抵押品价值低于清算阈值时，checkCanLiquidate 应返回 true
    expect(await ctx.pool.checkCanLiquidate(ctx.pid)).to.equal(true);
  });

  // 测试 liquidatePool 将池子状态变为 LIQUIDATION
  it("liquidatePool sets LIQUIDATION state", async function () {
    // 清算人执行清算
    const tx = await ctx.pool.connect(ctx.liquidator).liquidatePool(ctx.pid);
    // 验证 LiquidatePool 事件发出
    await expect(tx).to.emit(ctx.pool, "LiquidatePool");
    // 验证池子状态变为 LIQUIDATION（3）
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(3n);
    // 读取结算数据信息
    const dataInfo = await ctx.pool.poolDataInfoList(ctx.pid);
    // 验证 liquidationAmountLend（索引 4）大于 0，说明有出借资产被清算
    expect(dataInfo[4]).to.be.gt(0n);
  });
});

// ============================================================
// 12. 紧急提取
// 测试池子处于 UNDONE 状态时的紧急提取功能
// ============================================================

// "Emergency withdraw" 测试块：验证紧急提取功能
describe("PledgePool — Emergency withdraw", function () {
  // 测试出借人在 UNDONE 状态下的紧急提取
  it("emergencyWithdrawLend in UNDONE state", async function () {
    const ctx = await deployBase();
    // 出借人先存入 1000 * 1e18
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 推进时间到结算后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 通过 Safe 多签结算（由于没有借款人，池子进入 UNDONE 状态）
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    // 验证池子状态为 UNDONE（4）
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(4n);
    // 出借人调用 emergencyWithdrawLend 紧急取回资金
    await ctx.pool.connect(ctx.lender).emergencyWithdrawLend(ctx.pid);
    // 验证出借人收回了全部 500,000 * 1e18 个 tokenA
    // （初始 500,000 - 存入 1,000 + 取回 1,000 = 500,000）
    expect(await ctx.tokenA.balanceOf(ctx.lender.address)).to.equal(500_000n * 10n ** 18n);
  });

  // 测试借款人在 UNDONE 状态下的紧急提取
  it("emergencyWithdrawBorrow in UNDONE state", async function () {
    const ctx = await deployBase();
    // 借款人先质押 2000 * 1e18
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 推进时间到结算后
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    // 通过 Safe 多签结算（由于没有出借人，池子进入 UNDONE 状态）
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    // 验证池子状态为 UNDONE（4）
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(4n);
    // 借款人调用 emergencyWithdrawBorrow 紧急取回抵押品
    await ctx.pool.connect(ctx.borrower).emergencyWithdrawBorrow(ctx.pid);
    // 验证借款人收回了全部 500,000 * 1e18 个 tokenB
    expect(await ctx.tokenB.balanceOf(ctx.borrower.address)).to.equal(500_000n * 10n ** 18n);
  });

  // 测试在非 UNDONE 状态下调用紧急提取被拒绝
  it("emergencyWithdraw reverts in non-UNDONE state", async function () {
    const ctx = await deployBase();
    // 出借人存入（池子仍处于 MATCH 状态）
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 在 MATCH 状态下调用 emergencyWithdrawLend 应该被拒绝
    await expect(ctx.pool.connect(ctx.lender).emergencyWithdrawLend(ctx.pid))
      .to.be.revertedWith("Pool state must be UNDONE");
  });
});

// ============================================================
// 13. 完整生命周期
// 测试从存放到销毁的完整流程
// ============================================================

// "Full lifecycle" 测试块：验证完整的借贷生命周期
describe("PledgePool — Full lifecycle", function () {
  it("lend → borrow → settle → claim → finish → destroy", async function () {
    const ctx = await deployBase();
    // 设置出借手续费 1% 和借款手续费 0.5%
    await safeExec(ctx.safe, ctx.pool, "setFee", [1n * 10n ** 7n, 5n * 10n ** 6n], ctx.safeOwners);

    // --- 阶段 1：存入 ---
    // 出借人 1 存入 1000 * 1e18
    await ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n);
    // 出借人 2 存入 500 * 1e18
    await ctx.pool.connect(ctx.lender2).lend(ctx.pid, 500n * 10n ** 18n);
    // 借款人 1 质押 2000 * 1e18
    await ctx.pool.connect(ctx.borrower).borrow(ctx.pid, 2000n * 10n ** 18n);
    // 借款人 2 质押 1000 * 1e18
    await ctx.pool.connect(ctx.borrower2).borrow(ctx.pid, 1000n * 10n ** 18n);

    // 验证总存入和总质押量正确
    const preInfo = await ctx.pool.pledgePoolInfoList(ctx.pid);
    expect(preInfo[4]).to.equal(1500n * 10n ** 18n); // lendSupply
    expect(preInfo[5]).to.equal(3000n * 10n ** 18n); // borrowSupply

    // --- 阶段 2：结算（通过 Safe 多签） ---
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.settleTime + 1]);
    await safeExec(ctx.safe, ctx.pool, "settlePool", [ctx.pid], ctx.safeOwners);
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(1n); // EXECUTION

    // --- 阶段 3：领取凭证 ---
    await ctx.pool.connect(ctx.lender).claimLendDebtToken(ctx.pid);
    await ctx.pool.connect(ctx.lender2).claimLendDebtToken(ctx.pid);
    await ctx.pool.connect(ctx.borrower).claimBorrow(ctx.pid);
    await ctx.pool.connect(ctx.borrower2).claimBorrow(ctx.pid);

    // --- 阶段 4：finish（通过 Safe 多签） ---
    await ethers.provider.send("evm_setNextBlockTimestamp", [ctx.endTime + 1]);
    await safeExec(ctx.safe, ctx.pool, "finishPool", [ctx.pid], ctx.safeOwners);
    expect(await ctx.pool.getPoolState(ctx.pid)).to.equal(2n); // FINISH

    // --- 阶段 5：销毁凭证取回资产 ---
    const sp1 = await ctx.lendDebt.balanceOf(ctx.lender.address);
    const sp2 = await ctx.lendDebt.balanceOf(ctx.lender2.address);
    await ctx.pool.connect(ctx.lender).destroyLendDebtToken(ctx.pid, sp1);
    await ctx.pool.connect(ctx.lender2).destroyLendDebtToken(ctx.pid, sp2);

    const jp1 = await ctx.borrowDebt.balanceOf(ctx.borrower.address);
    const jp2 = await ctx.borrowDebt.balanceOf(ctx.borrower2.address);
    await ctx.pool.connect(ctx.borrower).destroyBorrowDebtToken(ctx.pid, jp1);
    await ctx.pool.connect(ctx.borrower2).destroyBorrowDebtToken(ctx.pid, jp2);

    // 验证所有债务代币已销毁完毕
    expect(await ctx.lendDebt.totalSupply()).to.equal(0n);
    expect(await ctx.borrowDebt.totalSupply()).to.equal(0n);
    // 验证手续费已被收取（feeCollector 账户有余额）
    expect(await ctx.tokenA.balanceOf(ctx.feeCollector.address)).to.be.gt(0n);
  });
});

// ============================================================
// 14. 校验函数 & 边界
// 测试各种校验函数和边界条件
// ============================================================

// "Check functions & edge cases" 测试块：验证辅助校验函数和各种边界情况
describe("PledgePool — Check functions & edge cases", function () {
  let ctx: PoolContext;
  before(async () => { ctx = await deployBase(); });

  // 验证 checkCanSettle 在结算时间之前返回 false
  it("checkCanSettle returns false before settleTime", async function () {
    // 当前时间在结算时间之前，settle 条件不满足
    expect(await ctx.pool.checkCanSettle(ctx.pid)).to.equal(false);
  });

  // 验证 checkCanFinish 在 MATCH 状态下返回 false
  it("checkCanFinish returns false for MATCH pool", async function () {
    // 池子还处于 MATCH 状态（未结算），finish 条件不满足
    expect(await ctx.pool.checkCanFinish(ctx.pid)).to.equal(false);
  });

  // 验证对不存在的池子 ID 进行操作时触发 panic
  it("reverts for non-existent pool", async function () {
    // 对不存在的池子 ID（999）调用 settlePool 会导致数组越界
    // Solidity 的数组越界访问会触发 panic(0x32)，而不是普通的 revert
    // 使用 provider.call 模拟从 Safe 地址调用，绕过 onlyOwner 检查并捕获 panic
    const safeAddr = await ctx.safe.getAddress();
    const data = ctx.pool.interface.encodeFunctionData("settlePool", [999n]);
    // provider.call 返回原始结果或抛出包含 revert 数据的异常
    await expect(
      ethers.provider.call({ from: safeAddr, to: ctx.pool.target, data })
    ).to.be.revertedWithPanic(0x32);
  });

  // 验证只有 owner 才能调用的管理函数都被正确保护
  it("onlyOwner functions revert for non-owner", async function () {
    // 非 owner 调用 setOracle 被拒绝
    await expect(ctx.pool.connect(ctx.lender).setOracle(ctx.lender.address))
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
    // 非 owner 调用 setSwapRouter 被拒绝
    await expect(ctx.pool.connect(ctx.lender).setSwapRouter(ctx.lender.address))
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
    // 非 owner 调用 setGlobalPaused 被拒绝
    await expect(ctx.pool.connect(ctx.lender).setGlobalPaused())
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
    // 非 owner 调用 settlePool 被拒绝
    await expect(ctx.pool.connect(ctx.lender).settlePool(ctx.pid))
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
    // 非 owner 调用 finishPool 被拒绝
    await expect(ctx.pool.connect(ctx.lender).finishPool(ctx.pid))
      .to.be.revertedWithCustomError(ctx.pool, "OwnableUnauthorizedAccount");
  });

  // 验证 lend 传入 0 金额被拒绝
  it("lend with zero amount reverts", async function () {
    // 存入 0 金额应该被拒绝
    await expect(ctx.pool.connect(ctx.lender).lend(ctx.pid, 0n))
      .to.be.revertedWith("ERC20 must be greater than 0");
  });

  // 验证全局暂停时所有操作被拒绝
  it("reverts when paused", async function () {
    // 通过 Safe 多签开启全局暂停
    await safeExec(ctx.safe, ctx.pool, "setGlobalPaused", [], ctx.safeOwners);
    // 暂停后调用 lend 应该被 "Global Paused" 拒绝
    await expect(ctx.pool.connect(ctx.lender).lend(ctx.pid, 1000n * 10n ** 18n))
      .to.be.revertedWith("Global Paused");
  });
});
