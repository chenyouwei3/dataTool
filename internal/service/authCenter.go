package service

import (
	"dataTool/initialize/global"
	"dataTool/internal/model"
	"dataTool/pkg/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func CreateUser(user model.User) utils.Response {
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		//查询账号重复
		var userDB model.User
		if err := tx.Where("account = ?", user.Account).First(&userDB).Error; (err != nil && !errors.Is(err, gorm.ErrRecordNotFound)) || userDB.Account == user.Account {
			return fmt.Errorf("查询账号错误:%w", err)
		}
		// 查询角色是否存在
		var roleDB []model.Role
		if err := global.RoleTable.Where("id IN ?", extractRoleIDs(user.Role)).Find(&roleDB).Error; err != nil {
			return fmt.Errorf("查询角色错误:%w", err)
		}
		if len(roleDB) != len(user.Role) { // 检查查询到的角色数量是否和传入的角色数量相等
			return fmt.Errorf("角色数量不相等")
		}
		//插入事务
		if err := global.RoleTable.Transaction(func(tx1 *gorm.DB) error {
			user.Id = global.ApiSnowFlake.Generate().Int64()
			if err := tx.Create(&user).Error; err != nil {
				return fmt.Errorf("创建角色失败:%w", err)
			}
			return nil
		}); err != nil {
			return fmt.Errorf("创建角色事务失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "1")
}

func extractRoleIDs(roles []model.Role) []int64 { // 提取角色ID列表(辅助函数)
	ids := make([]int64, len(roles))
	for i, role := range roles {
		ids[i] = role.Id
	}
	return ids
}

func DeletedUser(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		tx0 := global.UserRoleTable.Begin()
		if err := tx0.Model(&model.User{Id: id}).Association("Role").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("清除关联失败:%w", err)
		}
		tx0.Commit()
		// 删除用户记录
		if err := tx.Delete(&model.User{}, id).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func UpdatedUser(user model.User) utils.Response {
	if err := global.UserTable.Transaction(func(tx *gorm.DB) error {
		var userDB model.User
		if err := tx.Where("id = ?", user.Id).First(&userDB).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("查询失败%w", err)
		}
		userDB.Name = user.Name
		if err := tx.Save(&userDB).Error; err != nil {
			return fmt.Errorf("更新角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("修改用户成功", user.Id)
}

func GetUser(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err.Error())
	}
	tx := global.UserTable
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var userDB []model.User
	//Order("id desc")id降序排列
	res := tx.Preload("Role").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&userDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.User `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  userDB,
	})
}

func CreateRole(role model.Role) utils.Response {
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		//查询角色重复
		var roleDB model.Role
		if err := tx.Where("name = ?", roleDB.Name).Error; (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) || role.Name == roleDB.Name {
			return fmt.Errorf("角色重复:%w", err)
		}
		// 查询api是否存在
		var apiDB []model.Api
		if err := global.ApiTable.Where("id IN ?", extractRoleID(role.Api)).Find(&apiDB).Error; err != nil {
			return fmt.Errorf("查询api错误:%w", err)
		}
		if len(apiDB) != len(role.Api) { // 检查查询到的api数量是否和传入的api数量相等
			return fmt.Errorf("api数量不相等")
		}
		//插入事务
		if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
			role.Id = global.RoleSnowFlake.Generate().Int64()
			if err := tx.Create(&role).Error; err != nil {
				return fmt.Errorf("创建角色失败:%w", err)
			}
			return nil
		}); err != nil {
			return fmt.Errorf("创建角色事务失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "role")
}

func extractRoleID(apis []model.Api) []int64 { // 提取角色ID列表(辅助函数)
	ids := make([]int64, len(apis))
	for i, api := range apis {
		ids[i] = api.Id
	}
	return ids
}

func DeletedRole(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("字符串转化整数失败", err.Error())
	}
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		tx0 := global.RoleApiTable.Begin()
		if err := tx0.Model(&model.Role{Id: id}).Association("Api").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("清除关联失败:%w", err)
		}
		tx0.Commit()
		// 删除角色记录
		if err := tx.Delete(&model.Role{}, id).Error; err != nil {
			return fmt.Errorf("删除角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func UpdateRole(role model.Role) utils.Response {
	if err := global.RoleTable.Transaction(func(tx *gorm.DB) error {
		var roleDB model.Role
		if err := tx.Where("id = ?", role.Id).First(&roleDB).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("查询失败%w", err)
		}
		roleDB.Name = role.Name
		if err := tx.Save(&roleDB).Error; err != nil {
			return fmt.Errorf("更新角色失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("修改角色成功", role.Id)
}

func GetRole(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err.Error())
	}
	tx := global.RoleTable
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var roleDB []model.Role
	//Order("id desc")id降序排列
	res := tx.Preload("Api").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&roleDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.Role `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  roleDB,
	})
}

func CreateApi(api model.Api) utils.Response {
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		//查询角色重复
		var apiDB model.Api
		if err := tx.Where("url = ?", api.Url).First(&apiDB).Error; (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) || api.Url == apiDB.Url {
			return fmt.Errorf("api重复:%w", err)
		}
		if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
			api.Id = global.ApiSnowFlake.Generate().Int64()
			if err := tx.Create(&api).Error; err != nil {
				return fmt.Errorf("创建api失败:%w", err)
			}
			return nil
		}); err != nil {
			return fmt.Errorf("创建api事务失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "api")
}

func DeletedApi(idString string) utils.Response {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("字符串转化整数失败", err.Error())
	}
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		tx0 := global.RoleApiTable.Begin()
		if err := tx0.Model(&model.Api{Id: id}).Association("Role").Clear(); err != nil {
			tx0.Rollback()
			return fmt.Errorf("清除关联失败:%w", err)
		}
		tx0.Commit()
		// 删除api记录
		if err := tx.Delete(&model.Api{}, id).Error; err != nil {
			return fmt.Errorf("删除api失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("删除事务失败", err.Error())
	}
	return utils.SuccessMess("删除成功", id)
}

func UpdateApi(api model.Api) utils.Response {
	if err := global.ApiTable.Transaction(func(tx *gorm.DB) error {
		var apiDB model.Api
		if err := tx.Where("id = ?", api.Id).First(&apiDB).Error; err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("查询失败%w", err)
		}
		apiDB.Name = api.Name
		if err := tx.Save(&apiDB).Error; err != nil {
			return fmt.Errorf("更新api失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("修改api成功", api.Id)
}

func GetApi(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err.Error())
	}
	tx := global.ApiTable
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var apiDB []model.Api
	//Order("id desc")id降序排列
	res := tx.Preload("Role").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&apiDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error.Error())
	}
	return utils.SuccessMess("成功", struct {
		Count int64       `json:"count" bson:"count"`
		Data  []model.Api `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  apiDB,
	})
}
