package timer

import (
	"dataTool/internal/controller/timer/sukonCloud"
	"dataTool/pkg/utils"
)

func SuKonCloudTimer() {
	//if runtime.GOOS != "linux" {
	//	return
	//}
	utils.SukonToken() //更新全局变量SuKon-Token
	//c := cron.New()                                               //新建一个定时任务对象
	//_ = c.AddFunc(global.Spec, utils.SukonToken)                  //定时获取token
	//_ = c.AddFunc("0 */1 * * * *", sukonCloud.SuKonCloudProjects) //每分钟存储生产工艺数据
	//c.Start()
	//
	//select {} //阻塞住,保持程序运行
	sukonCloud.SuKonCloudProjects()
}

//
//func SukouCloudData() {
//	var producerWg sync.WaitGroup
//
//	SukouCloudProducer
//}
