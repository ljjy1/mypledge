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
	userCachePrefixKey = "user:"
	// UserExpireTime 用户缓存过期时间
	UserExpireTime = 5 * time.Minute
)

var _ UserCache = (*userCache)(nil)

// UserCache 用户缓存接口
type UserCache interface {
	// Set 设置用户缓存
	Set(ctx context.Context, id uint64, data *model.User, duration time.Duration) error
	// Get 获取用户缓存
	Get(ctx context.Context, id uint64) (*model.User, error)
	// MultiGet 批量获取用户缓存
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.User, error)
	// MultiSet 批量设置用户缓存
	MultiSet(ctx context.Context, data []*model.User, duration time.Duration) error
	// Del 删除用户缓存
	Del(ctx context.Context, id uint64) error
	// SetPlaceholder 设置占位符缓存（防止缓存穿透）
	SetPlaceholder(ctx context.Context, id uint64) error
	// IsPlaceholderErr 判断是否为占位符错误
	IsPlaceholderErr(err error) bool
}

// userCache 用户缓存结构体
type userCache struct {
	cache cache.Cache
}

// NewUserCache 创建用户缓存实例，根据 cacheType 选择 Redis 或 Memory 缓存
func NewUserCache(cacheType *database.CacheType) UserCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.User{}
		})
		return &userCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.User{}
		})
		return &userCache{cache: c}
	}

	return nil // no cache
}

// GetUserCacheKey 获取用户缓存键
func (c *userCache) GetUserCacheKey(id uint64) string {
	return userCachePrefixKey + utils.Uint64ToStr(id)
}

// Set 将用户数据写入缓存
func (c *userCache) Set(ctx context.Context, id uint64, data *model.User, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get 从缓存读取用户数据
func (c *userCache) Get(ctx context.Context, id uint64) (*model.User, error) {
	var data *model.User
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet 批量写入用户数据到缓存
func (c *userCache) MultiSet(ctx context.Context, data []*model.User, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet 批量从缓存读取用户数据，返回 map 的 key 为 id 值
func (c *userCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.User, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.User)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.User)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del 删除用户缓存
func (c *userCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder 将占位符值写入缓存，用于防止缓存穿透（空值缓存）
func (c *userCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr 判断错误是否为缓存占位符错误
func (c *userCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
