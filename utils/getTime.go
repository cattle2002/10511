package utils

import (
	"time"
)

func GetTime() string {
	currentTime := time.Now()

	// 格式化时间为年月日时分秒
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return formattedTime
}
func GetCurrentMillSecond(productTime int64) bool {
	ms := time.Now().Unix() * 1000
	if ms >= productTime {
		return true
	}
	return false
}
