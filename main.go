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
	if err := engine.Run(":8093"); err != nil {
		panic(err)
	}
}
