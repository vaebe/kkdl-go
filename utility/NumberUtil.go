package utility

import (
	"math/rand"
	"time"
)

// GenerateRandomNumber 生成指定范围随机数
func GenerateRandomNumber(max int, min int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
