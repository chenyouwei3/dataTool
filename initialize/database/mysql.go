package database

import (
	"dataTool/initialize/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// MysqlInit 初始化mysql
func MysqlInit() {
	var err error
	global.MysqlClient, err = gorm.Open(mysql.Open(global.MysqlAddress), &gorm.Config{})
	if err != nil {
		log.Fatalln("Mysql数据库连接失败:", err)
	}
	sqlDB, err := global.MysqlClient.DB()
	if err != nil {
		log.Fatalln("连接池创建失败")
	}
	sqlDB.SetMaxIdleConns(10)                  //最大空闲连接数
	sqlDB.SetMaxOpenConns(10)                  //最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute * 15) //设置连接空闲超时
	{
		global.UserTable = global.MysqlClient.Table("user")
		global.RoleTable = global.MysqlClient.Table("role")
		global.ApiTable = global.MysqlClient.Table("api")
		global.RoleApiTable = global.MysqlClient.Table("role_apis")
	}
	fmt.Println("mysql连接成功")
}
