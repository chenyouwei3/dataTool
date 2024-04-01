package initialize

import (
	"dataTool/initialize/config/database"
	"dataTool/initialize/config/messageQueue"
	"dataTool/initialize/config/socketServer"
	"dataTool/initialize/config/system"
)

func InitConfig() {
	database.MongodbInit(*system.Config.Mongodb)
	database.MysqlInit(*system.Config.Mysql)
	database.RedisInit(*system.Config.Redis)
	messageQueue.RabbitmqInit()
	system.SnowFlakeInit()
	system.LogInit()
	socketServer.SocketServerStart()

}
