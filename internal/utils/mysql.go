package utils

import (
	"code_snippet/internal/config"
	"code_snippet/internal/model"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var XormDb *xorm.Engine

func MysqlDbInit() {
	fmt.Println("初始化并连接 MySQL 数据库")
	// 获取配置实例
	configs := config.Instance().MysqlDatabase
	var err error
	XormDb, err = xorm.NewEngine("mysql", configs.Master)

	if err != nil {
		fmt.Printf("MySQL 数据库连接错误: %v", err)
		return
	}

	// 测试连接
	err = XormDb.Ping()
	if err != nil {
		fmt.Printf("MySQL 数据库连接错误: %v", err)
		return
	}

	// 设置时区
	XormDb.DatabaseTZ = time.Local
	XormDb.TZLocation = time.Local

	// 设置连接池
	XormDb.SetMaxIdleConns(10)
	XormDb.SetMaxOpenConns(30)

	// 设置数据表映射
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "")
	XormDb.SetTableMapper(tbMapper)

	// 打开调试模式和日志记录
	if configs.Debug {
		XormDb.ShowSQL(true)
		XormDb.Logger().SetLevel(core.LOG_DEBUG)
	}

	// 构建数据表
	model.Model(XormDb)

	fmt.Println("MySQL 数据库连接成功")
}
