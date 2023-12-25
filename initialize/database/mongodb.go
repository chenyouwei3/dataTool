package database

import (
	"context"
	"dataTool/initialize/global"
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
		global.DeviceColl = smartGraphiteHBClone.Collection("device")
		//速控云

	}
	sukonCloud := global.MongodbClient403.Database("sukouCloud")
	{
		//速控云
		global.ImmersionHisData = sukonCloud.Collection("ImmersionHisData")
		global.WestCraneCarHisData = sukonCloud.Collection("WestCraneCarHisData")
		global.GraphitingHisData = sukonCloud.Collection("GraphitingHisData")
		global.TunnelWetElectricHisDataColl = sukonCloud.Collection("TunnelWetElectricHisDataColl")
		global.RoastWetElectricHisDataColl = sukonCloud.Collection("RoastWetElectricHisDataColl")
		global.GraphitingWetElectricHisDataColl = sukonCloud.Collection("GraphitingWetElectricHisDataColl")
		global.EastCraneCarHisData = sukonCloud.Collection("EastCraneCarHisData")
		global.TunnelHisDataColl = sukonCloud.Collection("TunnelHisDataColl")
		global.CrucibleHisDataColl = sukonCloud.Collection("CrucibleHisDataColl")
		global.CalcinationHisDataColl = sukonCloud.Collection("CalcinationHisDataColl")
		global.FormPlcHisDataColl = sukonCloud.Collection("FormPlcHisDataColl")
		global.RoastDenitrificationHisColl = sukonCloud.Collection("RoastDenitrificationHisColl *mongo.Collection")
		global.FourSeaStoreFormHisColl = sukonCloud.Collection("FourSeaStoreFormHisColl")
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
