package utils

import (
	"math/rand"
)

func MsgID() int64 {
	// 设置随机数种子，以确保每次运行生成不同的随机数序列
	intn := rand.Intn(500)
	// 生成 100 到 1000 之间的随机整数
	randomNumber := intn + 500 // 生成 0 到 900 的随机整数再加上 100

	return int64(randomNumber)
}
