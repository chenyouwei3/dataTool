package initialize

import "dataTool/initialize/config"

func InitConfig() {
	config.MongodbInit(*config.Config.Mongodb)
	config.MysqlInit(*config.Config.Mysql)
	config.RedisInit(*config.Config.Redis)
	config.SnowFlakeInit()
	config.LogInit()
}
