package router

import (
	"dataTool/internal/controller"
	"github.com/gin-gonic/gin"
)

func AuthCenterRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.POST("/create", controller.CreateApi)     //增
		api.DELETE("/deleted", controller.DeletedApi) //删
		api.PUT("/update", controller.UpdatedApi)     //改
		api.GET("/get", controller.GetApi)            //查
	}
	//role := engine.Group("/role")
	//{
	//	role.POST("/create", controller.CreateRole) //增
	//	role.DELETE("/deleted", controller.DeletedRole) //删
	//	role.PUT("/update", controller.UpdatedRole)     //改
	//	role.GET("/get", controller.GetRole)            //查
	//}
}
