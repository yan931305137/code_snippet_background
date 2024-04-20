package main

import (
	route "code_snippet/internal/router"
	utils2 "code_snippet/internal/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	port = ":8099"
)

func main() {

	// 注册服务
	r := gin.Default()
	// 链接Mysql数据库
	utils2.MysqlDbInit()
	// 链接Redis数据库
	utils2.RedisDbInit()
	// 启动路由
	route.Routes(r)
	// 启动服务

	if err := r.Run(port); err != nil {
		return
	}

	log.Info("服务已启动!")
}
