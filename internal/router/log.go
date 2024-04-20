package router

import (
	handler "code_snippet/internal/handle/log"
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RouteLog 设置日志路由
func RouteLog(r *gin.Engine) {
	route := r.Group("/api", middleware.JWTAuthMiddleware())
	{
		// 日志列表
		route.GET("/logs/:size/:num", handler.Log.LogList)
		// 添加日志
		route.POST("/log", handler.Log.LogAdd)
		// 删除日志
		route.DELETE("/log", handler.Log.LogDelete)
		// 修改日志
		route.PUT("/log", handler.Log.LogAlter)
	}
}
