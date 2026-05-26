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
	// PooldataExpireTime expire time
	PooldataExpireTime = 5 * time.Minute
)

var _ PooldataCache = (*pooldataCache)(nil)

// PooldataCache cache interface
type PooldataCache interface {
	Set(ctx context.Context, id uint64, data *model.Pooldata, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Pooldata, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Pooldata, error)
	MultiSet(ctx context.Context, data []*model.Pooldata, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// pooldataCache define a cache struct
type pooldataCache struct {
	cache cache.Cache
}

// NewPooldataCache new a cache
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

// GetPooldataCacheKey cache key
func (c *pooldataCache) GetPooldataCacheKey(id uint64) string {
	return pooldataCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
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

// Get cache value
func (c *pooldataCache) Get(ctx context.Context, id uint64) (*model.Pooldata, error) {
	var data *model.Pooldata
	cacheKey := c.GetPooldataCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
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

// MultiGet multiple get cache, return key in map is id value
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

// Del delete cache
func (c *pooldataCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPooldataCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *pooldataCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetPooldataCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *pooldataCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
