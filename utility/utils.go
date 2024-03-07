package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"math/rand"
	"strings"
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

// GetFileSuffixName 获取文件后缀名称
func GetFileSuffixName(filename string) string {
	indexOfDot := strings.LastIndex(filename, ".") //获取文件后缀名前的.的位置
	if indexOfDot < 0 {
		return ""
	}
	suffix := filename[indexOfDot+1 : len(filename)] //获取后缀名
	return strings.ToLower(suffix)                   //后缀名统一小写处理
}

// SliceEqual 切片是否相等
func SliceEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}

	return true
}

// GetWeChatMiniProgramLoginCode 获取微信小程序登录用户 code key
func GetWeChatMiniProgramLoginCode(code string) string {
	return fmt.Sprintf("weChatMiniProgramLoginCode%s", code)
}
