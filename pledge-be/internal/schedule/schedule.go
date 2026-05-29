package schedule

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sasynq"

	"pledge-be/internal/config"
)

// 任务类型常量
const (
	TypePoolInfo   = "schedule:pool_info"
	TypePoolSettle = "schedule:pool_settle"
)

// EmptyPayload 无需参数的通用任务载荷
type EmptyPayload struct {
	Dummy string `json:"dummy"`
}

// ScheduleServer 封装 asynq Scheduler + Server，实现 app.IServer 接口
type ScheduleServer struct {
	scheduler *sasynq.Scheduler
	server    *sasynq.Server
}

// NewScheduleServer 创建定时任务服务
func NewScheduleServer() *ScheduleServer {
	return &ScheduleServer{}
}

// Start 启动 asynq server 和 scheduler
func (s *ScheduleServer) Start() error {
	redisCfg := parseRedisConfig()

	// 创建 asynq server，用于消费任务
	serverCfg := sasynq.DefaultServerConfig(sasynq.WithLogger(logger.Get()))
	srv := sasynq.NewServer(redisCfg, serverCfg)
	srv.Use(sasynq.LoggingMiddleware(sasynq.WithLogger(logger.Get())))

	// 注册所有任务处理器
	registerHandlers(srv)

	srv.Run()
	s.server = srv

	// 创建 asynq scheduler，用于定时注册任务
	scheduler := sasynq.NewScheduler(redisCfg,
		sasynq.WithSchedulerLogger(sasynq.WithLogger(logger.Get())),
	)

	// 注册所有定时任务
	if err := registerSchedulerTasks(scheduler); err != nil {
		return fmt.Errorf("register schedule tasks error: %w", err)
	}

	scheduler.Run()
	s.scheduler = scheduler

	// 启动时立即执行一次所有任务
	go runInitialTasks()

	logger.Info("[schedule] asynq scheduler and server started")
	return nil
}

// Stop 优雅关闭
func (s *ScheduleServer) Stop() error {
	if s.scheduler != nil {
		s.scheduler.Shutdown()
	}
	if s.server != nil {
		s.server.Shutdown()
	}
	logger.Info("[schedule] asynq scheduler and server stopped")
	return nil
}

// String 返回服务名称
func (s *ScheduleServer) String() string {
	return "schedule server"
}

// registerHandlers 注册所有任务处理器
func registerHandlers(srv *sasynq.Server) {
	sasynq.RegisterTaskHandler(srv.Mux(), TypePoolInfo, sasynq.HandleFunc(
		func(ctx context.Context, _ *EmptyPayload) error {
			return PoolService(ctx)
		},
	))
	sasynq.RegisterTaskHandler(srv.Mux(), TypePoolSettle, sasynq.HandleFunc(
		func(ctx context.Context, _ *EmptyPayload) error {
			return SettleService(ctx)
		},
	))
}

// registerSchedulerTasks 注册所有定时任务
func registerSchedulerTasks(scheduler *sasynq.Scheduler) error {
	tasks := []struct {
		cron     string
		typeName string
		payload  interface{}
		desc     string
	}{
		{cron: "@every 2m", typeName: TypePoolInfo, payload: &EmptyPayload{}, desc: "资金池数据同步"},
		{cron: "@every 5m", typeName: TypePoolSettle, payload: &EmptyPayload{}, desc: "资金池结算"},
	}

	for _, t := range tasks {
		id, err := scheduler.RegisterTask(t.cron, t.typeName, t.payload)
		if err != nil {
			return fmt.Errorf("register task %s(%s) error: %w", t.desc, t.typeName, err)
		}
		logger.Info("[schedule] registered task",
			logger.String("desc", t.desc),
			logger.String("cron", t.cron),
			logger.String("entryID", id),
		)
	}
	return nil
}

// runInitialTasks 启动时立即执行所有任务一次
func runInitialTasks() {
	logger.Info("[schedule] running initial tasks...")

	// 清空 Redis 缓存
	EnsureRedisFlush()

	ctx := context.Background()
	tasks := []struct {
		name string
		fn   func(context.Context) error
	}{
		{"PoolInfo", PoolService},
	}

	for _, t := range tasks {
		logger.Info("[schedule] initial run: " + t.name)
		if err := t.fn(ctx); err != nil {
			logger.Warn("[schedule] initial run error", logger.String("task", t.name), logger.Err(err))
		}
	}

	logger.Info("[schedule] all initial tasks completed")
}

// parseRedisConfig 从项目配置解析出 sasynq.RedisConfig
func parseRedisConfig() sasynq.RedisConfig {
	dsn := config.Get().Redis.Dsn
	cfg := sasynq.RedisConfig{
		Mode: sasynq.RedisModeSingle,
	}

	// 解析 DSN 中 @ 前的认证信息（密码）
	if idx := strings.LastIndex(dsn, "@"); idx >= 0 {
		authPart := dsn[:idx]
		rest := dsn[idx+1:]
		// 从认证信息中提取密码（格式 user:password 或 password）
		if colonIdx := strings.Index(authPart, ":"); colonIdx >= 0 {
			cfg.Password = authPart[colonIdx+1:]
		} else {
			cfg.Password = authPart
		}

		// 从 @ 后的部分解析地址和数据库编号（格式 host:port/db）
		if slashIdx := strings.Index(rest, "/"); slashIdx >= 0 {
			cfg.Addr = rest[:slashIdx]
			dbStr := rest[slashIdx+1:]
			if dbStr != "" {
				if db, err := strconv.Atoi(dbStr); err == nil {
					cfg.DB = db
				}
			}
		} else {
			cfg.Addr = rest
		}
	} else {
		// 无 @ 符号，整个 DSN 视为地址
		cfg.Addr = dsn
	}

	return cfg
}
