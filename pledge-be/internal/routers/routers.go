// Package routers 是专门用于注册路由的包，支持手动注册和自动注册路由。
package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/go-dev-frame/sponge/pkg/errcode"
	"github.com/go-dev-frame/sponge/pkg/gin/handlerfunc"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware/metrics"
	"github.com/go-dev-frame/sponge/pkg/gin/prof"
	"github.com/go-dev-frame/sponge/pkg/logger"

	"pledge-be/docs"
	"pledge-be/internal/auth"
	"pledge-be/internal/config"
	"pledge-be/internal/handler"
)

var (
	// apiV1RouterFns 存储 /api/v1 分组下的路由注册函数，通过各 router 文件的 init() 自动收集
	apiV1RouterFns []func(r *gin.RouterGroup)
)

// NewRouter 创建并配置 Gin 路由器引擎，包含全局中间件和路由注册。
func NewRouter() *gin.Engine {
	// 创建一个新的 Gin 引擎实例
	r := gin.New()

	// 添加全局中间件：异常恢复，防止单个 panic 导致整个服务崩溃
	r.Use(gin.Recovery())
	// 添加全局中间件：跨域资源共享 (CORS)，允许浏览器跨域请求
	r.Use(middleware.Cors())

	// 登录和注册路由：不需要 JWT 认证，注册在所有中间件之前
	publicHandler := handler.NewUserHandler()
	r.POST("/api/v1/login", publicHandler.Login)
	r.POST("/api/v1/register", publicHandler.Register)

	// 如果配置了 HTTP 超时时间，添加请求超时中间件
	if config.Get().HTTP.Timeout > 0 {
		r.Use(middleware.Timeout(time.Second * time.Duration(config.Get().HTTP.Timeout)))
	}

	// 添加全局中间件：为每个请求生成唯一的 request_id，便于日志追踪和链路排查
	r.Use(middleware.RequestID())

	// 添加全局中间件：请求日志记录，打印请求方法和路径等信息
	r.Use(middleware.Logging(
		middleware.WithLog(logger.Get()),
		middleware.WithRequestIDFromContext(),
		middleware.WithIgnoreRoutes("/metrics"), // 忽略 metrics 路径，避免日志过于冗长
	))

	// 如果开启了指标采集，添加 Prometheus 监控指标中间件
	if config.Get().App.EnableMetrics {
		r.Use(metrics.Metrics(r,
			metrics.WithIgnoreStatusCodes(http.StatusNotFound), // 忽略 404 状态码的指标采集
		))
	}

	// 如果开启了限流，添加自适应限流中间件，防止突发流量打垮服务
	if config.Get().App.EnableLimit {
		r.Use(middleware.RateLimit(
		//middleware.WithWindow(time.Second*5), // 时间窗口，默认 10s
		//middleware.WithBucket(1000),           // 桶容量，默认 100
		//middleware.WithCPUThreshold(750),      // CPU 阈值，默认 800
		))
	}

	// 如果开启了熔断，添加自适应熔断中间件，防止级联故障
	if config.Get().App.EnableCircuitBreaker {
		r.Use(middleware.CircuitBreaker(
			middleware.WithValidCode( // 配置触发熔断的错误码
				errcode.InternalServerError.Code(),
				errcode.ServiceUnavailable.Code(),
			),
		))
	}

	// 如果开启了链路追踪，添加 OpenTelemetry 链路追踪中间件
	if config.Get().App.EnableTrace {
		r.Use(middleware.Tracing(config.Get().App.Name))
	}

	// 如果开启了性能分析，注册 pprof 性能分析路由
	if config.Get().App.EnableHTTPProfile {
		prof.Register(r, prof.WithIOWaitTime())
	}

	// 注册健康检查、心跳、错误码列表等公开路由
	r.GET("/health", handlerfunc.CheckHealth)
	r.GET("/ping", handlerfunc.Ping)
	r.GET("/codes", handlerfunc.ListCodes)

	// 非生产环境注册配置展示和 Swagger API 文档路由
	if config.Get().App.Env != "prod" {
		r.GET("/config", gin.WrapF(errcode.ShowConfig([]byte(config.Show()))))
		// 注册 Swagger 路由，代码由 swag init 生成
		docs.SwaggerInfo.BasePath = ""
		// 访问路径: /swagger/index.html
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 注册 /api/v1 分组路由，所有 v1 接口都需要经过 JWT 认证中间件
	registerRouters(r, "/api/v1", apiV1RouterFns, auth.Middleware())

	return r
}

// registerRouters 将路由函数注册到指定的分组路径下，同时支持添加额外的中间件。
func registerRouters(r *gin.Engine, groupPath string, routerFns []func(*gin.RouterGroup), handlers ...gin.HandlerFunc) {
	// 创建路由分组，并应用传入的中间件
	rg := r.Group(groupPath, handlers...)
	// 遍历执行所有路由注册函数
	for _, fn := range routerFns {
		fn(rg)
	}
}
