package middleware

import (
	"code_snippet/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 基于JWT认证的中间件   验证token的中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := utils.GetToken(c) //获取请求中头部的token
		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token为空",
			})
			c.Abort() //授权失败，调用Abort以确保没有调用此请求的其余处理程序
			return
		}

		if _, err := utils.ParseToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
	}
}
