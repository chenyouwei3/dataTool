package ticker

import (
	"dataTool/initialize/global"
	"dataTool/internal/controller/ticker/sukonCloud"
	"dataTool/pkg/utils"
	"github.com/robfig/cron"
)

func CornTicker() {
	//if utils.IsProd() {
	//	return
	//}
	//utils.SukonToken() //更新全局变量SuKon-Token
	//sukonCloud.SuKonCloudProjects()
	c := cron.New()                                               //新建一个定时任务对象
	_ = c.AddFunc(global.Spec, utils.SukonToken)                  //定时获取token
	_ = c.AddFunc("0 */1 * * * *", sukonCloud.SuKonCloudProjects) //每分钟存储生产工艺数据
	c.Start()
	select {}
}

func addCornFunc(c *cron.Cron, spec string, cmd func()) {
	err := c.AddFunc(spec, cmd)
	if err != nil {

	}
}
