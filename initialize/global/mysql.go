package global

import "gorm.io/gorm"

var (
	MysqlClient   *gorm.DB
	UserTable     *gorm.DB
	UserRoleTable *gorm.DB
	RoleTable     *gorm.DB
	ApiTable      *gorm.DB
	RoleApiTable  *gorm.DB
	LogTable      *gorm.DB
)
