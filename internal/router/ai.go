package router

import (
	handler "code_snippet/internal/handle/ai"
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RouteAi(r *gin.Engine) {

	route := r.Group("/api", middleware.JWTAuthMiddleware())
	{ // ai路由
		route.GET("/aiKnow", handler.Ai.GetAiKnow)
		route.PUT("/aiKnow", handler.Ai.PutAiKnow)
		route.DELETE("/aiKnow/deleteList", handler.Ai.DeleteListAiKnow)

	}
}
