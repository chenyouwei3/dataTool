package redis

import (
	"dataTool/initialize/global"
	"fmt"
)

func (r Redis) SetValueHash(key, field, value string) error {
	err := global.RedisClient.HSet(key, field, value).Err()
	if err != nil {
		return fmt.Errorf("redis(sds)设置失败:%w", err)
	}
	return nil
}

func (r Redis) GetValueHash(key, field string) (string, error) {
	value, err := global.RedisClient.HGet(key, field).Result()
	if err != nil {
		return "", fmt.Errorf("redis(sds)读取失败:%w", err)
	}
	return value, nil
}

func (r Redis) DeleteValueHash(key, field string) error {
	err := global.RedisClient.HDel(key, field).Err()
	if err != nil {
		return fmt.Errorf("redis(hash)删除失败:%w", err)
	}
	return nil
}

//func (redis Redis) GetAndUpdateSukouCloud(key, field string, value *[]model.Device) ([]model.Device, error) {
//	value, err := global.RedisClient.HGet(key, field).Result()
//	if err != nil {
//		return nil, fmt.Errorf("redis(sds)读取设备失败:%w", err)
//	}
//
//}
