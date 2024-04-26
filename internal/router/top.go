package router

import (
	handler "code_snippet/internal/handle/top"
	"github.com/gin-gonic/gin"
)

func RouteTop(r *gin.Engine) {

	route := r.Group("/api")
	{
		// top路由
		route.GET("/top/slider", handler.Top.GetTopSlider)
		route.GET("/top/hot", handler.Top.GetTopHop)
		route.GET("/top/new", handler.Top.GetTopNew)
		route.GET("/top/focus", handler.Top.GetTopFocus)
	}
}
