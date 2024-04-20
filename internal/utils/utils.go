package utils

import (
	"github.com/gin-gonic/gin"
)

// GetToken 获取客户端Token
func GetToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("Authorization") //获取请求中头部的token
	return token
}
