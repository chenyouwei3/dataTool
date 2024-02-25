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
	if err := engine.Run(":8095"); err != nil {
		panic(err)
	}
}
