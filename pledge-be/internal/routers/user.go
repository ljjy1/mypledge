package routers

import (
	"github.com/gin-gonic/gin"

	"pledge-be/internal/handler"
)

// init 自动向 apiV1RouterFns 注册用户模块的路由函数
func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		userRouter(group, handler.NewUserHandler())
	})
}

// userRouter 注册用户 CRUD 路由
func userRouter(group *gin.RouterGroup, h handler.UserHandler) {
	g := group.Group("/user")

	// JWT 认证参考文档: https://go-sponge.com/component/transport/gin.html#jwt-authorization-middleware

	// 以下所有路由默认都使用 JWT 认证，也可以使用 middleware.Auth(middleware.WithExtraVerify(fn))
	//g.Use(middleware.Auth())

	// 如果不需要所有路由都走 JWT 认证，可以单独为某些路由添加认证中间件。
	// 这种情况下，不要使用上面的 g.Use(middleware.Auth())

	g.POST("/", h.Create)          // [post] /api/v1/user
	g.DELETE("/:id", h.DeleteByID) // [delete] /api/v1/user/:id
	g.PUT("/:id", h.UpdateByID)    // [put] /api/v1/user/:id
	g.GET("/:id", h.GetByID)       // [get] /api/v1/user/:id
	g.POST("/list", h.List)        // [post] /api/v1/user/list
}
