# Pledge Schedule 定时任务模块技术文档

## 一、模块概述

`schedule` 是 Pledge 项目的**链上数据同步引擎**，作为独立服务运行（`pledge-task`）。其核心职责是：

1. **定时从 BSC 区块链读取数据**（资金池、价格、代币信息）
2. **写入 MySQL 和 Redis**，供 API 服务查询
3. **向链上写入数据**（管理员签名设置预言机价格）
4. **异常监控告警**（合约余额不足时邮件通知）

### 数据流全景

```
┌──────────────────────────────────────────────────────────────────────────┐
│                          BSC Blockchain                                  │
│                                                                          │
│  PledgePool 合约                  BscPledgeOracle 合约    ERC20 代币合约   │
│  ├─ PoolLength() (只读)           ├─ GetPrice() (只读)    ├─ symbol()     │
│  ├─ PoolBaseInfo() (只读)         └─ SetPrice() (写入)◄───┘               │
│  ├─ PoolDataInfo() (只读)                                   (管理员私钥签名)│
│  ├─ BorrowFee() (只读)                                                │
│  └─ LendFee() (只读)                                                 │
└──────────────────────────────┬───────────────────────────────────────────┘
                               │
                    ethclient.Dial() RPC 调用
                               │
                               ▼
┌──────────────────────────────────────────────────────────────────────────┐
│                      schedule（pledge-task 定时任务）                       │
│                                                                          │
│  ┌─────────────────┐   ┌──────────────────┐   ┌──────────────────┐      │
│  │  poolService    │   │  tokenPriceService│   │  tokenSymbolService│    │
│  │  ← PledgePool   │   │  ← Oracle 合约    │   │  ← ERC20 合约    │      │
│  │  每 2 分钟       │   │  每 1 分钟         │   │  每 2 小时        │      │
│  └───────┬─────────┘   └──────┬───────────┘   └──────┬───────────┘      │
│          │                    │                       │                  │
│          ▼                    ▼                       ▼                  │
│  ┌──────────────────────────────────────────────────────────────────┐   │
│  │                     MySQL + Redis                                 │   │
│  │  表: poolbases, pooldata, token_info                              │   │
│  │  Redis Key 前缀: base_info:, data_info:, token_info:             │   │
│  └──────────────────────────────────────────────────────────────────┘   │
│                                                                          │
│  ┌──────────────────┐     ┌─────────────────────┐                       │
│  │  tokenLogoService│     │  balanceMonitor     │                       │
│  │  ← PancakeSwap   │     │  ← BSC 节点         │                       │
│  │  每 2 小时        │     │  每 30 分钟 → 邮件告警 │                      │
│  └──────┬───────────┘     └─────────────────────┘                       │
│         ▼                                                               │
│    token_info 表                                                         │
└──────────────────────────────────────────────────────────────────────────┘
                               │
                               ▼
                    API 服务 (pledge-api)
                    读取 MySQL/Redis 提供 HTTP API
```

---

## 二、目录结构

```
schedule/
├── pledge_task.go          # 服务入口 main()
├── pledge-task.service     # Systemd 服务文件 (Ubuntu)
├── README.md               # gocron 使用说明
├── common/
│   └── common.go           # 环境变量读取 (管理员私钥)
├── models/
│   ├── table_init.go       # GORM AutoMigrate 表结构初始化
│   ├── poolBase.go         # PoolBase 模型 + SavePoolBase 逻辑
│   ├── poolData.go         # PoolData 模型 + SavePoolData 逻辑
│   ├── tokenInfo.go        # TokenInfo 模型 + 缓存查询
│   ├── tokenLogo.go        # PancakeSwap 代币列表结构体
│   ├── redisToken.go       # Redis 缓存结构体
│   └── abi.go              # ABI 下载响应结构体
├── services/
│   ├── poolService.go      # 资金池数据同步 (PledgePool 合约)
│   ├── tokenPriceService.go # 代币价格同步/写入 (BscPledgeOracle 合约)
│   ├── tokenSymbolService.go# 代币 Symbol 同步 (ERC20 合约)
│   ├── tokenLogoService.go  # 代币 Logo 同步 (PancakeSwap + 本地)
│   └── balanceMonitor.go   # 合约余额监控告警
└── tasks/
    └── task.go             # 任务调度器 (gocron)
```

---

## 三、启动流程

```
pledge_task.go main()
    │
    ├── 1. db.InitMysql()       ─── 连接 MySQL
    ├── 2. db.InitRedis()       ─── 连接 Redis
    ├── 3. models.InitTable()   ─── 自动建表 (4 张表)
    └── 4. tasks.Task()         ─── 启动任务调度
                │
                ├── common.GetEnv()          ── 读取环境变量 plgr_admin_private_key
                ├── db.RedisFlushDB()        ── 清空 Redis 缓存
                │
                ├── [首次立即执行] ───────────────────────────────┐
                │   ├── UpdateAllPoolInfo()     ← 同步资金池      │
                │   ├── UpdateContractPrice()   ← 同步价格        │
                │   ├── UpdateContractSymbol()  ← 同步 Symbol     │
                │   ├── UpdateTokenLogo()       ← 同步 Logo       │
                │   ├── Monitor()               ← 余额检查        │
                │   └── SavePlgrPriceTestNet()   ← 写入 PLGR 价格 │
                │                                                │
                └── [gocron 定时调度] ────────────────────────────┘
                     ├── 每 1 分钟  → UpdateContractPrice()
                     ├── 每 2 分钟  → UpdateAllPoolInfo()
                     ├── 每 30 分钟 → Monitor()
                     ├── 每 30 分钟 → SavePlgrPriceTestNet()
                     └── 每 2 小时  → UpdateContractSymbol()
                                    → UpdateTokenLogo()
```

---

## 四、六大核心任务详解

### 4.1 poolService — 资金池数据同步

**文件**：`services/poolService.go`

**执行频率**：每 2 分钟

**功能**：从链上 PledgePool 合约读取所有资金池的静态信息和动态数据，同步到 MySQL。

**详细流程**：

```
UpdateAllPoolInfo()
    │
    └── UpdatePoolInfo(contractAddr, netUrl, chainId)
        │  当前: 仅同步测试网 (chain_id=97)
        │  主网 (chain_id=56) 代码被注释
        │
        ├── ethclient.Dial(netUrl)           ← 连接 BSC 节点 RPC
        ├── bindings.NewPledgePoolToken(...) ← abigen 生成的合约绑定
        │
        ├── pledgePoolToken.BorrowFee(nil)   ← 读取全局借款费率
        ├── pledgePoolToken.LendFee(nil)     ← 读取全局出借费率
        └── pledgePoolToken.PoolLength(nil)  ← 获取资金池总数 N
            │
            └── for i := 0; i < N; i++
                │
                ├── 1. PoolBaseInfo(i)       ← 读链: 第 i 个池的基础信息
                │   ├─ settleTime, endTime
                │   ├─ interestRate, maxSupply
                │   ├─ lendSupply, borrowSupply
                │   ├─ martgageRate
                │   ├─ lendToken, borrowToken (地址)
                │   ├─ spCoin, jpCoin
                │   ├─ autoLiquidateThreshold
                │   └─ state (0=未开始, 1=进行中, 2=已清算, ...)
                │
                ├── 2. GetTokenInfo()         ← 查询 Redis/MySQL 获取借贷代币的 Symbol/Logo/Price
                │
                ├── 3. MD5 去重判断           ← 计算 MD5 与 Redis 缓存比对
                │   key: "base_info:pool_{chainId}_{poolId}"
                │
                ├── 4. SavePoolBase()         ← 写库: poolbases 表
                │   ├─ 自动创建/更新 token_info (借贷代币)
                │   └─ INSERT or UPDATE poolbases
                │
                ├── 5. PoolDataInfo(i)        ← 读链: 第 i 个池的动态数据
                │   ├─ finishAmountBorrow
                │   ├─ finishAmountLend
                │   ├─ liquidationAmounBorrow
                │   ├─ liquidationAmounLend
                │   ├─ settleAmountBorrow
                │   └─ settleAmountLend
                │
                ├── 6. MD5 去重判断
                │   key: "data_info:pool_{chainId}_{poolId}"
                │
                └── 7. SavePoolData()         ← 写库: pooldata 表
                    └─ INSERT or UPDATE pooldata
```

**涉及 MySQL 表**：
- `poolbases` — 资金池基础信息
- `pooldata` — 资金池动态数据
- `token_info` — 自动创建借贷代币记录

**MD5 去重机制**：
```go
func (s *poolService) GetPoolMd5(baseInfo *models.PoolBase, key string) (bool, string, string) {
    baseInfoBytes, _ := json.Marshal(baseInfo)
    baseInfoMd5Str := utils.Md5(string(baseInfoBytes))
    resInfoBytes, _ := db.RedisGet(key)
    // 比较 Redis 中缓存的 MD5 与当前数据的 MD5
    // 相同 → 跳过写入; 不同 → 写入 MySQL 并更新 Redis MD5
}
```

---

### 4.2 tokenPriceService — 代币价格同步 & 预言机价格写入

**文件**：`services/tokenPriceService.go`

#### 4.2.1 UpdateContractPrice — 读取价格

**执行频率**：每 1 分钟

**功能**：从 BscPledgeOracle 预言机合约读取所有代币价格，同步到 MySQL。

```
UpdateContractPrice()
    │
    ├── db.Mysql.Find(&tokens)           ← 从 token_info 表读取所有代币
    │
    └── for _, t := range tokens
        │
        ├── GetTestNetTokenPrice(t.Token)   ← (测试网) 调用预言机合约
        │   ├── ethclient.Dial(TestNet.NetUrl)
        │   ├── bindings.NewBscPledgeOracleTestnetToken(addr, conn)
        │   └── oracle.GetPrice(nil, tokenAddr) → 返回价格 (uint256)
        │
        ├── GetMainNetTokenPrice(t.Token)   ← (主网) 代码被注释，未启用
        │
        ├── CheckPriceData()               ← Redis 去重
        │   key: "token_info:{chainId}:{token}"
        │
        └── SavePriceData()                ← 更新 MySQL token_info.price
            └── UPDATE token_info SET price = ? WHERE token = ? AND chain_id = ?
```

#### 4.2.2 SavePlgrPrice / SavePlgrPriceTestNet — 写入价格

**执行频率**：每 30 分钟（仅测试网启用）

**功能**：使用管理员私钥签名交易，调用预言机合约的 `SetPrice` 方法，将 PLGR 代币价格写入链上。

```
SavePlgrPriceTestNet()
    │
    ├── ethclient.Dial(TestNet.NetUrl)                           ← 连接 BSC 节点
    ├── bindings.NewBscPledgeOracleMainnetToken(addr, conn)      ← 绑定预言机合约
    │
    ├── crypto.HexToECDSA(plgrAdminPrivateKey)                   ← 从环境变量加载私钥
    ├── bind.NewKeyedTransactorWithChainID(pkey, chainId)        ← 创建签名器
    │
    ├── 构造 TransactOpts{From, Signer, Context(5s超时)}         ← 交易参数
    │
    ├── oracle.SetPrice(&opts, plgrAddress, big.NewInt(22222))   ← 发送交易！！
    │  (这是有 Gas 费的状态修改交易，不是只读调用)
    │
    ├── 记录日志 "SavePlgrPrice" + error
    │
    └── GetTestNetTokenPrice(plgrAddress)                        ← 读取写入后的价格验证
```

**重要说明**：
- 测试网固定写死价格为 `22222`（硬编码，不是真实价格）
- 主网 `SavePlgrPrice` 从 Redis 读取 Kucoin 的 PLGR-USDT 价格再乘以 `10^8` 后写入（当前被注释）
- 写入链上后**不会写回 MySQL**，MySQL 中的价格由 `UpdateContractPrice` 定时读取更新

---

### 4.3 tokenSymbolService — 代币 Symbol 同步

**文件**：`services/tokenSymbolService.go`

**执行频率**：每 2 小时

**功能**：从链上 ERC20 代币合约读取 `symbol()`，写入 MySQL token_info 表。

```
UpdateContractSymbol()
    │
    ├── db.Mysql.Find(&tokens)           ← 从 token_info 表读取所有代币
    │
    └── for _, t := range tokens
        │
        ├── [测试网] GetContractSymbolOnTestNet(t.Token)
        │   ├── abifile.GetAbiByToken("erc20")    ← 使用本地 erc20.abi
        │   ├── bind.NewBoundContract(addr, abi, conn)  ← 通用合约绑定
        │   └── contract.Call(nil, &res, "symbol")      ← 调用 ERC20 symbol()
        │
        ├── [主网] GetContractSymbolOnMainNet(t.Token)
        │   ├── GetRemoteAbiFileByToken()          ← 先通过 BSC Scan API 下载 ABI
        │   │   ├── HTTP GET api.etherscan.io      ← 下载合约 ABI JSON
        │   │   ├── 写入本地文件 {tokenAddress}.abi  ← 保存到 contract/abi/ 目录
        │   │   └── UPDATE token_info SET abi_file_exist=1
        │   ├── abifile.GetAbiByToken(t.Token)     ← 使用下载的 ABI 文件
        │   └── contract.Call(nil, &res, "symbol")
        │
        ├── CheckSymbolData()                     ← Redis 去重
        │
        └── SaveSymbolData()                      ← UPDATE token_info SET symbol = ?
```

**注意**：主网上 `GetRemoteAbiFileByToken` 的 URL 写的是 `api-sepolia.etherscan.io`，与 BSC 主网不匹配，当前阶段实测应该会失败。

---

### 4.4 tokenLogoService — 代币 Logo 同步

**文件**：`services/tokenLogoService.go`

**执行频率**：每 2 小时

**功能**：从 PancakeSwap 官方代币列表 + 本地硬编码列表，同步代币的 Logo URL、Symbol、Decimals 到 MySQL。

```
UpdateTokenLogo()
    │
    ├── [第一步] 从 PancakeSwap 远程获取
    │   ├── HTTP GET config.Token.LogoUrl           ← PancakeSwap top-100 列表
    │   ├── json.Unmarshal → TokenLogoRemote 结构体
    │   └── for _, token := range tokens
    │       ├── CheckLogoData(token.Address, ...)    ← Redis 去重
    │       └── SaveLogoData(...)                    ← 更新 logo, symbol, decimals
    │
    └── [第二步] 从本地硬编码列表覆盖（优先级更高）
        └── LocalTokenLogo 常量 (见下方)
            ├── BNB (主网+测试网)
            ├── BTC/BTCB
            ├── BUSD
            ├── DAI
            ├── ETH
            ├── USDT
            ├── CAKE
            └── PLGR
```

**本地 Logo 优先级高于远程**：代码逻辑上先更新远程数据，再覆盖本地硬编码数据。

**硬编码代币列表覆盖 9 种代币**，每种包含测试网和主网两个环境下的地址、精度、Logo URL。

---

### 4.5 balanceMonitor — 合约余额监控告警

**文件**：`services/balanceMonitor.go`

**执行频率**：每 30 分钟

**功能**：检查链上资金池合约的 BNB 余额，低于阈值时发送邮件告警。

```
Monitor()
    │
    ├── [测试网] GetBalance(TestNet.NetUrl, TestNet.PledgePoolToken)
    │   └── ethclient.BalanceAt(ctx, poolAddress, nil)  ← 查询合约 BNB 余额
    │
    ├── tokenPoolBalance < threshold (10^17 wei = 0.1 BNB)  ← 阈值判断
    │   └── utils.SendEmail(邮件内容, 2)                    ← 发送告警邮件
    │       └── 收件人: config.Email.To
    │
    └── [主网] 代码被注释，未启用
```

阈值配置：`pledge_pool_token_threshold_bnb = "100000000000000000"`（0.1 BNB）

---

### 4.6 tokenLogoService — 代币 Logo 同步 (续)

**LocalTokenLogo 数据表**：

| 代币 | 测试网地址 | 主网地址 | 精度 |
|------|-----------|---------|------|
| BNB | 0x0000...0000 | 0x0000...0000 | 18 |
| BTC | 0xB5514a4F... | 0x7130d2A1... | 8 |
| BTCB | 同上 | 同上 | 8 |
| BUSD | 0xE676Dcd7... | 0xe9e7CEA3... | 18 |
| DAI | 0x490BC3FC... | 0x1AF3F329... | 18 |
| ETH | (空) | 0x2170ed08... | 18 |
| USDT | (空) | 0x55d39832... | 18 |
| CAKE | 0xEAEd0816... | 0x0e09fabb... | 18 |
| PLGR | (空) | 0x6Aa91CbF... | 18 |

---

## 五、数据模型

### 5.1 poolbases 表 — 资金池基础信息

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int (PK) | 自增主键 |
| pool_id | int | 资金池编号（从 1 开始） |
| chain_id | string | 链 ID (97=测试网, 56=主网) |
| settle_time | string | 结算时间戳 |
| end_time | string | 结束时间戳 |
| interest_rate | string | 利率（如 5000000 = 5%） |
| max_supply | string | 最大供应量 |
| lend_supply | string | 当前出借供应量 |
| borrow_supply | string | 当前借款供应量 |
| martgage_rate | string | 抵押率 |
| lend_token | string | 出借代币地址 |
| lend_token_info | text (JSON) | 出借代币详情（费率、Logo、名称、价格） |
| borrow_token | string | 借款代币地址 |
| borrow_token_info | text (JSON) | 借款代币详情 |
| state | string | 资金池状态 (0=未开始, 1=进行中, 2=已清算...) |
| sp_coin | string | SP Coin 地址 |
| jp_coin | string | JP Coin 地址 |
| lend_token_symbol | string | 出借代币符号 |
| borrow_token_symbol | string | 借款代币符号 |
| auto_liquidate_threshold | string | 自动清算阈值 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

### 5.2 pooldata 表 — 资金池动态数据

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int (PK) | 自增主键 |
| pool_id | string | 资金池编号 |
| chain_id | string | 链 ID |
| finish_amount_borrow | string | 已完成借款量 |
| finish_amount_lend | string | 已完成出借量 |
| liquidation_amoun_borrow | string | 借款清算量 |
| liquidation_amoun_lend | string | 出借清算量 |
| settle_amount_borrow | string | 借款结算量 |
| settle_amount_lend | string | 出借结算量 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

### 5.3 token_info 表 — 代币信息

| 字段 | 类型 | 说明 |
|------|------|------|
| id | int (PK) | 自增主键 |
| logo | string | 代币 Logo URL |
| token | string | 代币合约地址 |
| symbol | string | 代币符号 (如 BNB, BUSD) |
| chain_id | string | 链 ID |
| price | string | 代币价格 |
| decimals | int | 代币精度 |
| abi_file_exist | int | 主网 ABI 文件是否已下载 (0/1) |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

### 5.4 RedisTokenInfo — Redis 缓存结构体

| 字段 | 说明 |
|------|------|
| Logo | 代币 Logo URL |
| Token | 代币合约地址 |
| Symbol | 代币符号 |
| ChainId | 链 ID |
| Price | 代币价格 |

---

## 六、Redis 缓存策略

### Key 设计

| Key 格式 | 用途 | 过期时间 | 所属服务 |
|----------|------|---------|---------|
| `token_info:{chainId}:{token}` | 代币信息缓存（价格/Logo/Symbol） | 无过期 | 所有服务共用 |
| `base_info:pool_{chainId}_{poolId}` | 资金池基础信息 MD5 | 30 分钟 | poolService |
| `data_info:pool_{chainId}_{poolId}` | 资金池动态数据 MD5 | 30 分钟 | poolService |

### 启动时行为

`RedisFlushDB()` — 服务启动时**清空整个 Redis DB**，确保缓存与数据库一致。

### 去重逻辑（所有服务通用模式）

```
1. 从链上/远程获取最新数据
2. 查询 Redis 中该 key 的缓存值
   ├── 无缓存 → 直接写入 MySQL + Redis
   └── 有缓存 → 比对是否变化
       ├── 无变化 → 跳过 (return)
       └── 有变化 → 更新 MySQL + 更新 Redis
```

---

## 七、区块链交互汇总

### 7.1 连接的合约

| 合约 | 地址（测试网） | 地址（主网 V2.1） | 用途 |
|------|--------------|-----------------|------|
| PledgePool | `0x216f718A983FCCb462b338FA9c60f2A89199490c` | `0x25C3f3d3E3299d7C56700CE54303Fbe1E6a16fee` | 资金池管理 |
| BscPledgeOracle | `0xd96DBDC193617A0cD4bbf38E78a0fB4799A8E554` | `0x4Aa9EB3149089D7208C9C0403BF1b9bA25ff05BD` | 价格预言机 |
| PLGR Token | `0X6AA91CBFE045F9D154050226FCC830DDBA886CED` | `0x6aa91cbfe045f9d154050226fcc830ddba886ced` | 平台代币 |

### 7.2 合约调用方式

| 方式 | 方法 | 代码示例 |
|------|------|---------|
| 只读调用 (Call) | `PoolBaseInfo`, `PoolDataInfo`, `PoolLength`, `GetPrice`, `symbol()` | `pledgePoolToken.PoolBaseInfo(nil, big.NewInt(i))` |
| 交易签名写入 (Transact) | `SetPrice` | `oracle.SetPrice(&transactOpts, addr, price)` |

### 7.3 交易签名流程（仅 SavePlgrPrice* 使用）

```go
privateKeyEcdsa, _ := crypto.HexToECDSA(plgrAdminPrivateKey)         // 加载私钥
auth, _ := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, chainId) // 创建签名器
transactOpts := bind.TransactOpts{
    From:   auth.From,
    Signer: auth.Signer,       // 签名方法
    Value:  big.NewInt(0),     // 不发送 BNB
    Context: ctx,              // 5 秒超时
    NoSend: false,             // 发送交易
}
tx, err := oracle.SetPrice(&transactOpts, plgrAddress, price)        // 签名并发送
```

---

## 八、任务调度配置总表

| 任务 | 周期 | 启动时执行 | 涉及链操作 | 写 MySQL | 写 Redis | 发邮件 |
|------|------|-----------|-----------|---------|---------|-------|
| `UpdateAllPoolInfo` | 每 2 分钟 | ✅ | 读 PledgePool | ✅ poolbases + pooldata | ✅ MD5 缓存 | ❌ |
| `UpdateContractPrice` | 每 1 分钟 | ✅ | 读 Oracle | ✅ token_info.price | ✅ 价格缓存 | ❌ |
| `UpdateContractSymbol` | 每 2 小时 | ✅ | 读 ERC20 symbol() | ✅ token_info.symbol | ✅ Symbol 缓存 | ❌ |
| `UpdateTokenLogo` | 每 2 小时 | ✅ | HTTP 远程拉取 | ✅ token_info.logo | ✅ Logo 缓存 | ❌ |
| `Monitor` | 每 30 分钟 | ✅ | 读合约余额 | ❌ | ❌ | ✅ 余额告警 |
| `SavePlgrPriceTestNet` | 每 30 分钟 | ✅ | **写 Oracle** SetPrice | ❌ | ❌ | ❌ |

---

## 九、当前限制与注释代码

### 9.1 主网代码全部被注释

poolService.go:28 — `// s.UpdatePoolInfo(config.Config.MainNet...)`
tokenPriceService.go:48-58 — 主网价格获取逻辑被注释
balanceMonitor.go:42-54 — 主网余额监控被注释

### 9.2 硬编码值

SavePlgrPriceTestNet — 测试网 PLGR 价格写死为 `22222`

### 9.3 潜在问题

1. **`GetRemoteAbiFileByToken` URL 写的是 `api-sepolia.etherscan.io`**，应该是 `api.bscscan.com`（BSC 主网），可能存在兼容问题
2. **所有价格返回类型为 `int64`**，代币价格若超过 `int64` 范围会溢出
3. **没有重试机制** — 链上 RPC 调用失败直接跳过
4. **Redis key 无过期时间**（`RedisSet` 的 ttl=0），可能造成内存泄漏
5. **`contract/abi/pledge_pool.abi` 文件内容并非标准 ABI JSON**，而是模拟的 API 响应数据，真正用于合约绑定的 ABI 在编译合约时生成

---

## 十、部署与运维

### Systemd 服务

```ini
[Unit]
Description=pledge task service

[Service]
Type=simple
Restart=always
User=root
EnvironmentFile=/etc/systemd/pledge.env
ExecStart=/home/ubuntu/codespace/pledge-backend/schedule/pledge_task
```

### 环境变量

| 变量名 | 说明 |
|--------|------|
| `plgr_admin_private_key` | 管理员私钥（16进制，不含 0x 前缀），用于签名 SetPrice 交易 |

### 编译与运行

```bash
cd schedule
go build -o pledge_task .
./pledge_task
```
