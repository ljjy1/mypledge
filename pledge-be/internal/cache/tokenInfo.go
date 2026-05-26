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
	tokenInfoCachePrefixKey = "tokenInfo:"
	// TokenInfoExpireTime expire time
	TokenInfoExpireTime = 5 * time.Minute
)

var _ TokenInfoCache = (*tokenInfoCache)(nil)

// TokenInfoCache cache interface
type TokenInfoCache interface {
	Set(ctx context.Context, id uint64, data *model.TokenInfo, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.TokenInfo, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.TokenInfo, error)
	MultiSet(ctx context.Context, data []*model.TokenInfo, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// tokenInfoCache define a cache struct
type tokenInfoCache struct {
	cache cache.Cache
}

// NewTokenInfoCache new a cache
func NewTokenInfoCache(cacheType *database.CacheType) TokenInfoCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.TokenInfo{}
		})
		return &tokenInfoCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.TokenInfo{}
		})
		return &tokenInfoCache{cache: c}
	}

	return nil // no cache
}

// GetTokenInfoCacheKey cache key
func (c *tokenInfoCache) GetTokenInfoCacheKey(id uint64) string {
	return tokenInfoCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *tokenInfoCache) Set(ctx context.Context, id uint64, data *model.TokenInfo, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetTokenInfoCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *tokenInfoCache) Get(ctx context.Context, id uint64) (*model.TokenInfo, error) {
	var data *model.TokenInfo
	cacheKey := c.GetTokenInfoCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *tokenInfoCache) MultiSet(ctx context.Context, data []*model.TokenInfo, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetTokenInfoCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *tokenInfoCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.TokenInfo, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetTokenInfoCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.TokenInfo)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.TokenInfo)
	for _, id := range ids {
		val, ok := itemMap[c.GetTokenInfoCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *tokenInfoCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetTokenInfoCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *tokenInfoCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetTokenInfoCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *tokenInfoCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
