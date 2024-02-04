package sukonCloud

import (
	"context"
	"dataTool/initialize/global"
	"dataTool/internal/model"
	"dataTool/pkg/utils"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func SuKonCloudProjects() { //获取项目
	URL := "http://sukon-cloud.com/api/v1/base/projects"
	urlValues := url.Values{}
	urlValues.Add("token", global.SukonCloudToken)
	var data model.SuKonProject
	res, err := http.PostForm(URL, urlValues)
	if err != nil {
		log.Println("请求错误:", err)
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("获取速控云项目失败:", err)
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("响应错误:", err)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("解析错误:", err)
		return
	}
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
	res, err := http.PostForm(URL, urlValues)
	if err != nil {
		log.Println("请求错误:", err)
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("获取box错误:", err)
		}
	}()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("响应错误:", err)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("解析错误:", err)
		return
	}
	if data.Success == false {
		log.Println("获取box异常", data)
	}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, box := range data.Data {
		fmt.Println(box.Name)
		mutex.Lock()
		switch box.Status {
		case "0":
			updateTime := utils.TimeFormat(time.Now())
			update := bson.M{"$set": bson.M{"status": "离线", "updateTime": updateTime}}
			err = global.DeviceColl.FindOneAndUpdate(context.TODO(), bson.M{"code": box.BoxId}, update).Decode(bson.M{})
			if err != nil && err != mongo.ErrNoDocuments {
				log.Println("0:", err)
				mutex.Unlock()
				continue
			}
			fmt.Println(box.Name + "设备离线")
			mutex.Unlock()
			continue
		case "1":
			updateTime := utils.TimeFormat(time.Now())
			update := bson.M{"$set": bson.M{"status": "正常", "updateTime": updateTime}}
			err = global.DeviceColl.FindOneAndUpdate(context.TODO(), bson.M{"code": box.BoxId}, update).Decode(bson.M{})
			if err != nil && err != mongo.ErrNoDocuments {
				log.Println("1:", err)
				mutex.Unlock()
				continue
			}
			wg.Add(1)
			go func(boxId string) {
				defer wg.Done()
				BoxPlc(boxId)
			}(box.BoxId)
			//BoxPlc(box.BoxId)
			fmt.Println(box.Name + "设备在线")
			mutex.Unlock()
		default:
			fmt.Println("没有设备")
		}
	}
	wg.Wait()
	fmt.Println("等待组是否结束")
}
