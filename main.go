package main

import (
	"dataTool/initialize"
	"dataTool/internal/controller/timer"
	"dataTool/internal/router"
)

func init() {
	initialize.InitDataBase()
	initialize.InitChan()
}

func main() {
	go timer.SuKonCloudTimer()
	engine := router.GetEngine()
	if err := engine.Run(":8091"); err != nil {
		panic(err)
	}
}
