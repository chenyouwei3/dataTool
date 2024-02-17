package config

import (
	"dataTool/initialize/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func MysqlInit(config MysqlConfig) {
	var err error
	global.MysqlClient, err = gorm.Open(mysql.Open(config.Address), &gorm.Config{})
	if err != nil {
		log.Fatalln("Mysql数据库连接失败:", err)
	}
	sqlDB, err := global.MysqlClient.DB()
	if err != nil {
		log.Fatalln("Mysql连接池创建失败")
	}
	//sqlDB.SetMaxIdleConns(config.SetMaxIdleConns)       //最大空闲连接数
	//sqlDB.SetMaxOpenConns(config.SetMaxOpenConns)       //最大连接数
	sqlDB.SetConnMaxLifetime(config.SetConnMaxLifetime) //设置连接空闲超时
	{
		global.UserTable = global.MysqlClient.Table("users")
		global.RoleTable = global.MysqlClient.Table("roles")
		global.UserRoleTable = global.MysqlClient.Table("user_roles")
		global.ApiTable = global.MysqlClient.Table("api")
		global.LogTable = global.MysqlClient.Table("log")
	}
	fmt.Println("mysql连接成功")
}
