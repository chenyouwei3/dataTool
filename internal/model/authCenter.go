package model

import "time"

type User struct {
	Id         int64  `db:"id" json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name       string `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Account    string `db:"account" json:"account" gorm:"column:account;type:varchar(20);not null"`
	Password   string `db:"password" json:"password" gorm:"column:password;type:varchar(60);not null"`
	Salt       string `db:"salt" json:"salt" gorm:"column:salt;type:varchar(60);not null"`
	RoleId     int64  `db:"roleId" json:"roleId" gorm:"column:roleId;type:text；not null"`
	CreateTime string `db:"createTime" json:"createTime" gorm:"column:createTime;type:datetime;not null"`
	UpdateTime string `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:datetime"`
	Role       Role
}

//	type Role struct {
//		Id         int64      `json:"id" gorm:"foreignKey:id;constraint:OnDelete:CASCADE"`
//		Name       string     `json:"name" gorm:"column:name;type:varchar(20);not null"`
//		CreateTime *time.Time `json:"createTime" gorm:"column:createTime;type:datetime;;not null"`
//		UpdateTime *time.Time `json:"updateTime" gorm:"column:updateTime;type:datetime"`
//		RoleAPIs   []RoleAPI  `json:"roleAPIs "gorm:"many2many:role_apis;foreignKey:ID;constraint:OnDelete:CASCADE" `
//	}
//
//	type Api struct {
//		Id         int64      `json:"id" gorm:"column:id;type:bigint(20);primaryKey;not null"`
//		Name       string     ` json:"name" gorm:"column:name;type:varchar(20);not null"`
//		Url        string     ` json:"url" gorm:"column:url;type:varchar(20);not null"`
//		Method     string     ` json:"method" gorm:"column:method;type:varchar(10);not null"`
//		Desc       string     ` json:"desc" gorm:"column:desc;type:varchar(144)"`
//		CreateTime *time.Time ` json:"createTime" gorm:"column:createTime;type:datetime;not null"`
//		UpdateTime *time.Time `json:"updateTime" gorm:"column:updateTime;type:datetime"`
//		RoleAPIs   []RoleAPI  ` json:"roleAPIs" gorm:"many2many:role_apis;"`
//	}
//
//	type RoleAPI struct {
//		RoleID int64 `gorm:"column:role_id;index;constraint:OnDelete:CASCADE" json:"role_id"`
//		APIID  int64 `gorm:"column:api_id;index;constraint:OnDelete:CASCADE" json:"api_id"`
//	}
type Role struct {
	Id         int64      `json:"id" gorm:"foreignKey:id;constraint:OnDelete:CASCADE"`
	Name       string     `json:"name" gorm:"column:name;type:varchar(20);not null"`
	CreateTime *time.Time `json:"createTime" gorm:"column:createTime;type:datetime;;not null"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:updateTime;type:datetime"`
	RoleAPIs   []RoleAPI  `json:"roleAPIs "gorm:"many2many:role_apis;constraint:OnDelete:CASCADE" `
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
