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
	pooldataCachePrefixKey = "pooldata:"
	// PooldataExpireTime 资金池数据缓存过期时间
	PooldataExpireTime = 5 * time.Minute
)

var _ PooldataCache = (*pooldataCache)(nil)

// PooldataCache 资金池数据缓存接口
type PooldataCache interface {
	// Set 设置资金池数据缓存
	Set(ctx context.Context, id uint64, data *model.Pooldata, duration time.Duration) error
	// Get 获取资金池数据缓存
	Get(ctx context.Context, id uint64) (*model.Pooldata, error)
	// MultiGet 批量获取资金池数据缓存
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Pooldata, error)
	// MultiSet 批量设置资金池数据缓存
	MultiSet(ctx context.Context, data []*model.Pooldata, duration time.Duration) error
	// Del 删除资金池数据缓存
	Del(ctx context.Context, id uint64) error
	// SetPlaceholder 设置占位符缓存（防止缓存穿透）
	SetPlaceholder(ctx context.Context, id uint64) error
	// IsPlaceholderErr 判断是否为占位符错误
	IsPlaceholderErr(err error) bool
}

// pooldataCache 资金池数据缓存结构体
type pooldataCache struct {
	cache cache.Cache
}

// NewPooldataCache 创建资金池数据缓存实例，根据 cacheType 选择 Redis 或 Memory 缓存
func NewPooldataCache(cacheType *database.CacheType) PooldataCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Pooldata{}
		})
		return &pooldataCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Pooldata{}
		})
		return &pooldataCache{cache: c}
	}

	return nil // no cache
}

// GetPooldataCacheKey 获取资金池数据缓存键
func (c *pooldataCache) GetPooldataCacheKey(id uint64) string {
	return pooldataCachePrefixKey + utils.Uint64ToStr(id)
}

// Set 将资金池数据写入缓存
func (c *pooldataCache) Set(ctx context.Context, id uint64, data *model.Pooldata, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPooldataCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get 从缓存读取资金池数据
func (c *pooldataCache) Get(ctx context.Context, id uint64) (*model.Pooldata, error) {
	var data *model.Pooldata
	cacheKey := c.GetPooldataCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet 批量写入资金池数据到缓存
func (c *pooldataCache) MultiSet(ctx context.Context, data []*model.Pooldata, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPooldataCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet 批量从缓存读取资金池数据，返回 map 的 key 为 id 值
func (c *pooldataCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Pooldata, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPooldataCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Pooldata)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Pooldata)
	for _, id := range ids {
		val, ok := itemMap[c.GetPooldataCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del 删除资金池数据缓存
func (c *pooldataCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPooldataCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder 将占位符值写入缓存，用于防止缓存穿透（空值缓存）
func (c *pooldataCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetPooldataCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr 判断错误是否为缓存占位符错误
func (c *pooldataCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
