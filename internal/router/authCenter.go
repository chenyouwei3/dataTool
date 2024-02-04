package router

import (
	"dataTool/internal/controller"
	"github.com/gin-gonic/gin"
)

func AuthCenterRouter(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.POST("/create", controller.CreateApi)     //增api@
		api.DELETE("/deleted", controller.DeletedApi) //删api@
		api.PUT("/update", controller.UpdatedApi)     //改api@
		api.GET("/get", controller.GetApi)            //查api@
	}
	role := engine.Group("/role")
	{
		role.POST("/create", controller.CreateRole)              //增role@
		role.POST("/association", controller.AssociationRoleApi) //增加role_api关系@
		role.DELETE("/deleted", controller.DeletedRole)          //删role@
		//role.PUT("/update", controller.UpdatedRole)              //改role
		role.GET("/get", controller.GetRole) //查role
	}
}
