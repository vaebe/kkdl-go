// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"compressURL/internal/model"

	"golang.org/x/net/context"
)

type (
	IShortURL interface {
		// CreateShortURL 创建短链
		CreateShortURL(ctx context.Context, in model.ShortURLCreateInput) (string, error)
		// GetShortURL 获取短链
		GetShortURL(ctx context.Context, url string) (string, error)
	}
)

var (
	localShortURL IShortURL
)

func ShortURL() IShortURL {
	if localShortURL == nil {
		panic("implement not found for interface IShortURL, forgot register?")
	}
	return localShortURL
}

func RegisterShortURL(i IShortURL) {
	localShortURL = i
}
