// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"golang.org/x/net/context"
)

type (
	IShortUrlCode interface {
		// BatchCreateCode 批量创建短链
		BatchCreateCode(ctx context.Context, num int) error
		// UnusedCodeCount 获取未使用的 code 数量
		UnusedCodeCount(ctx context.Context) (int, error)
	}
)

var (
	localShortUrlCode IShortUrlCode
)

func ShortUrlCode() IShortUrlCode {
	if localShortUrlCode == nil {
		panic("implement not found for interface IShortUrlCode, forgot register?")
	}
	return localShortUrlCode
}

func RegisterShortUrlCode(i IShortUrlCode) {
	localShortUrlCode = i
}
