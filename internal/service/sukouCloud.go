package service

import (
	"dataTool/initialize/global"
	"dataTool/internal/model"
	"dataTool/pkg/utils"
)

func StoreDippingData(box model.Box, data model.RealtimeData) { // 存储浸渍数据
	utils.SKCloudHisData(box, data, global.ImmersionHisData, "2da580adb26b4a12accd4aec80e04656:")
	return
}

func StoreWestAirCarData(box model.Box, data model.RealtimeData) { // 存储西跨吸料天车数据
	utils.SKCloudHisData(box, data, global.WestCraneCarHisData, "b46a0faf11cc4000a4c290eba5cc949a:")
	return
}

func StoreGraphitingData(box model.Box, data model.RealtimeData) { //存储石墨化数据
	utils.SKCloudHisData(box, data, global.GraphitingHisData, "be67c2b8216e49e8981a95663413f115:")
	return
}

func StoreTunnelWetElectricData(box model.Box, data model.RealtimeData) { // 存储隧道窑湿电数据
	utils.SKCloudHisData(box, data, global.TunnelWetElectricHisDataColl, "5cba298477bc456ab1a2bd06e35cb0d8:")
	return
}

func StoreRoastingWetElectricData(box model.Box, data model.RealtimeData) { // 存储焙烧湿电数据
	utils.SKCloudHisData(box, data, global.RoastWetElectricHisDataColl, "01e844f884844aa2bb5d1cab87316c17:")
	return
}

func StoreGraphiteWetElectricData(box model.Box, data model.RealtimeData) { // 存储石墨化湿电数据
	utils.SKCloudHisData(box, data, global.GraphitingWetElectricHisDataColl, "69fb82a9cba744188cab9da766787f25:")
	return
}

func StoreEarthAirCarData(box model.Box, data model.RealtimeData) { // 存储东跨跨吸料天车数据
	utils.SKCloudHisData(box, data, global.EastCraneCarHisData, "f73fe0d8688046e088bb073849aa0c3f:")
	return
}

func StoreTunnelData(box model.Box, data model.RealtimeData) { // 存储隧道窑数据
	utils.SKCloudHisData(box, data, global.TunnelHisDataColl, "9f62bc0edbd542b2bec159ac8f023509:")
	return
}

func StoreCrucibleData(box model.Box, data model.RealtimeData) { // 存储坩埚数据
	utils.SKCloudHisData(box, data, global.CrucibleHisDataColl, "9bd62f734af94dc0b0641817ac2807e9:")
	return
}

func StoreCalcinationData(box model.Box, data model.RealtimeData) { //存储煅烧脱销
	utils.SKCloudHisData(box, data, global.CalcinationHisDataColl, "ef62aa2e44204b5d82463b72a86f9621:")
}

func StoreFormPlcData(box model.Box, data model.RealtimeData) { //存储压型
	utils.SKCloudHisData(box, data, global.FormPlcHisDataColl, "65d27a491d744a0e91b4d8e6db628887:")
}

func StoreRoastDenitrificationData(box model.Box, data model.RealtimeData) { //焙烧脱硝
	utils.SKCloudHisData(box, data, global.RoastDenitrificationHisColl, "52980204e2dc4ce9907196441c6f9a32:")
}

func FourSeaStoreFormPlcData(box model.Box, data model.RealtimeData) {
	utils.SKCloudHisData(box, data, global.FourSeaStoreFormHisColl, "97509d2212bf4b1cb5bc3ea8dd8649d7:")
}
