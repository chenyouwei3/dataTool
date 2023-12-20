package main

import (
	"dataTool/internal/router"
)

func init() {
	InitDataBase()
}

func main() {
	engine := router.GetEngine()
	if err := engine.Run(":8091"); err != nil {
		panic(err)
	}
}
