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

const TypePrintMessage = "schedule:print"

// PrintMessagePayload 定时打印消息任务载荷
type PrintMessagePayload struct {
	Message string `json:"message"`
}

// HandlePrintMessage 处理定时打印消息任务
func HandlePrintMessage(ctx context.Context, p *PrintMessagePayload) error {
	logger.Info("[定时任务] " + p.Message)
	return nil
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
	sasynq.RegisterTaskHandler(srv.Mux(), TypePrintMessage, sasynq.HandleFunc(HandlePrintMessage))
	srv.Run()
	s.server = srv

	// 创建 asynq scheduler，用于定时注册任务
	scheduler := sasynq.NewScheduler(redisCfg,
		sasynq.WithSchedulerLogger(sasynq.WithLogger(logger.Get())),
	)
	_, err := scheduler.RegisterTask("@every 1m", TypePrintMessage, &PrintMessagePayload{
		Message: "Hello! This message is printed every minute by the scheduled task.",
	})
	if err != nil {
		return fmt.Errorf("register schedule task error: %w", err)
	}
	scheduler.Run()
	s.scheduler = scheduler

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

// parseRedisConfig 从项目配置解析出 sasynq.RedisConfig
// DSN 格式: [user]:<password>@host:port/db
func parseRedisConfig() sasynq.RedisConfig {
	dsn := config.Get().Redis.Dsn
	cfg := sasynq.RedisConfig{
		Mode: sasynq.RedisModeSingle,
	}

	// 提取 password 和 host:port/db
	if idx := strings.LastIndex(dsn, "@"); idx >= 0 {
		authPart := dsn[:idx]
		rest := dsn[idx+1:]
		// authPart 格式: [user]:password
		if colonIdx := strings.Index(authPart, ":"); colonIdx >= 0 {
			cfg.Password = authPart[colonIdx+1:]
		} else {
			cfg.Password = authPart
		}

		// rest 格式: host:port/db
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
		cfg.Addr = dsn
	}

	return cfg
}
