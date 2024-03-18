package main

import (
	"context"
	"dataTool/initialize"
	"dataTool/initialize/global"
	"dataTool/internal/model"
	"dataTool/pkg/redis"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	initialize.InitConfig()
}

func main() {
	//code := "2da580adb26b4a12accd4aec80e04656"
	//firstDB, _ := redis.Redis{}.GetValueHash("sh", code) //读取设备
	//var DB model.Device
	//err := json.Unmarshal([]byte(firstDB), DB) //解码
	//if err != nil {
	//	fmt.Println("解码失败:", err)
	//}
	//DB.Status = "离线"                             //修改状态
	//DB.UpdateTime = utils.TimeFormat(time.Now()) //修改状态
	//lastDB, err := json.Marshal(DB)
	//if err != nil {
	//	fmt.Println("转化失败:", err)
	//}
	//err = redis.Redis{}.SetValueHash("sh", code, string(lastDB))

	var device []model.Device
	cur, _ := global.DeviceColl.Find(context.TODO(), bson.M{})
	cur.All(context.TODO(), &device)
	for _, v := range device {
		lastDB, _ := json.Marshal(v)
		redis.Redis{}.SetValueHash("rtts", v.Code, string(lastDB))
	}
	for _, v := range device {
		fmt.Println(v.Code)
		str, _ := redis.Redis{}.GetValueHash("rtts", v.Code)
		fmt.Println("*--------------------->", str)
		//err := redis.Redis{}.DeleteValueHash("rtts", v.Code)
		//if err != nil {
		//	fmt.Println(err)
		//}
	}
}
