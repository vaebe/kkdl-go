// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/shortURL/v1"
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	IShortURL interface {
		// BatchCreateShortUrlCode 批量生成短链 code
		BatchCreateShortUrlCode(ctx context.Context, num int) error
		// BatchImport 批量导入
		BatchImport(ctx context.Context, in []entity.ShortUrl) ([]string, error)
		// CreateShortURL 创建短链
		CreateShortURL(ctx context.Context, in entity.ShortUrl) (string, error)
		// GetShortURL 获取短链
		GetShortURL(ctx context.Context, url string) (string, error)
		// GetList 短链列表
		GetList(ctx context.Context, in v1.GetListReq, userId string) ([]entity.ShortUrl, int, error)
		// Delete 删除短链
		Delete(ctx context.Context, id string) error
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
