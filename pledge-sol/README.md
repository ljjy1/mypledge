## 安装依赖

```shell
# safe-contracts 目前不兼容 ethers v6，强制安装
npm install --legacy-peer-deps
```

## 测试

```shell
#编译
npx hardhat compile

# 运行所有测试（solidity + mocha）
npx hardhat test

# 运行 solidity 测试
npx hardhat test solidity

# 运行 mocha 测试
npx hardhat test mocha
```

### 指定 describe / it 运行

使用 Mocha 的 `--grep` 按名称过滤测试：

```shell
# 只运行 describe 名包含 "Sepolia" 的测试
npx hardhat test mocha --grep "Sepolia"

# 只运行 it 名包含 "多签" 的测试
npx hardhat test mocha --grep "多签"

# 指定具体文件 + describe 名称
npx hardhat test mocha test/pledgePool-sepolia.ts --network sepolia --grep "多签操作"

# 正则匹配：it 名包含 "setOracle" 或 "owner"
npx hardhat test mocha --grep "/setOracle|owner/"

# 反向匹配：跳过包含 "sepolia" 的测试（忽略大小写）
npx hardhat test mocha --grep "/sepolia/i" --invert
```

## 部署到本地节点

本地部署分两步：先启动 Hardhat 节点，再部署合约到该节点。之后 Go 等客户端通过 RPC 连接进行测试。

### 1. 启动 Hardhat 本地节点

```shell
# terminal 1
npx hardhat node
```

启动后输出 20 个预 funded 账户（每个 10000 ETH），JSON-RPC 服务监听在 `http://127.0.0.1:8545`，chainId 为 `31337`。

### 2. 部署合约

```shell
# 新终端terminal 2：传入必需参数部署合约到本地节点
npx hardhat ignition deploy ignition/modules/PledgeProtocol.ts \
  --network localhost \
  --parameters '<参数JSON>'
```

参数说明：

| 参数 | 说明 | 本地测试推荐值 |
|------|------|----------------|
| `owner` | 合约 owner 地址 | `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`（账户 #0） |
| `swapRouter` | DEX 路由地址 | 本地测试可填任意地址 |
| `feeAddress` | 手续费接收地址 | 本地测试可填任意地址 |

完整例子（直接传参）：

```shell
npx hardhat ignition deploy ignition/modules/PledgeProtocol.ts \
  --network localhost \
  --parameters '{
    "PledgeProtocol": {
      "owner": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
      "swapRouter": "0x0000000000000000000000000000000000000001",
      "feeAddress": "0x0000000000000000000000000000000000000002"
    }
  }'
```

部署成功输出示例：

```
[ PledgeProtocol ] successfully deployed 🚀

Deployed Addresses

PledgeProtocol#BorrowDebtToken - 0x...
PledgeProtocol#BscPledgeOracle    - 0x...
PledgeProtocol#LendDebtToken     - 0x...
PledgeProtocol#PledgePool        - 0x...
```

记下输出的合约地址，后续 Go 测试需要用到。

### 3. 部署后操作

部署完成后通过 owner 账户执行：

1. **设置价格** — `BscPledgeOracle.setPrice(asset, price)` 为各资产配置价格
2. **设置手续费** — `PledgePool.setFee(lendFee, borrowFee)`
3. **创建借贷池** — `PledgePool.createPledgePool(params)`


### 本地节点账户

Hardhat 节点预置 20 个测试账户，账户 #0 可直接用作 owner：

| 字段 | 值 |
|------|-----|
| 地址 | `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266` |
| 私钥 | `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80` |

### 常用命令

```shell
# 清除本地部署记录（重新部署前需要执行）
npx hardhat ignition clear
```

## 部署到测试网/主网

```shell
# 部署到 Sepolia
npx hardhat ignition deploy ignition/modules/PledgeProtocol.ts \
  --network sepolia \
  --parameters ignition/parameters/sepolia.json

# 部署到 BSC 主网（需配置 bsc 网络）
npx hardhat ignition deploy ignition/modules/PledgeProtocol.ts \
  --network bsc \
  --parameters ignition/parameters/bsc.json
```

## Sepolia 测试网部署测试

`test/pledgePool-sepolia.ts` 是 Sepolia 测试网的部署 + 验证一体化脚本。

### 首次部署

先在 `.env` 中填入 Sepolia RPC URL 和部署账户私钥：

```
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_API_KEY
SEPOLIA_PRIVATE_KEY=0xYOUR_DEPLOYER_PRIVATE_KEY
```

执行部署（多签地址和测试参数已在脚本中硬编码）：

```shell
# 完整部署所有合约（耗时 2-5 分钟）
npx hardhat test mocha test/pledgePool-sepolia.ts \
  --network sepolia
```

首次执行会：
1. 部署 UniswapV2 + Mock ERC20 + Oracle + DebtToken + PledgePool
2. 配置 minter/手续费/创建测试池
3. 将 owner 转移给多签 `0xDF39470F7c82e62174A06BE933B8D94c5A48Dc11`
4. 保存合约地址到 `.sepolia-deployments.json`

### 再次运行（跳过部署，仅验证）

```shell
npx hardhat test mocha test/pledgePool-sepolia.ts \
  --network sepolia
```

检测到 `.sepolia-deployments.json` 时自动跳过部署，直接从链上读取合约状态并运行 14 个验证测试。

### 强制重新部署

```shell
REDEPLOY=true npx hardhat test mocha test/pledgePool-sepolia.ts \
  --network sepolia
```

### 多签操作验证

部署完成后，运行脚本时会输出多签 `setFeeAddress` 操作指引，按提示在 <https://app.safe.global> 中执行：

```
Safe → New Transaction → Contract Interaction
  目标合约: 0x...PledgePool 地址
  方法:     setFeeAddress(address payable _feeAddress)
  参数:     _feeAddress: 0x...手续费接收地址
```

执行后重新运行脚本，`验证多签 setFeeAddress 已生效` 测试会验证操作结果。

### 环境变量

| 变量 | 说明 |
|------|------|
| `REDEPLOY=true` | 强制重新部署 |
| `DEPLOYMENTS_FILE=xxx.json` | 自定义部署记录路径（默认 `.sepolia-deployments.json`） |

## BSC 主网参考地址

| 合约 | 地址 |
|------|------|
| PancakeSwap Router v2 | `0x10ED43C718714eb63d5aA57B78B54704E256024E` |
| WBNB | `0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c` |


## 提取ABI和BIN 后续用于后端连接 (需要合约已经编译了)
```shell
#在终端 pledge-sol目录下运行
jq '.abi' artifacts/contracts/pledge/PledgePool.sol/PledgePool.json > PledgePool.abi  
 
jq -r '.bytecode' artifacts/contracts/pledge/PledgePool.sol/PledgePool.json > PledgePool.bin
```

```shell
#脚本生成 go合约绑定文件和对应abi和bin json 需要提前安装好jq和abigen
#在终端 pledge-sol目录下运行
# 默认 --pkg bindcode
node abi-bin-go-code/extract.mjs
# 自定义包名mycustompkg
node abi-bin-go-code/extract.mjs mycustompkg
```
