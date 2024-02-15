package model

import (
	"time"
)

type User struct {
	Id       int64  `db:"id" json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name     string `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Account  string `db:"account" json:"account" gorm:"column:account;type:varchar(20);not null"`
	Password string `db:"password" json:"password" gorm:"column:password;type:varchar(60);not null"`
	Role     []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	Id     int64  `json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name   string `json:"name" gorm:"column:name;type:varchar(20);not null"`
	UserId int64
	User   []User `gorm:"many2many:user_roles;"`
}

type Api struct {
	Id         int64      `json:"id" gorm:"column:id;type:bigint(20);primaryKey;not null"`
	Name       string     ` json:"name" gorm:"column:name;type:varchar(20);not null"`
	Url        string     ` json:"url" gorm:"column:url;type:varchar(20);not null"`
	Method     string     ` json:"method" gorm:"column:method;type:varchar(10);not null"`
	Desc       string     ` json:"desc" gorm:"column:desc;type:varchar(144)"`
	CreateTime *time.Time ` json:"createTime" gorm:"column:createTime;type:datetime;not null"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:updateTime;type:datetime"`
	RoleAPIs   []RoleAPI  ` json:"roleAPIs" gorm:"many2many:role_apis;"`
}

type RoleAPI struct {
	RoleID int64 `gorm:"column:role_id;index;constraint:OnDelete:CASCADE" json:"role_id"`
	APIID  int64 `gorm:"column:api_id;index;constraint:OnDelete:CASCADE" json:"api_id"`
}

//
//func CreateApi(api model.Api) utils.Response {
//	tx := global.ApiTable.Begin()
//	if api.Name == " " || api.Url == " " || len(api.Url) >= 20 || len(api.Name) >= 10 || (api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
//		return utils.ErrorMess("参数错误", nil)
//	}
//	var apiDB model.Api
//	res := tx.Where("name = ?", api.Name).Or("url = ? and method = ?", api.Url, api.Method).First(&apiDB)
//	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
//		tx.Rollback()
//		return utils.ErrorMess("查重错误", res.Error.Error())
//	}
//	if res.Error == gorm.ErrRecordNotFound {
//		//api.CreateTime = utils.GetNowTime()
//		api.Id = global.ApiSnowFlake.Generate().Int64()
//		res = tx.Create(&api)
//		if res.Error != nil {
//			tx.Rollback()
//			return utils.ErrorMess("失败", res.Error.Error())
//		}
//		tx.Commit()
//		return utils.SuccessMess("成功", res.RowsAffected)
//	}
//	return utils.ErrorMess("创建失败", "此api已存在")
//}
