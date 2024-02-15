package sukonCloud

import (
	"context"
	"dataTool/initialize/global"
	"dataTool/internal/model"
	"dataTool/pkg/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/url"
	"sync"
	"time"
)

func SuKonCloudProjects() { //获取项目
	URL := "http://sukon-cloud.com/api/v1/base/projects"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	var data model.SuKonProject
	data, _ = utils.Test(data, URL, urlValues)
	if data.Data == nil && data.Msg == "token已过期" {
		fmt.Println("project数据为空,获取失败---", data)
		utils.SukonToken()
		return
	}
	for _, project := range data.Data {
		if project.Id == "rKWw9LNBQYH" { //瑞通碳素项目id
			continue
		}
		suKonCloudBox(project.Id)
	}
}

func suKonCloudBox(projectId string) { //获取box并且更新box状态
	URL := "http://sukon-cloud.com/api/v1/base/projectBoxes"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	urlValues.Add("projectId", projectId)
	var data model.ProjectBox
	data, _ = utils.Test(data, URL, urlValues)
	if data.Success == false {
		log.Println("获取box异常", data)
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i, box := range data.Data {
		mutex.Lock()
		switch box.Status {
		case "0":
			update := bson.M{"$set": bson.M{"status": "离线", "updateTime": utils.TimeFormat(time.Now())}}
			err := global.DeviceColl.FindOneAndUpdate(context.TODO(), bson.M{"code": box.BoxId}, update).Decode(bson.M{})
			if err != nil && err != mongo.ErrNoDocuments {
				log.Println("0:", err)
				mutex.Unlock()
				continue
			}
			fmt.Println(box.Name + "设备离线")
			mutex.Unlock()
			continue
		case "1":
			update := bson.M{"$set": bson.M{"status": "正常", "updateTime": utils.TimeFormat(time.Now())}}
			err := global.DeviceColl.FindOneAndUpdate(context.TODO(), bson.M{"code": box.BoxId}, update).Decode(bson.M{})
			if err != nil && err != mongo.ErrNoDocuments {
				log.Println("1:", err)
				mutex.Unlock()
				continue
			}
			wg.Add(1)
			go func(boxId string, i int) {
				defer wg.Done()
				BoxPlc(boxId)
			}(box.BoxId, i)
			fmt.Println(box.Name + "设备在线")
			mutex.Unlock()
		default:
			fmt.Println("没有设备")
		}
	}
	wg.Wait()
}
