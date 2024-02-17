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
			fmt.Println(err)
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
		return utils.ErrorMess("事务失败", err)
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
		return utils.ErrorMess("失败", err)
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
		return utils.ErrorMess("删除事务失败", err)
	}
	return utils.SuccessMess("删除成功", id)
}

func UpdatedUser(user model.User) utils.Response {
	if len(user.Name) > 20 {
		return utils.ErrorMess("字段过长", "重新更改")
	}
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
		return utils.ErrorMess("事务失败", err)
	}
	return utils.SuccessMess("修改用户成功", user.Id)
}

func GetUser(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("数据转化失败", err)
	}
	tx := global.UserTable
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var userDB []model.User
	res := tx.Preload("Role").Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&userDB).Count(&count)
	if res.Error != nil {
		return utils.ErrorMess("失败", res.Error)
	}
	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.User `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  userDB,
	})
}

func CreateApi(api model.Api) utils.Response {
	tx := global.ApiTable.Begin()
	if api.Name == " " || api.Url == " " || len(api.Url) >= 20 || len(api.Name) >= 10 || (api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("参数错误", nil)
	}
	var apiDB model.Api
	res := tx.Where("name = ?", api.Name).Or("url = ? and method = ?", api.Url, api.Method).First(&apiDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return utils.ErrorMess("查重错误", res.Error.Error())
	}
	if res.Error == gorm.ErrRecordNotFound {
		//api.CreateTime = utils.GetNowTime()
		api.Id = global.ApiSnowFlake.Generate().Int64()
		res = tx.Create(&api)
		if res.Error != nil {
			tx.Rollback()
			return utils.ErrorMess("失败", res.Error.Error())
		}
		tx.Commit()
		return utils.SuccessMess("成功", res.RowsAffected)
	}
	return utils.ErrorMess("创建失败", "此api已存在")
}

func DeletedApi(idString string) utils.Response {
	tx := global.ApiTable.Begin()
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		tx.Rollback()
		return utils.ErrorMess("失败", err.Error())
	}
	res := tx.Delete(&model.Api{}, id)
	if res.Error != nil {
		tx.Rollback()
		return utils.ErrorMess("失败", res.Error.Error())
	}
	tx.Commit()
	return utils.SuccessMess("成功", res.RowsAffected)
}

func UpdateApi(api model.Api) utils.Response {
	tx := global.ApiTable.Begin()
	if api.Id == 0 || api.Name == "" || api.Url == "" || len(api.Url) >= 20 || len(api.Name) >= 10 || (api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("失败,参数错误", nil)
	}
	var apiDB model.Api
	res := tx.Where("id=?", api.Id).First(&apiDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return utils.ErrorMess("失败,该API不存在", nil)
	}
	Temp := apiDB.CreateTime
	apiDB = api
	apiDB.CreateTime = Temp
	//apiDB.UpdateTime = utils.GetNowTime()
	res = tx.Where("id = ?", api.Id).Save(&apiDB)
	if res.Error != nil {
		tx.Rollback()
		return utils.ErrorMess("失败", res.Error.Error())
	}
	tx.Commit()
	return utils.SuccessMess("成功", res.RowsAffected)
}

func GetApi(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	tx := global.ApiTable.Begin()
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var apiDB []model.Api
	res := tx.Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&apiDB).Count(&count)
	if res.Error != nil {
		tx.Rollback()
		return utils.ErrorMess("失败", res.Error.Error())
	}
	tx.Commit()
	return utils.SuccessMess("成功", struct {
		Count int64       `json:"count" bson:"count"`
		Data  []model.Api `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  apiDB,
	})
}

func CreateRole(role model.Role) utils.Response {
	tx := global.RoleTable.Begin()
	var roleDB model.Role
	res := tx.Where("name = ?", role.Name).First(&roleDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return utils.ErrorMess("查重错误", res.Error.Error())
	}
	if res.Error == gorm.ErrRecordNotFound {
		//role.CreateTime = utils.GetNowTime()
		role.Id = global.ApiSnowFlake.Generate().Int64()
		res = tx.Create(&role)
		if res.Error != nil {
			tx.Rollback()
			return utils.ErrorMess("失败", res.Error.Error())
		}
		tx.Commit()
		return utils.SuccessMess("成功", res.RowsAffected)
	}
	tx.Rollback()
	return utils.ErrorMess("创建失败", "此role已存在")
}

func AssociationRoleApi(roleId int64, apiIds []int64) utils.Response {
	tx0 := global.RoleApiTable.Begin()
	tx1 := global.RoleTable.Begin()
	tx2 := global.ApiTable.Begin()
	// 查询要关联的 role
	var role model.Role
	if err := tx1.First(&role, roleId).Error; err != nil {
		tx1.Rollback()
		return utils.ErrorMess("查询角色失败", err.Error())
	}
	// 查询要关联的 APIs
	var apis []model.Api
	if err := tx2.Find(&apis, apiIds).Error; err != nil {
		tx2.Rollback()
		return utils.ErrorMess("查询 APIs 失败", err.Error())
	}
	// 将关联信息写入 role_apis 表
	for _, api := range apis {
		roleAPI := model.RoleAPI{RoleID: role.Id, APIID: api.Id}
		if err := tx0.Create(&roleAPI).Error; err != nil {
			tx0.Rollback()
			return utils.ErrorMess("关联失败", err.Error())
		}
	}
	tx0.Commit()
	tx1.Commit()
	tx2.Commit()
	return utils.SuccessMess("成功", "关联数据成功")
}

func DeletedRole(idString string) utils.Response {
	tx0 := global.RoleApiTable.Begin()
	tx1 := global.RoleTable.Begin()
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	var apiroleDB []model.RoleAPI
	res0 := tx0.Where("role_id = ?", id).Find(&apiroleDB)
	if res0.Error != nil {
		tx0.Rollback()
		return utils.ErrorMess("没有此role关联信息", res0.Error.Error())
	}
	for _, temp := range apiroleDB {
		if err := tx0.Delete(&model.RoleAPI{}, temp.RoleID).Error; err != nil {
			tx0.Rollback()
			return utils.ErrorMess("删除关联失败", err.Error())
		}
	}
	res1 := tx1.Delete(&model.Role{}, id)
	if res1.Error != nil {
		tx1.Rollback()
		return utils.ErrorMess("删除role失败", res1.Error.Error())
	}
	tx0.Commit()
	tx1.Commit()
	return utils.SuccessMess("成功", apiroleDB)
}

func DeletedRole0(idString string) utils.Response {
	tx1 := global.RoleTable.Begin()
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	res1 := tx1.Unscoped().Delete(&model.Role{}, id)
	if res1.Error != nil {
		tx1.Rollback()
		return utils.ErrorMess("删除role失败", res1.Error.Error())
	}
	tx1.Commit()
	return utils.SuccessMess("成功", res1.RowsAffected)
}

//func UpdateRole(role model.Role) utils.Response {
//	tx := global.RoleTable.Begin()
//
//}

func GetRole(name, currPage, pageSize, startTime, endTime string) utils.Response {
	skip, limit, err := utils.GetPage(currPage, pageSize)
	if err != nil {
		return utils.ErrorMess("失败", err.Error())
	}
	tx := global.RoleTable.Begin()
	if startTime != "" && endTime != "" {
		tx = tx.Where("createTime >= ? and createTime <=?", startTime, endTime)
	}
	var count int64
	var roleDB []model.Role
	res := tx.Order("id desc").Where("name like ?", "%"+name+"%").Limit(limit).Offset(skip).Find(&roleDB).Count(&count)
	if res.Error != nil {
		tx.Rollback()
		return utils.ErrorMess("失败", res.Error.Error())
	}
	tx.Commit()
	return utils.SuccessMess("成功", struct {
		Count int64        `json:"count" bson:"count"`
		Data  []model.Role `json:"data" bson:"data"`
	}{
		Count: count,
		Data:  roleDB,
	})
}
