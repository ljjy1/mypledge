package cache

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-dev-frame/sponge/pkg/cache"
	"github.com/go-dev-frame/sponge/pkg/encoding"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/database"
	"pledge-be/internal/model"
)

const (
	// cache prefix key, must end with a colon
	contractCachePrefixKey = "contract:"
	// ContractExpireTime 合约缓存过期时间
	ContractExpireTime = 5 * time.Minute
)

var _ ContractCache = (*contractCache)(nil)

// ContractCache 合约缓存接口
type ContractCache interface {
	// Set 设置合约缓存
	Set(ctx context.Context, id uint64, data *model.Contract, duration time.Duration) error
	// Get 获取合约缓存
	Get(ctx context.Context, id uint64) (*model.Contract, error)
	// MultiGet 批量获取合约缓存
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Contract, error)
	// MultiSet 批量设置合约缓存
	MultiSet(ctx context.Context, data []*model.Contract, duration time.Duration) error
	// Del 删除合约缓存
	Del(ctx context.Context, id uint64) error
	// SetPlaceholder 设置占位符缓存（防止缓存穿透）
	SetPlaceholder(ctx context.Context, id uint64) error
	// IsPlaceholderErr 判断是否为占位符错误
	IsPlaceholderErr(err error) bool
}

// contractCache 合约缓存结构体
type contractCache struct {
	cache cache.Cache
}

// NewContractCache 创建合约缓存实例，根据 cacheType 选择 Redis 或 Memory 缓存
func NewContractCache(cacheType *database.CacheType) ContractCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Contract{}
		})
		return &contractCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Contract{}
		})
		return &contractCache{cache: c}
	}

	return nil // no cache
}

// GetContractCacheKey 获取合约缓存键
func (c *contractCache) GetContractCacheKey(id uint64) string {
	return contractCachePrefixKey + utils.Uint64ToStr(id)
}

// Set 将合约数据写入缓存
func (c *contractCache) Set(ctx context.Context, id uint64, data *model.Contract, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetContractCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get 从缓存读取合约数据
func (c *contractCache) Get(ctx context.Context, id uint64) (*model.Contract, error) {
	var data *model.Contract
	cacheKey := c.GetContractCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet 批量写入合约数据到缓存
func (c *contractCache) MultiSet(ctx context.Context, data []*model.Contract, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetContractCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet 批量从缓存读取合约数据，返回 map 的 key 为 id 值
func (c *contractCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Contract, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetContractCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Contract)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Contract)
	for _, id := range ids {
		val, ok := itemMap[c.GetContractCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del 删除合约缓存
func (c *contractCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetContractCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder 将占位符值写入缓存，用于防止缓存穿透（空值缓存）
func (c *contractCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetContractCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr 判断错误是否为缓存占位符错误
func (c *contractCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
