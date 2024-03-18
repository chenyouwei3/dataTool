package router

import (
	"dataTool/internal/controller"
	"dataTool/internal/middleware"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	//限流/路由日志/跨域
	engine.Use(middleware.OperationLogMiddleware(), middleware.CorsMiddleware()) //跨域
	engine.POST("/login", controller.Login)
	//权限/jwt/cookie/session
	engine.Use(middleware.AuthCookieMiddleware())
	AuthCenterRouter(engine)
	return engine
}
