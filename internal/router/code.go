package router

import (
	handler "code_snippet/internal/handle/code"
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RouteCode(r *gin.Engine) {

	route := r.Group("/api")
	{ // code路由
		route.POST("/code/PostCode", middleware.JWTAuthMiddleware(), handler.Code.PostCode)
		route.GET("/code/GetMyCode", middleware.JWTAuthMiddleware(), handler.Code.GetMyCode)
		route.POST("/code/SearchGetCodes", handler.Code.SearchGetCodes)
		route.GET("/code/GetCodes", handler.Code.GetCodes)
	}
}
