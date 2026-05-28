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
	// TokenInfoExpireTime 代币信息缓存过期时间
	TokenInfoExpireTime = 5 * time.Minute
)

var _ TokenInfoCache = (*tokenInfoCache)(nil)

// TokenInfoCache 代币信息缓存接口
type TokenInfoCache interface {
	// Set 设置代币信息缓存
	Set(ctx context.Context, id uint64, data *model.TokenInfo, duration time.Duration) error
	// Get 获取代币信息缓存
	Get(ctx context.Context, id uint64) (*model.TokenInfo, error)
	// MultiGet 批量获取代币信息缓存
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.TokenInfo, error)
	// MultiSet 批量设置代币信息缓存
	MultiSet(ctx context.Context, data []*model.TokenInfo, duration time.Duration) error
	// Del 删除代币信息缓存
	Del(ctx context.Context, id uint64) error
	// SetPlaceholder 设置占位符缓存（防止缓存穿透）
	SetPlaceholder(ctx context.Context, id uint64) error
	// IsPlaceholderErr 判断是否为占位符错误
	IsPlaceholderErr(err error) bool
}

// tokenInfoCache 代币信息缓存结构体
type tokenInfoCache struct {
	cache cache.Cache
}

// NewTokenInfoCache 创建代币信息缓存实例，根据 cacheType 选择 Redis 或 Memory 缓存
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

// GetTokenInfoCacheKey 获取代币信息缓存键
func (c *tokenInfoCache) GetTokenInfoCacheKey(id uint64) string {
	return tokenInfoCachePrefixKey + utils.Uint64ToStr(id)
}

// Set 将代币信息写入缓存
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

// Get 从缓存读取代币信息
func (c *tokenInfoCache) Get(ctx context.Context, id uint64) (*model.TokenInfo, error) {
	var data *model.TokenInfo
	cacheKey := c.GetTokenInfoCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet 批量写入代币信息到缓存
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

// MultiGet 批量从缓存读取代币信息，返回 map 的 key 为 id 值
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

// Del 删除代币信息缓存
func (c *tokenInfoCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetTokenInfoCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder 将占位符值写入缓存，用于防止缓存穿透（空值缓存）
func (c *tokenInfoCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetTokenInfoCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr 判断错误是否为缓存占位符错误
func (c *tokenInfoCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
