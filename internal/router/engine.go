package router

import (
	"dataTool/internal/middleware"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.OperationLogMiddleware(), middleware.CorsMiddleware())
	AuthCenterRouter(engine)
	return engine
}
