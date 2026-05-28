package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/app"
	"github.com/go-dev-frame/sponge/pkg/httpsrv"

	"pledge-be/internal/config"
	"pledge-be/internal/routers"
)

var _ app.IServer = (*httpServer)(nil)

// httpServer HTTP 服务器结构体，封装了监听地址和 sponge 的 Server 实例
type httpServer struct {
	addr   string
	server *httpsrv.Server
}

// Start 启动 HTTP 服务并阻塞等待，返回时记录错误
func (s *httpServer) Start() error {
	if err := s.server.Run(); err != nil {
		return fmt.Errorf("run %s service error: %v", s.server.Scheme(), err)
	}
	return nil
}

// Stop 优雅关闭 HTTP 服务，设置 3 秒超时上下文
func (s *httpServer) Stop() error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second) //nolint
	return s.server.Shutdown(ctx)
}

// String 返回 HTTP 服务的协议类型和监听地址摘要信息
func (s *httpServer) String() string {
	return s.server.Scheme() + " service address is " + s.addr
}

// newServer 根据 TLS 配置创建 httpsrv.Server，支持自签名、自动加密、外部证书和无 TLS 四种模式
func newServer(server *http.Server, tls config.TLS) *httpsrv.Server {
	var c *httpsrv.Server
	switch httpsrv.Mode(tls.EnableMode) {
	case httpsrv.ModeTLSSelfSigned:
		c = httpsrv.New(server, httpsrv.NewTLSSelfSignedConfig())
	case httpsrv.ModeTLSEncrypt:
		c = httpsrv.New(server,
			httpsrv.NewTLSEAutoEncryptConfig(
				tls.Domain,
				tls.Email,
				// enable http redirect to https, port 80 to 443, default is false
				//httpsrv.WithTLSEncryptEnableRedirect(),
			),
		)
	case httpsrv.ModeTLSExternal:
		c = httpsrv.New(server, httpsrv.NewTLSExternalConfig(tls.CertFile, tls.KeyFile))
	default:
		c = httpsrv.New(server)
	}
	return c
}

// NewHTTPServer 创建 HTTP 服务实例，根据环境设置 Gin 模式，配置路由和超时参数
func NewHTTPServer(addr string, opts ...HTTPOption) app.IServer {
	o := defaultHTTPOptions()
	o.apply(opts...)

	if o.isProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := routers.NewRouter()
	server := &http.Server{
		Addr:    addr,
		Handler: router,
		//ReadTimeout:    time.Second*30,
		//WriteTimeout:   time.Second*60,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}

	return &httpServer{
		addr:   addr,
		server: newServer(server, o.tls),
	}
}
