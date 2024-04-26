package router

import (
	handler "code_snippet/internal/handle/user"
	"code_snippet/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RouteUser(r *gin.Engine) {

	route := r.Group("/api")
	{ // 用户路由
		// 获取用户信息
		route.GET("/user/Information", middleware.JWTAuthMiddleware(), handler.User.GetInformation)
		route.POST("/user/avatar", middleware.JWTAuthMiddleware(), handler.User.PostAvatar)
		route.PUT("/user/Information", middleware.JWTAuthMiddleware(), handler.User.PutInformation)

		// 登录
		route.POST("/user/login", handler.User.Login)

		// 登录退出
		route.GET("/user/login/exit", middleware.JWTAuthMiddleware(), handler.User.Exit)

		// 注册
		route.POST("/user/register", handler.User.Register)

		// 获取验证码
		route.GET("/user/captcha", handler.User.Captcha)

	}
}
