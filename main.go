package main

import (
	"dataTool/initialize"
	"dataTool/internal/controller/ticker"
	"dataTool/internal/router"
)

func init() {
	initialize.InitConfig()
}

func main() {
	go ticker.CornTicker()
	engine := router.GetEngine()
	if err := engine.Run(":8099"); err != nil {
		panic(err)
	}
	//utils.SukonToken() //更新全局变量SuKon-Token
	//sukonCloud.SuKonCloudProjects()
}
