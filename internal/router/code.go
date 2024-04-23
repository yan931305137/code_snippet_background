package router

import (
	handler "code_snippet/internal/handle/code"
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RouteCode(r *gin.Engine) {

	route := r.Group("/api", middleware.JWTAuthMiddleware())
	{ // code路由
		route.POST("/code/PostCode", handler.Code.PostCode)
		route.GET("/code/GetMyCode", handler.Code.GetMyCode)
	}
}
