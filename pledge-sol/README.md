## 安装依赖
```shell
#safe-contracts现在不兼容ethers6版本强制安装下
npm install --legacy-peer-deps
```

## 测试
```shell
#运行所有测试
npx hardhat test

#运行solidity测试
npx hardhat test solidity

#运行mocha测试
npx hardhat test mocha
```

```shell
npx hardhat ignition deploy ignition/modules/Counter.ts
```

