// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/shortUrl/v1"
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	IShortUrl interface {
		// BatchImport 批量导入
		BatchImport(ctx context.Context, in []entity.ShortUrl) ([]string, error)
		// CreateShortUrl 创建短链
		CreateShortUrl(ctx context.Context, in entity.ShortUrl) (string, error)
		// GetShortUrl 获取短链
		GetShortUrl(ctx context.Context, url string) (string, error)
		// GetList 短链列表
		GetList(ctx context.Context, in v1.GetListReq, userId string) ([]entity.ShortUrl, int, error)
		// Delete 删除短链
		Delete(ctx context.Context, id string) error
	}
)

var (
	localShortUrl IShortUrl
)

func ShortUrl() IShortUrl {
	if localShortUrl == nil {
		panic("implement not found for interface IShortUrl, forgot register?")
	}
	return localShortUrl
}

func RegisterShortUrl(i IShortUrl) {
	localShortUrl = i
}
