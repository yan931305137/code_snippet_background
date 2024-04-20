package router

import (
	handler "code_snippet/internal/handle/ai"
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RouteAi(r *gin.Engine) {

	route := r.Group("/api")
	{ // ai路由
		route.GET("/aiKnow", middleware.JWTAuthMiddleware(), handler.Ai.AiKnow)
	}
}
