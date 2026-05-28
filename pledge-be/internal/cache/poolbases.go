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
	poolbasesCachePrefixKey = "poolbases:"
	// PoolbasesExpireTime 资金池缓存过期时间
	PoolbasesExpireTime = 5 * time.Minute
)

var _ PoolbasesCache = (*poolbasesCache)(nil)

// PoolbasesCache 资金池缓存接口
type PoolbasesCache interface {
	// Set 设置资金池缓存
	Set(ctx context.Context, id uint64, data *model.Poolbases, duration time.Duration) error
	// Get 获取资金池缓存
	Get(ctx context.Context, id uint64) (*model.Poolbases, error)
	// MultiGet 批量获取资金池缓存
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Poolbases, error)
	// MultiSet 批量设置资金池缓存
	MultiSet(ctx context.Context, data []*model.Poolbases, duration time.Duration) error
	// Del 删除资金池缓存
	Del(ctx context.Context, id uint64) error
	// SetPlaceholder 设置占位符缓存（防止缓存穿透）
	SetPlaceholder(ctx context.Context, id uint64) error
	// IsPlaceholderErr 判断是否为占位符错误
	IsPlaceholderErr(err error) bool
}

// poolbasesCache 资金池缓存结构体
type poolbasesCache struct {
	cache cache.Cache
}

// NewPoolbasesCache 创建资金池缓存实例，根据 cacheType 选择 Redis 或 Memory 缓存
func NewPoolbasesCache(cacheType *database.CacheType) PoolbasesCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Poolbases{}
		})
		return &poolbasesCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Poolbases{}
		})
		return &poolbasesCache{cache: c}
	}

	return nil // no cache
}

// GetPoolbasesCacheKey 获取资金池缓存键
func (c *poolbasesCache) GetPoolbasesCacheKey(id uint64) string {
	return poolbasesCachePrefixKey + utils.Uint64ToStr(id)
}

// Set 将资金池数据写入缓存
func (c *poolbasesCache) Set(ctx context.Context, id uint64, data *model.Poolbases, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPoolbasesCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get 从缓存读取资金池数据
func (c *poolbasesCache) Get(ctx context.Context, id uint64) (*model.Poolbases, error) {
	var data *model.Poolbases
	cacheKey := c.GetPoolbasesCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet 批量写入资金池数据到缓存
func (c *poolbasesCache) MultiSet(ctx context.Context, data []*model.Poolbases, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPoolbasesCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet 批量从缓存读取资金池数据，返回 map 的 key 为 id 值
func (c *poolbasesCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Poolbases, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPoolbasesCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Poolbases)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Poolbases)
	for _, id := range ids {
		val, ok := itemMap[c.GetPoolbasesCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del 删除资金池缓存
func (c *poolbasesCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPoolbasesCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder 将占位符值写入缓存，用于防止缓存穿透（空值缓存）
func (c *poolbasesCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetPoolbasesCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr 判断错误是否为缓存占位符错误
func (c *poolbasesCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
