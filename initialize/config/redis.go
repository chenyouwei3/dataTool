package config

import (
	"dataTool/initialize/global"
	"fmt"
	"github.com/go-redis/redis"
)

// RedisInit 开启RedisPool
func RedisInit(config RedisConfig) {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.Address,      // Redis 服务器地址
		Password:     config.Password,     // Redis 服务器密码
		DB:           config.DB,           // Redis 数据库索引
		PoolSize:     config.PoolSize,     // 连接池大小
		MinIdleConns: config.MinIdleConns, // 最小空闲连接数
	})
	ping, err := global.RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败", ping, err)
		return
	}
	fmt.Println("redis连接成功", ping)
}
