package database

import (
	"context"
	"dataTool/config/global"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func MongodbInit() {
	if global.MongodbClient403 == nil {
		global.MongodbClient403 = getMongodbClient(global.MongodbAddress)
	}
	smartGraphiteHBClone := global.MongodbClient403.Database("smartGraphiteHB-Clone")
	{
		global.Device = smartGraphiteHBClone.Collection("device")
	}
}

func getMongodbClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)                    // 创建一个新的客户端选项实例
	mongodbClient, err := mongo.Connect(context.TODO(), clientOptions) //使用提供的客户端选项连接到 mongodb
	if err != nil {
		log.Fatalln("客户端连接mongodb失败:", err)
	}
	err = mongodbClient.Ping(context.TODO(), nil) // 验证与mongodb的连接是否成功
	if err != nil {
		log.Fatalln("客户端连接mongodb失败:", err)
	}
	fmt.Println("mongodb连接成功")
	return mongodbClient
}
