package model

type User struct {
	Id       int64  `db:"id" json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name     string `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Account  string `db:"account" json:"account" gorm:"column:account;type:varchar(20);not null"`
	Password string `db:"password" json:"password" gorm:"column:password;type:varchar(60);not null"`
	Role     []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	Id   int64  `json:"id" gorm:"column:id;type:bigint;primarykey;not null"`
	Name string `json:"name" gorm:"column:name;type:varchar(20);not null"`
	User []User `gorm:"many2many:user_roles;"`
	Api  []Api  `gorm:"many2many:role_apis;"`
}

type Api struct {
	Id     int64  `json:"id" gorm:"column:id;type:bigint(20);primaryKey;not null"`
	Name   string ` json:"name" gorm:"column:name;type:varchar(20);not null"`
	Url    string ` json:"url" gorm:"column:url;type:varchar(20);not null"`
	Method string ` json:"method" gorm:"column:method;type:varchar(10);not null"`
	Desc   string ` json:"desc" gorm:"column:desc;type:varchar(144)"`
	//CreateTime *time.Time ` json:"createTime" gorm:"column:createTime;type:datetime;not null"`
	//UpdateTime *time.Time `json:"updateTime" gorm:"column:updateTime;type:datetime"`
	Role []Role `gorm:"many2many:role_apis;"`
}
