package service

import (
	"dataTool/config/global"
	"dataTool/internal/model"
	"dataTool/pkg/utils"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

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
		time0 := time.Now()
		api.CreateTime = &time0
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
	fmt.Println(api)
	if api.Id == 0 || api.Name == "" || api.Url == "" || len(api.Url) >= 20 || len(api.Name) >= 10 || (api.Method != "GET" && api.Method != "POST" && api.Method != "PUT" && api.Method != "DELETE") {
		return utils.ErrorMess("失败,参数错误", nil)
	}
	var apiDB model.Api
	res := tx.Where("id=?", api.Id).First(&apiDB)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		tx.Rollback()
		return utils.ErrorMess("失败,该API不存在", nil)
	}
	apiDB = api
	time0 := time.Now()
	apiDB.UpdateTime = &time0
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
