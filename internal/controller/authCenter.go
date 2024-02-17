package controller

import (
	"dataTool/internal/model"
	"dataTool/internal/service"
	"dataTool/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.CreateUser(user))
}

func DeletedUser(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.DeletedUser(id))
}

func UpdatedUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.UpdatedUser(user))
}

func GetUser(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetUser(name, currPage, pageSize, startTime, endTime))
}

func CreateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.CreateApi(api))
}

func DeletedApi(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.DeletedApi(id))
}

func UpdatedApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.UpdateApi(api))
}

func GetApi(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetApi(name, currPage, pageSize, startTime, endTime))
}

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
		return
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}

func AssociationRoleApi(c *gin.Context) {
	roleIdStr := c.Query("roleId")
	apiIdsStr := c.QueryArray("apiIds")
	// 将 roleId 转为 int64
	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		// 处理转换错误
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roleId"})
		return
	}
	// 将 apiIds 转为 []int64
	var apiIds []int64
	for _, apiIdStr := range apiIdsStr {
		apiId, err := strconv.ParseInt(apiIdStr, 10, 64)
		if err != nil {
			// 处理转换错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid apiId"})
			return
		}
		apiIds = append(apiIds, apiId)
	}
	c.JSON(http.StatusOK, service.AssociationRoleApi(roleId, apiIds))
}

func DeletedRole(c *gin.Context) {
	id := c.Query("id")
	if id == " " {
		c.JSON(e.ParameterError, e.GetMsg(e.ParameterError))
		return
	}
	c.JSON(http.StatusOK, service.DeletedRole0(id))
}

//func UpdatedRole(c *gin.Context) {
//	var role model.Role
//	if err := c.Bind(&role); err != nil {
//		c.JSON(e.ParameterStructError, e.GetMsg(e.ParameterStructError))
//		return
//	}
//	c.JSON(http.StatusOK, service.UpdateRole(role))
//}

func GetRole(c *gin.Context) {
	name := c.Query("name")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	c.JSON(http.StatusOK, service.GetRole(name, currPage, pageSize, startTime, endTime))
}
