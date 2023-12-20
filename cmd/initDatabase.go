package main

import (
	"context"
	"dataTool/config/database"
	"dataTool/config/global"
	"fmt"
)

func InitDataBase() {
	//defer CloseDB()
	database.MongodbInit()
	database.SnowFlakeInit()
	database.MysqlInit()
	database.RedisInit()
}

func CloseDB() {
	redisErr := global.RedisClient.Close()
	if redisErr != nil {
		fmt.Println("Error on closing Redis Service client.")
	}
	sql, MysqlErr := global.MysqlClient.DB()
	if MysqlErr != nil {
		fmt.Println("Error on closing Mysql Service client.")
	}
	sql.Close()
	MongodbErr := global.MongodbClient403.Disconnect(context.TODO()) // 延迟关闭 MongoDB 客户端连接
	if MongodbErr != nil {
		fmt.Println("Error on closing Mongodb Service client.")
	}
}
