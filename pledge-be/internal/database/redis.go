package database

import (
	"sync"
	"time"

	"github.com/go-dev-frame/sponge/pkg/goredis"
	"github.com/go-dev-frame/sponge/pkg/tracer"

	"pledge-be/internal/config"
)

var (
	// ErrCacheNotFound No hit cache
	ErrCacheNotFound = goredis.ErrRedisNotFound
)

var (
	redisCli     *goredis.Client
	redisCliOnce sync.Once

	cacheType     *CacheType
	cacheTypeOnce sync.Once
)

// CacheType 缓存类型配置，包含缓存类型（memory 或 redis）和 Redis 客户端
type CacheType struct {
	// CType 缓存类型，可选 "memory" 或 "redis"
	CType string
	// Rdb Redis 客户端，当 CType 为 "redis" 时必填
	Rdb *goredis.Client
}

// InitCache 根据传入的缓存类型初始化缓存（memory 或 redis），若为 redis 则同时初始化 Redis 客户端
func InitCache(cType string) {
	cacheType = &CacheType{
		CType: cType,
	}

	if cType == "redis" {
		cacheType.Rdb = GetRedisCli()
	}
}

// GetCacheType 获取缓存类型实例（单例，懒加载）
func GetCacheType() *CacheType {
	if cacheType == nil {
		cacheTypeOnce.Do(func() {
			InitCache(config.Get().App.CacheType)
		})
	}

	return cacheType
}

// InitRedis 初始化 Redis 客户端连接，读取配置并设置连接超时参数，可选启用链路追踪
func InitRedis() {
	redisCfg := config.Get().Redis
	opts := []goredis.Option{
		goredis.WithDialTimeout(time.Duration(redisCfg.DialTimeout) * time.Second),
		goredis.WithReadTimeout(time.Duration(redisCfg.ReadTimeout) * time.Second),
		goredis.WithWriteTimeout(time.Duration(redisCfg.WriteTimeout) * time.Second),
	}
	if config.Get().App.EnableTrace {
		opts = append(opts, goredis.WithTracing(tracer.GetProvider()))
	}

	var err error
	redisCli, err = goredis.Init(redisCfg.Dsn, opts...)
	if err != nil {
		panic("goredis.Init error: " + err.Error())
	}
}

// GetRedisCli 获取 Redis 客户端实例（单例，懒加载）
func GetRedisCli() *goredis.Client {
	if redisCli == nil {
		redisCliOnce.Do(func() {
			InitRedis()
		})
	}

	return redisCli
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() error {
	return goredis.Close(redisCli)
}
