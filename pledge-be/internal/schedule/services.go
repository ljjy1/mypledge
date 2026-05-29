package schedule

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"

	"github.com/go-dev-frame/sponge/pkg/logger"

	"pledge-be/internal/config"
	"pledge-be/internal/contract/bindcode"
	"pledge-be/internal/database"
	"pledge-be/internal/model"
)

// getRedisCli 获取 Redis 客户端
func getRedisCli() *redis.Client {
	return database.GetRedisCli()
}

// ==================== Redis 活跃池缓存（按链分组）====================

const (
	activePoolsCacheKeyPrefix   = "active_pools:"
	activePoolsCacheTTL         = 24 * time.Hour
	pledgePoolContractsCacheKey = "pledge_pool_contracts"
	pledgePoolContractsCacheTTL = 24 * time.Hour
)

// activePoolEntry 活跃池条目，记录所属合约 ID 和池子 ID
type activePoolEntry struct {
	ContractID uint64
	PoolID     int
}

func (e activePoolEntry) encode() string {
	return fmt.Sprintf("%d:%d", e.ContractID, e.PoolID)
}

func decodeActivePoolEntry(s string) (activePoolEntry, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return activePoolEntry{}, fmt.Errorf("invalid active pool entry: %s", s)
	}
	cid, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return activePoolEntry{}, fmt.Errorf("invalid contract id in entry %q: %w", s, err)
	}
	pid, err := strconv.Atoi(parts[1])
	if err != nil {
		return activePoolEntry{}, fmt.Errorf("invalid pool id in entry %q: %w", s, err)
	}
	return activePoolEntry{ContractID: cid, PoolID: pid}, nil
}

// getActivePoolEntries 从 Redis 缓存或 DB 获取指定链上活跃的资金池条目（MATCH + EXECUTION 状态）
// 缓存 key 为 active_pools:<chainID>，存储格式为 contractID:poolID 的列表，TTL 24 小时
func getActivePoolEntries(ctx context.Context, chainID string) ([]activePoolEntry, error) {
	redisCli := getRedisCli()
	key := fmt.Sprintf("%s%s", activePoolsCacheKeyPrefix, chainID)

	// 1. 尝试从 Redis 列表读取缓存
	vals, err := redisCli.LRange(ctx, key, 0, -1).Result()
	if err == nil && len(vals) > 0 {
		entries := make([]activePoolEntry, 0, len(vals))
		for _, v := range vals {
			entry, err := decodeActivePoolEntry(v)
			if err != nil {
				logger.Warn("[cache] skip invalid active pool entry", logger.String("value", v), logger.Err(err))
				continue
			}
			entries = append(entries, entry)
		}
		if len(entries) > 0 {
			return entries, nil
		}
	}

	// 2. 缓存未命中，从 DB 查询该链上所有活跃池
	db := database.GetDB()
	var activePools []model.Poolbases
	if err := db.WithContext(ctx).
		Where("chain_id = ? AND state IN ?", chainID, []model.PoolState{model.PoolStateMatch, model.PoolStateExecution}).
		Find(&activePools).Error; err != nil {
		return nil, fmt.Errorf("query active pools error: %w", err)
	}

	entries := make([]activePoolEntry, len(activePools))
	for i, p := range activePools {
		entries[i] = activePoolEntry{ContractID: p.ContractID, PoolID: p.PoolID}
	}

	if len(entries) > 0 {
		// 3. 写入 Redis 列表缓存，TTL 24 小时
		for _, e := range entries {
			redisCli.RPush(ctx, key, e.encode())
		}
		redisCli.Expire(ctx, key, activePoolsCacheTTL)
	}

	return entries, nil
}

// removePoolEntryFromCache 从 Redis 缓存中移除指定的资金池条目
func removePoolEntryFromCache(ctx context.Context, chainID string, entry activePoolEntry) {
	redisCli := getRedisCli()
	key := fmt.Sprintf("%s%s", activePoolsCacheKeyPrefix, chainID)
	redisCli.LRem(ctx, key, 0, entry.encode())
}

// pushPoolEntryToCache 向 Redis 缓存中添加一条资金池条目（用于新建池子后主动入缓存）
func pushPoolEntryToCache(ctx context.Context, chainID string, entry activePoolEntry) {
	redisCli := getRedisCli()
	key := fmt.Sprintf("%s%s", activePoolsCacheKeyPrefix, chainID)
	redisCli.RPush(ctx, key, entry.encode())
	redisCli.Expire(ctx, key, activePoolsCacheTTL)
}

// ==================== 合约缓存 ====================

// getPledgePoolContracts 从 Redis 缓存或 DB 获取所有 PledgePool 合约列表，缓存 TTL 24 小时
func getPledgePoolContracts(ctx context.Context) ([]model.Contract, error) {
	redisCli := getRedisCli()

	// 1. 尝试从 Redis 读取缓存
	data, err := redisCli.Get(ctx, pledgePoolContractsCacheKey).Bytes()
	if err == nil {
		var contracts []model.Contract
		if err := json.Unmarshal(data, &contracts); err == nil && len(contracts) > 0 {
			return contracts, nil
		}
	}

	// 2. 缓存未命中，从 DB 查询
	db := database.GetDB()
	var contracts []model.Contract
	if err := db.WithContext(ctx).
		Where("contract_name = ?", "PledgePool").
		Find(&contracts).Error; err != nil {
		return nil, fmt.Errorf("query pledge pool contracts error: %w", err)
	}

	// 3. 写入缓存
	if len(contracts) > 0 {
		jsonBytes, _ := json.Marshal(contracts)
		redisCli.Set(ctx, pledgePoolContractsCacheKey, jsonBytes, pledgePoolContractsCacheTTL)
	}

	return contracts, nil
}

// ==================== 1. poolService — 资金池数据同步 ====================

// PoolService 资金池数据同步，每 2 分钟执行一次
// 从 contract 表读取所有 PledgePool 合约，按 chainID 分组，
// 同一链共享一个 RPC 连接，从 Redis 缓存或 DB 获取活跃池列表并同步链上数据
func PoolService(ctx context.Context) error {
	logger.Info("[schedule] PoolService start: syncing pool info from chain")

	// 1. 从缓存或 DB 获取所有 PledgePool 合约
	contracts, err := getPledgePoolContracts(ctx)
	if err != nil {
		return err
	}

	if len(contracts) == 0 {
		logger.Warn("[schedule] PoolService: no PledgePool contract found in database")
		return nil
	}

	// 2. 按 chainID 分组，建立 contractID → contract 映射
	chainContracts := make(map[string][]*model.Contract)
	contractMap := make(map[uint64]*model.Contract)
	for i := range contracts {
		c := &contracts[i]
		if c.ContractAddress == "" || c.NodeURL == "" {
			logger.Warn("[schedule] PoolService: contract missing address or node url",
				logger.Uint64("contractID", c.ID))
			continue
		}
		chainContracts[c.ChainID] = append(chainContracts[c.ChainID], c)
		contractMap[c.ID] = c
	}

	if len(chainContracts) == 0 {
		logger.Warn("[schedule] PoolService: no valid contracts to sync")
		return nil
	}

	// 3. 按链逐条处理，同一链共享一个 RPC 连接
	for chainID, contractsOnChain := range chainContracts {
		rpcURL := contractsOnChain[0].NodeURL
		logger.Info("[schedule] PoolService: syncing chain",
			logger.String("chainID", chainID),
			logger.Int("contractCount", len(contractsOnChain)))

		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			logger.Warn("[schedule] PoolService: dial rpc error",
				logger.String("chainID", chainID),
				logger.String("rpcURL", rpcURL),
				logger.Err(err))
			continue
		}

		syncChainPools(ctx, client, chainID, contractsOnChain, contractMap)
		client.Close()
	}

	logger.Info("[schedule] PoolService completed")
	return nil
}

// syncChainPools 同步指定链上所有合约的活跃资金池数据，共享一个 RPC 连接
func syncChainPools(ctx context.Context, client *ethclient.Client, chainID string, contracts []*model.Contract, contractMap map[uint64]*model.Contract) {
	// 1. 从缓存或 DB 获取该链上的活跃池列表
	entries, err := getActivePoolEntries(ctx, chainID)
	if err != nil {
		logger.Warn("[schedule] syncChainPools: get active pool entries error",
			logger.String("chainID", chainID), logger.Err(err))
		return
	}
	if len(entries) == 0 {
		logger.Info("[schedule] syncChainPools: no active pools for chain",
			logger.String("chainID", chainID))
		return
	}

	// 2. 为该链上每个合约创建调用绑定（共享 client）
	type poolBinding struct {
		contract *model.Contract
		instance *bindcode.PledgePool
	}
	bindings := make(map[uint64]*poolBinding)
	for _, c := range contracts {
		instance, err := bindcode.NewPledgePool(common.HexToAddress(c.ContractAddress), client)
		if err != nil {
			logger.Warn("[schedule] syncChainPools: new pledge pool error",
				logger.Uint64("contractID", c.ID),
				logger.String("address", c.ContractAddress),
				logger.Err(err))
			continue
		}
		bindings[c.ID] = &poolBinding{contract: c, instance: instance}
	}

	// 3. 遍历活跃池逐个同步
	for _, e := range entries {
		binding, ok := bindings[e.ContractID]
		if !ok {
			logger.Warn("[schedule] syncChainPools: no binding for contract",
				logger.Uint64("contractID", e.ContractID))
			continue
		}

		pid := big.NewInt(int64(e.PoolID))

		// 从链上读取池子基础信息
		info, err := binding.instance.PledgePoolInfoList(nil, pid)
		if err != nil {
			logger.Warn("[schedule] syncChainPools: poolInfo error",
				logger.Int("poolID", e.PoolID),
				logger.Err(err))
			continue
		}

		chainState := poolStateFromChain(info.State)

		// 同步 poolbases
		poolBase := &model.Poolbases{
			ContractID:             e.ContractID,
			PoolID:                 e.PoolID,
			SettleTime:             info.SettleTime.String(),
			EndTime:                info.EndTime.String(),
			InterestRate:           info.InterestRate.String(),
			MaxSupply:              info.MaxSupply.String(),
			LendSupply:             info.LendSupply.String(),
			BorrowSupply:           info.BorrowSupply.String(),
			MortgageRate:           info.MortgageRate.String(),
			LendToken:              info.LendToken.Hex(),
			BorrowToken:            info.BorrowToken.Hex(),
			State:                  chainState,
			LendDebtToken:          info.LendDebtToken.Hex(),
			BorrowDebtToken:        info.BorrowDebtToken.Hex(),
			AutoLiquidateThreshold: info.AutoLiquidateThreshold.String(),
			ChainID:                chainID,
		}

		md5Key := fmt.Sprintf("base_info:pool_%s_%d", chainID, e.PoolID)
		if err := saveWithDedup(ctx, md5Key, poolBase, func() error {
			return upsertPoolBase(ctx, poolBase, e.ContractID, e.PoolID)
		}); err != nil {
			logger.Warn("[schedule] syncChainPools: save pool base error",
				logger.Int("poolID", e.PoolID), logger.Err(err))
		}

		// 如果链上状态已不再是 MATCH 或 EXECUTION，从活跃缓存中移除
		if chainState != model.PoolStateMatch && chainState != model.PoolStateExecution {
			removePoolEntryFromCache(ctx, chainID, e)
			logger.Info("[schedule] syncChainPools: pool state changed, remove from active cache",
				logger.Int("poolID", e.PoolID),
				logger.String("state", string(chainState)))
		}

		// 从链上读取池子动态数据
		data, err := binding.instance.PoolDataInfoList(nil, pid)
		if err != nil {
			logger.Warn("[schedule] syncChainPools: poolData error",
				logger.Int("poolID", e.PoolID), logger.Err(err))
			continue
		}

		poolBase.SettleAmountLend = data.SettleAmountLend.String()
		poolBase.SettleAmountBorrow = data.SettleAmountBorrow.String()
		poolBase.FinishAmountLend = data.FinishAmountLend.String()
		poolBase.FinishAmountBorrow = data.FinishAmountBorrow.String()
		poolBase.LiquidationAmounLend = data.LiquidationAmountLend.String()
		poolBase.LiquidationAmounBorrow = data.LiquidationAmountBorrow.String()

		dataMd5Key := fmt.Sprintf("data_info:pool_%s_%d", chainID, e.PoolID)
		if err := saveWithDedup(ctx, dataMd5Key, poolBase, func() error {
			return upsertPoolBase(ctx, poolBase, e.ContractID, e.PoolID)
		}); err != nil {
			logger.Warn("[schedule] syncChainPools: save pool data error",
				logger.Int("poolID", e.PoolID), logger.Err(err))
		}
	}
}

// upsertPoolBase 创建或更新 poolbases 记录
func upsertPoolBase(ctx context.Context, base *model.Poolbases, contractID uint64, poolID int) error {
	db := database.GetDB()

	var existing model.Poolbases
	err := db.WithContext(ctx).
		Where("contract_id = ? AND pool_id = ?", contractID, poolID).
		First(&existing).Error

	if err != nil {
		return db.WithContext(ctx).Create(base).Error
	}

	base.ID = existing.ID
	return db.WithContext(ctx).Model(base).Updates(map[string]interface{}{
		"settle_time":              base.SettleTime,
		"end_time":                 base.EndTime,
		"interest_rate":            base.InterestRate,
		"max_supply":               base.MaxSupply,
		"lend_supply":              base.LendSupply,
		"borrow_supply":            base.BorrowSupply,
		"mortgage_rate":            base.MortgageRate,
		"lend_token":               base.LendToken,
		"borrow_token":             base.BorrowToken,
		"state":                    base.State,
		"lend_debt_token":          base.LendDebtToken,
		"borrow_debt_token":        base.BorrowDebtToken,
		"auto_liquidate_threshold": base.AutoLiquidateThreshold,
		"settle_amount_lend":       base.SettleAmountLend,
		"settle_amount_borrow":     base.SettleAmountBorrow,
		"finish_amount_lend":       base.FinishAmountLend,
		"finish_amount_borrow":     base.FinishAmountBorrow,
		"liquidation_amoun_lend":   base.LiquidationAmounLend,
		"liquidation_amoun_borrow": base.LiquidationAmounBorrow,
	}).Error
}

// poolStateFromChain 将链上 uint8 状态转换为 PoolState 枚举
func poolStateFromChain(state uint8) model.PoolState {
	switch state {
	case 0:
		return model.PoolStateMatch
	case 1:
		return model.PoolStateExecution
	case 2:
		return model.PoolStateFinish
	case 3:
		return model.PoolStateLiquidation
	case 4:
		return model.PoolStateUndone
	default:
		return model.PoolStateUndone
	}
}

// ==================== 2. settleService — 结算检测与执行 ====================

// SettleService 结算检测定时任务，每 5 分钟执行一次
// 查询 MATCH 状态且 SettleTime >= 当前时间的池子，调用链上 checkCanSettle 检测，
// 可结算则执行 SettlePool 交易
func SettleService(ctx context.Context) error {
	logger.Info("[schedule] SettleService start: checking settleable pools")

	// 1. 从缓存或 DB 获取所有 PledgePool 合约
	contracts, err := getPledgePoolContracts(ctx)
	if err != nil {
		return fmt.Errorf("get pledge pool contracts error: %w", err)
	}
	if len(contracts) == 0 {
		return nil
	}

	contractMap := make(map[uint64]*model.Contract)
	for i := range contracts {
		c := &contracts[i]
		if c.ContractAddress == "" || c.NodeURL == "" {
			continue
		}
		contractMap[c.ID] = c
	}

	// 2. 查询 MATCH 状态且 SettleTime >= 当前时间的池子
	db := database.GetDB()
	nowStr := strconv.FormatInt(time.Now().Unix(), 10)
	var settlePools []model.Poolbases
	if err := db.WithContext(ctx).
		Where("state = ? AND settle_time >= ?", model.PoolStateMatch, nowStr).
		Find(&settlePools).Error; err != nil {
		return fmt.Errorf("query settleable pools error: %w", err)
	}
	if len(settlePools) == 0 {
		logger.Info("[schedule] SettleService: no settleable pools found")
		return nil
	}

	// 3. 按 chainID 分组
	chainPools := make(map[string][]model.Poolbases)
	for _, p := range settlePools {
		chainPools[p.ChainID] = append(chainPools[p.ChainID], p)
	}

	// 4. 获取 operator 私钥
	privKeyHex := config.Get().Settle.OperatorPrivateKey
	if privKeyHex == "" {
		return fmt.Errorf("settle operator private key not configured")
	}

	// 5. 按链处理
	for chainID, pools := range chainPools {
		// 查找该链上任意一个合约获取 RPC URL
		var rpcURL string
		for _, p := range pools {
			if c, ok := contractMap[p.ContractID]; ok && c.NodeURL != "" {
				rpcURL = c.NodeURL
				break
			}
		}
		if rpcURL == "" {
			logger.Warn("[schedule] SettleService: no rpc url for chain", logger.String("chainID", chainID))
			continue
		}

		// 连接 RPC（读+写共享一个连接，每笔交易独立获取 nonce）
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			logger.Warn("[schedule] SettleService: dial rpc error",
				logger.String("chainID", chainID), logger.Err(err))
			continue
		}

		// 为该链上每个合约创建只读 caller 用于 checkCanSettle
		type poolCaller struct {
			contract *model.Contract
			instance *bindcode.PledgePoolCaller
		}
		callers := make(map[uint64]*poolCaller)
		for _, p := range pools {
			if _, ok := callers[p.ContractID]; ok {
				continue
			}
			c, ok := contractMap[p.ContractID]
			if !ok {
				continue
			}
			instance, err := bindcode.NewPledgePoolCaller(common.HexToAddress(c.ContractAddress), client)
			if err != nil {
				logger.Warn("[schedule] SettleService: new pledge pool caller error",
					logger.Uint64("contractID", p.ContractID), logger.Err(err))
				continue
			}
			callers[p.ContractID] = &poolCaller{contract: c, instance: instance}
		}

		// 遍历检测并执行结算
		for _, p := range pools {
			caller, ok := callers[p.ContractID]
			if !ok {
				continue
			}

			pid := big.NewInt(int64(p.PoolID))
			canSettle, err := caller.instance.CheckCanSettle(nil, pid)
			if err != nil {
				logger.Warn("[schedule] SettleService: checkCanSettle error",
					logger.Int("poolID", p.PoolID),
					logger.Uint64("contractID", p.ContractID),
					logger.Err(err))
				continue
			}
			if !canSettle {
				logger.Info("[schedule] SettleService: pool not ready for settle",
					logger.Int("poolID", p.PoolID))
				continue
			}

			logger.Info("[schedule] SettleService: pool is settleable, executing settle",
				logger.Int("poolID", p.PoolID))

			// 执行结算交易（独立连接，确保 nonce 正确）
			if err := settlePoolTx(ctx, rpcURL, privKeyHex, caller.contract.ContractAddress, p.PoolID); err != nil {
				logger.Warn("[schedule] SettleService: settle pool tx error",
					logger.Int("poolID", p.PoolID), logger.Err(err))
			}
		}
		client.Close()
	}

	logger.Info("[schedule] SettleService completed")
	return nil
}

// settlePoolTx 执行 SettlePool 链上交易
func settlePoolTx(ctx context.Context, rpcURL, privKeyHex, contractAddr string, poolID int) error {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("dial rpc error: %w", err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		return fmt.Errorf("parse private key error: %w", err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return fmt.Errorf("get nonce error: %w", err)
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("suggest gas price error: %w", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("get chain id error: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("create transactor error: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(8000000)

	pool, err := bindcode.NewPledgePool(common.HexToAddress(contractAddr), client)
	if err != nil {
		return fmt.Errorf("new pledge pool error: %w", err)
	}

	tx, err := pool.SettlePool(auth, big.NewInt(int64(poolID)))
	if err != nil {
		return fmt.Errorf("settle pool tx error: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("wait mined error: %w", err)
	}
	if receipt.Status == 0 {
		return fmt.Errorf("settle transaction reverted for pool %d", poolID)
	}

	logger.Info("[schedule] settlePoolTx confirmed",
		logger.Int("poolID", poolID),
		logger.String("txHash", tx.Hash().Hex()))
	return nil
}

// ==================== 通用工具 ====================

// saveWithDedup MD5 去重通用函数
// 计算 value 的 MD5，与 Redis 缓存比较，仅在有变化时执行 saveFn
func saveWithDedup(ctx context.Context, key string, value interface{}, saveFn func() error) error {
	redisCli := getRedisCli()

	// 计算当前值的 MD5
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	currentMd5 := fmt.Sprintf("%x", jsonBytes) // 简化版，直接使用 json 字符串作为指纹

	// 查询 Redis 中缓存的指纹
	cachedMd5, err := redisCli.Get(ctx, key).Result()
	if err == nil && cachedMd5 == currentMd5 {
		// 未发生变化，跳过
		return nil
	}

	// 执行保存逻辑
	if err := saveFn(); err != nil {
		return err
	}

	// 更新 Redis 缓存，30 分钟过期
	redisCli.Set(ctx, key, currentMd5, 30*time.Minute)
	return nil
}

// EnsureRedisFlush 初始化时清空 Redis 缓存（由 ScheduleServer 调用）
func EnsureRedisFlush() {
	redisCli := getRedisCli()
	if err := redisCli.FlushDB(context.Background()).Err(); err != nil {
		logger.Warn("[schedule] RedisFlushDB error", logger.Err(err))
	} else {
		logger.Info("[schedule] Redis cache flushed")
	}
}
