package config

import (
	"code_snippet/internal/types"
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	instance *types.Config
	once     sync.Once
)

// Instance 返回配置文档实例
func Instance() *types.Config {
	once.Do(func() {
		var conf types.Config
		// 获取当前工作目录
		path, err := os.Getwd()
		if err != nil {
			fmt.Println("获取当前工作目录失败:", err)
			return
		}
		// 拼接配置文件路径
		filePath := path + "/internal/config/config.toml"
		// 解析配置文件
		if _, err := toml.DecodeFile(filePath, &conf); err != nil {
			fmt.Println("解析配置文件失败:", err)
			return
		}
		instance = &conf
	})
	return instance
}
