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

type Role struct {
	Id         int64      `db:"id" json:"id" gorm:"column:id;type:bigint;primaryKey;not null"`
	Name       string     `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	CreateTime *time.Time `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
	UpdateTime *time.Time `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:date"`
	RoleAPIs   []RoleAPI  `gorm:"many2many:role_apis;" json:"roleAPIs"`
}

type Api struct {
	Id         int64      `db:"id" json:"id" gorm:"column:id;type:bigint(20);primaryKey;not null"`
	Name       string     `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Url        string     `db:"url" json:"url" gorm:"column:url;type:varchar(20);not null"`
	Method     string     `db:"method" json:"method" gorm:"column:method;type:varchar(10);not null"`
	Desc       string     `db:"desc" json:"desc" gorm:"column:desc;type:varchar(144)"`
	CreateTime *time.Time `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
	UpdateTime *time.Time `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:varchar(20)"`
	RoleAPIs   []RoleAPI  `gorm:"many2many:role_apis;" json:"roleAPIs"`
}

type RoleAPI struct {
	RoleID int64 `gorm:"column:role_id;index"`
	APIID  int64 `gorm:"column:api_id;index"`
}
