package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/jwt"
)

const signKey = "pledge-be-jwt-secret-key"

// GenerateToken 生成 JWT token
func GenerateToken(uid string) (string, error) {
	_, token, err := jwt.GenerateToken(uid, jwt.WithGenerateTokenSignKey([]byte(signKey)))
	return token, err
}

// Middleware 返回 JWT 认证中间件
func Middleware() gin.HandlerFunc {
	return middleware.Auth(middleware.WithSignKey([]byte(signKey)))
}
