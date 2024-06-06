package utils

import (
	"math/rand"
	"time"
)

func Shuffle(arr []any) {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)           // 生成0到i之间的随机数
		arr[i], arr[j] = arr[j], arr[i] // 交换位置
	}
}
