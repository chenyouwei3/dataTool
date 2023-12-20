package router

import "github.com/gin-gonic/gin"

func GetEngine() *gin.Engine {
	engine := gin.Default()
	AuthCenterRouter(engine)
	return engine
}
