package database

import (
	"dataTool/config/global"
	"fmt"
	"github.com/go-redis/redis"
)

// RedisInit 开启RedisPool
func RedisInit() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:         global.RedisAddress, // Redis 服务器地址
		Password:     "",                  // Redis 服务器密码
		DB:           0,                   // Redis 数据库索引
		PoolSize:     20,                  // 连接池大小
		MinIdleConns: 5,                   // 最小空闲连接数
	})
	ping, err := global.RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败", ping, err)
		return
	}
	fmt.Println("redis连接成功", ping)
}
