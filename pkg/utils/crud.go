package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

func GetPage(currPage, pageSize string) (int, int, error) {
	curr, err := strconv.Atoi(currPage)
	if err != nil {
		return 0, 0, err
	}
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		return 0, 0, err
	}
	skip := (curr - 1) * size
	return skip, size, nil
}

func GetTime(startTime, endTime string) (filter bson.M) {
	if startTime != "" && endTime != "" {
		filter = bson.M{
			"createTime": bson.M{
				"$gte": startTime,
				"$lte": endTime,
			},
		}
		return filter
	}
	filter = bson.M{}
	return filter
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetNowTime() *time.Time {
	time := time.Now()
	return &time
}
