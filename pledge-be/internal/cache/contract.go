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
	// ContractExpireTime expire time
	ContractExpireTime = 5 * time.Minute
)

var _ ContractCache = (*contractCache)(nil)

// ContractCache cache interface
type ContractCache interface {
	Set(ctx context.Context, id uint64, data *model.Contract, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Contract, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Contract, error)
	MultiSet(ctx context.Context, data []*model.Contract, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// contractCache define a cache struct
type contractCache struct {
	cache cache.Cache
}

// NewContractCache new a cache
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

// GetContractCacheKey cache key
func (c *contractCache) GetContractCacheKey(id uint64) string {
	return contractCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
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

// Get cache value
func (c *contractCache) Get(ctx context.Context, id uint64) (*model.Contract, error) {
	var data *model.Contract
	cacheKey := c.GetContractCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
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

// MultiGet multiple get cache, return key in map is id value
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

// Del delete cache
func (c *contractCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetContractCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *contractCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetContractCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *contractCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
