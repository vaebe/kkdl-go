package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"math/rand"
	"time"
)

// GenerateRandomNumber 生成指定范围随机数
func GenerateRandomNumber(max int, min int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}
