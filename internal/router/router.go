package router

import (
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Routes 设置所有路由
func Routes(r *gin.Engine) {
	// 使用跨域中间件
	r.Use(middleware.Cros())

	//r.Use(middleware.JWTAuthMiddleware())
	// 记录请求日志
	r.Use(middleware.LogRequest)

	// 设置各个模块的路由
	RouteUser(r)
	RouteLog(r)
	RouteAi(r)
	RouteCode(r)
}
