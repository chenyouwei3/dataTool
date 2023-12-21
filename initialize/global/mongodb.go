package global

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongodbClient403 *mongo.Client
	DeviceColl       *mongo.Collection
	//浸渍
	ImmersionHisData *mongo.Collection //浸渍
	//西跨吸料天车
	WestCraneCarHisData *mongo.Collection
	//隧道窑湿电
	GraphitingHisData *mongo.Collection
	//石墨化
	TunnelWetElectricHisDataColl *mongo.Collection
	//焙烧湿电
	RoastWetElectricHisDataColl *mongo.Collection
	//石墨化湿电
	GraphitingWetElectricHisDataColl *mongo.Collection
	//东跨跨吸料天车
	EastCraneCarHisData *mongo.Collection
	//隧道窑
	TunnelHisDataColl *mongo.Collection
	//坩埚
	CrucibleHisDataColl *mongo.Collection
	//煅烧脱销
	CalcinationHisDataColl *mongo.Collection
	//压型
	FormPlcHisDataColl *mongo.Collection
	//存储焙烧脱硝
	RoastDenitrificationHisColl *mongo.Collection
	//四海成型plc
	FourSeaStoreFormHisColl *mongo.Collection
)
