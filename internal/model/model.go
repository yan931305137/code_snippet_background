package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

func Model(XormDb *xorm.Engine) {
	//直接通过结构体，在数据库中创建对应的表【同步结构体与数据表】
	if err := XormDb.Sync(new(User)); err != nil {
		fmt.Println("用户表结构同步失败")
	}
	if err := XormDb.Sync(new(Log)); err != nil {
		fmt.Println("日志表结构同步失败")
	}
}
