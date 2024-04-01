package messageQueue

import (
	"dataTool/initialize/global"
	"dataTool/pkg/rabbitmqUtils"
	"dataTool/pkg/redisUtils"
)

func RabbitmqInit() {
	global.RabbitCache = rabbitmqUtils.NewRabbitMQ("redisCache", "", "")
	RabbitmqConsume()
}

func RabbitmqConsume() {
	global.RabbitCache.ConsumeSimple(redisUtils.Redis{}.DeletedValue)
}
