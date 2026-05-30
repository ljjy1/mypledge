package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-dev-frame/sponge/pkg/app"
	"github.com/go-dev-frame/sponge/pkg/logger"

	"pledge-be/internal/database"
	"pledge-be/internal/model"
)

var _ app.IServer = (*EventListener)(nil)

const (
	maxEventsPerChain = 200
	reconnectDelay    = 5 * time.Second
)

// EventListener 合约事件监听服务，监听所有链上所有合约的事件并存入 Redis
type EventListener struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewEventListener 创建事件监听服务实例
func NewEventListener() *EventListener {
	ctx, cancel := context.WithCancel(context.Background())
	return &EventListener{
		ctx:    ctx,
		cancel: cancel,
	}
}

// Start 启动事件监听，查询所有合约并按链 ID 分组，每个链启动一个 goroutine 监听事件
func (l *EventListener) Start() error {
	db := database.GetDB()
	var contracts []model.Contract
	if err := db.WithContext(l.ctx).Find(&contracts).Error; err != nil {
		return fmt.Errorf("[listener] query contracts error: %w", err)
	}
	if len(contracts) == 0 {
		logger.Warn("[listener] no contracts found, event listener skipped")
		return nil
	}

	// 按 chainID 分组
	chainMap := make(map[string][]model.Contract)
	for _, c := range contracts {
		if c.ContractAddress == "" || c.NodeURL == "" {
			continue
		}
		chainMap[c.ChainID] = append(chainMap[c.ChainID], c)
	}

	for chainID, chainContracts := range chainMap {
		wsURL := toWSURL(chainContracts[0].NodeURL)
		if wsURL == "" {
			logger.Warn("[listener] unsupported protocol, skip chain",
				logger.String("chainID", chainID),
				logger.String("nodeURL", chainContracts[0].NodeURL))
			continue
		}

		// 收集该链下所有合约地址和名称映射
		addrMap := make(map[string]string)
		addresses := make([]common.Address, 0, len(chainContracts))
		for _, c := range chainContracts {
			addr := common.HexToAddress(c.ContractAddress)
			addresses = append(addresses, addr)
			addrMap[strings.ToLower(c.ContractAddress)] = c.ContractName
		}

		l.wg.Add(1)
		go l.watchChain(l.ctx, chainID, wsURL, addresses, addrMap)
		logger.Info("[listener] watching chain",
			logger.String("chainID", chainID),
			logger.String("wsURL", wsURL),
			logger.Int("contractCount", len(addresses)))
	}

	logger.Info("[listener] event listener started",
		logger.Int("chainCount", len(chainMap)))
	return nil
}

// Stop 优雅关闭所有事件监听
func (l *EventListener) Stop() error {
	l.cancel()
	l.wg.Wait()
	logger.Info("[listener] event listener stopped")
	return nil
}

// String 返回服务名称
func (l *EventListener) String() string {
	return "event listener"
}

// toWSURL 将 http/https RPC URL 转换为 ws/wss WebSocket URL
func toWSURL(raw string) string {
	switch {
	case strings.HasPrefix(raw, "https://"):
		return "wss://" + strings.TrimPrefix(raw, "https://")
	case strings.HasPrefix(raw, "http://"):
		return "ws://" + strings.TrimPrefix(raw, "http://")
	case strings.HasPrefix(raw, "wss://") || strings.HasPrefix(raw, "ws://"):
		return raw
	default:
		return ""
	}
}

// watchChain 监听单条链上所有合约的事件，支持断线自动重连
func (l *EventListener) watchChain(ctx context.Context, chainID, wsURL string, addresses []common.Address, addrMap map[string]string) {
	defer l.wg.Done()

	filterQ := ethereum.FilterQuery{
		Addresses: addresses,
	}

	var retries int
	const maxRetries = 12

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		client, err := ethclient.Dial(wsURL)
		if err != nil {
			retries++
			logger.Warn("[listener] dial ws error, retrying",
				logger.String("chainID", chainID),
				logger.Int("retry", retries),
				logger.Int("maxRetries", maxRetries),
				logger.Err(err))
			if retries >= maxRetries {
				logger.Error("[listener] max retries reached, stop watching chain",
					logger.String("chainID", chainID),
					logger.Int("retries", retries))
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(reconnectDelay):
			}
			continue
		}

		logs := make(chan types.Log, 100)
		sub, err := client.SubscribeFilterLogs(ctx, filterQ, logs)
		if err != nil {
			logger.Warn("[listener] subscribe filter logs error, retrying",
				logger.String("chainID", chainID),
				logger.Err(err))
			client.Close()
			select {
			case <-ctx.Done():
				return
			case <-time.After(reconnectDelay):
			}
			continue
		}

		logger.Info("[listener] connected and subscribed",
			logger.String("chainID", chainID),
			logger.String("wsURL", wsURL))

		retries = 0 // 连接成功，重置重连计数

		// 监听事件直到 context 取消或 subscription 出错
		func() {
			defer func() {
				sub.Unsubscribe()
				client.Close()
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					logger.Warn("[listener] subscription error, reconnecting",
						logger.String("chainID", chainID),
						logger.Err(err))
					return
				case vLog := <-logs:
					l.handleLog(ctx, chainID, vLog, addrMap)
				}
			}
		}()
	}
}

// EventRecord Redis 中存储的事件记录结构
type EventRecord struct {
	ChainID         string `json:"chainID"`
	ContractAddress string `json:"contractAddress"`
	ContractName    string `json:"contractName"`
	BlockNumber     uint64 `json:"blockNumber"`
	TxHash          string `json:"txHash"`
	LogIndex        uint   `json:"logIndex"`
	EventSignature  string `json:"eventSignature"`
	Data            string `json:"data"`
	Timestamp       int64  `json:"timestamp"`
}

// handleLog 处理收到的日志事件，序列化后存储到 Redis（按链 ID 区分，最多 200 条）
func (l *EventListener) handleLog(ctx context.Context, chainID string, vLog types.Log, addrMap map[string]string) {
	addrL := strings.ToLower(vLog.Address.Hex())
	contractName := addrMap[addrL]

	eventSig := ""
	if len(vLog.Topics) > 0 {
		eventSig = vLog.Topics[0].Hex()
	}

	record := EventRecord{
		ChainID:         chainID,
		ContractAddress: vLog.Address.Hex(),
		ContractName:    contractName,
		BlockNumber:     vLog.BlockNumber,
		TxHash:          vLog.TxHash.Hex(),
		LogIndex:        vLog.Index,
		EventSignature:  eventSig,
		Data:            common.Bytes2Hex(vLog.Data),
		Timestamp:       time.Now().Unix(),
	}

	data, err := json.Marshal(record)
	if err != nil {
		logger.Warn("[listener] marshal event error", logger.Err(err))
		return
	}

	redisCli := database.GetRedisCli()
	key := fmt.Sprintf(database.RedisKeyEvents, chainID)

	// LPush 插入头部，LTrim 截断保留前 N 条，实现最多 200 条的滑动窗口
	if err := redisCli.LPush(ctx, key, data).Err(); err != nil {
		logger.Warn("[listener] redis LPush error", logger.Err(err))
		return
	}
	if err := redisCli.LTrim(ctx, key, 0, maxEventsPerChain-1).Err(); err != nil {
		logger.Warn("[listener] redis LTrim error", logger.Err(err))
	}
}
