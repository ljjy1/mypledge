package server

import (
	"pledge-be/internal/config"
)

// HTTPOption setting up http
type HTTPOption func(*httpOptions)

// httpOptions HTTP 服务的内部配置选项
type httpOptions struct {
	isProd bool
	tls    config.TLS
}

// defaultHTTPOptions 返回默认的 HTTP 配置选项
func defaultHTTPOptions() *httpOptions {
	return &httpOptions{
		isProd: false,
	}
}

// apply 依次应用传入的 HTTPOption 配置函数到 httpOptions
func (o *httpOptions) apply(opts ...HTTPOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithHTTPIsProd setting up production environment markers
func WithHTTPIsProd(isProd bool) HTTPOption {
	return func(o *httpOptions) {
		o.isProd = isProd
	}
}

// WithHTTPTLS setting up tls
func WithHTTPTLS(tls config.TLS) HTTPOption {
	return func(o *httpOptions) {
		o.tls = tls
	}
}
