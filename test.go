package main

import (
	"context"
	"dataTool/initialize"
	"dataTool/initialize/global"
	"dataTool/internal/model"
	"dataTool/pkg/redis"
	"dataTool/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func init() {
	initialize.InitConfig()
}

func main() {
	code := "2da580adb26b4a12accd4aec80e04656"
	var db model.Device
	update := bson.M{"$set": bson.M{"status": "离线", "updateTime": utils.TimeFormat(time.Now())}}
	err := global.DeviceColl.FindOneAndUpdate(context.TODO(), bson.M{"code": code}, update).Decode(&db)
	if err != nil && err != mongo.ErrNoDocuments {
		logrus.Error("0:", err)
	}
	fmt.Println("寻找到的设备", db)

	jsonDB, err := json.Marshal(db)
	if err != nil {
		fmt.Println("转化失败:", err)
	}
	fmt.Println("转化成功:", jsonDB)
	err = redis.Redis{}.SetValueHash("sh", code, string(jsonDB))
	if err != nil {
		fmt.Println("redis:Redis", err)
	}

	JJ, _ := redis.Redis{}.GetValueHash("sh", code)
	fmt.Println("读取成功:", JJ)
	//err := redis.Redis{}.DeleteValueHash("sh", code)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
