package main

import (
	"dataTool/initialize"
	"dataTool/internal/router"
)

func init() {
	initialize.InitConfig()
}

func main() {
	////go ticker.CornTicker()
	//global.MysqlClient.AutoMigrate(model.User{}, model.Role{})
	engine := router.GetEngine()
	if err := engine.Run(":8095"); err != nil {
		panic(err)
	}
}
