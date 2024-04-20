package utils

import (
	"code_snippet/internal/config"
	"github.com/go-redis/redis/v8"
	"strconv"

	"context"
	"fmt"
	"time"
)

var ctx = context.Background()
var RD *redis.Client

// RedisDbInit 初始化并连接 Redis 数据库
func RedisDbInit() {
	fmt.Println("初始化并连接 Redis 数据库")
	// 获取 Redis 配置实例
	redisConfig := config.Instance().RedisDatabase
	// 创建 Redis 客户端
	RD = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	// 测试连接
	if err := RD.Ping(ctx).Err(); err != nil {
		fmt.Errorf("连接 Redis 数据库失败: %w", err)
	}

	fmt.Println("Redis 数据库连接成功")
}

// SaveTokenToRedis 将 id 存储到 Redis 中
func SaveTokenToRedis(token string, id int) error {
	// 设置 id 到 Redis，并设置过期时间
	expiration := 24 * time.Hour // 设置过期时间为 24 小时
	if err := RD.Set(ctx, token, id, expiration).Err(); err != nil {
		return fmt.Errorf("无法存储 id 到 Redis: %v", err)
	}

	fmt.Println("Token 存储到 Redis 成功")
	return nil
}

// GetTokenToRedis 从Redis 中将 id 取出
func GetTokenToRedis(token string) (int, error) {
	// 取出id
	id, err := RD.Get(ctx, token).Result()
	if err != nil {
		return 0, fmt.Errorf("无法从 Redis 取出 id: %v", err)
	}
	i, _ := strconv.ParseInt(id, 10, 32)
	fmt.Println("从 Redis 取出 Token 成功")
	return int(i), nil
}
