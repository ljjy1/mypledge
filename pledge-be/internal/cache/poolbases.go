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
	// PoolbasesExpireTime expire time
	PoolbasesExpireTime = 5 * time.Minute
)

var _ PoolbasesCache = (*poolbasesCache)(nil)

// PoolbasesCache cache interface
type PoolbasesCache interface {
	Set(ctx context.Context, id uint64, data *model.Poolbases, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Poolbases, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Poolbases, error)
	MultiSet(ctx context.Context, data []*model.Poolbases, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// poolbasesCache define a cache struct
type poolbasesCache struct {
	cache cache.Cache
}

// NewPoolbasesCache new a cache
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

// GetPoolbasesCacheKey cache key
func (c *poolbasesCache) GetPoolbasesCacheKey(id uint64) string {
	return poolbasesCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
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

// Get cache value
func (c *poolbasesCache) Get(ctx context.Context, id uint64) (*model.Poolbases, error) {
	var data *model.Poolbases
	cacheKey := c.GetPoolbasesCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
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

// MultiGet multiple get cache, return key in map is id value
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

// Del delete cache
func (c *poolbasesCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPoolbasesCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *poolbasesCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetPoolbasesCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *poolbasesCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
