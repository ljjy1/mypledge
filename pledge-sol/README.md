## 安装依赖

```shell
# safe-contracts 目前不兼容 ethers v6，强制安装
npm install --legacy-peer-deps
```

## 测试

```shell
# 运行所有测试
npx hardhat test

# 运行 solidity 测试
npx hardhat test solidity

# 运行 mocha 测试
npx hardhat test mocha
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

## BSC 主网参考地址

| 合约 | 地址 |
|------|------|
| PancakeSwap Router v2 | `0x10ED43C718714eb63d5aA57B78B54704E256024E` |
| WBNB | `0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c` |
