// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/short_url_visits/v1"
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	IShortUrlVisits interface {
		Create(ctx context.Context, in entity.ShortUrlVisits) error
		// GetList 短链列表
		GetList(ctx context.Context, in v1.GetListReq, userId string) ([]entity.ShortUrlVisits, int, error)
	}
)

var (
	localShortUrlVisits IShortUrlVisits
)

func ShortUrlVisits() IShortUrlVisits {
	if localShortUrlVisits == nil {
		panic("implement not found for interface IShortUrlVisits, forgot register?")
	}
	return localShortUrlVisits
}

func RegisterShortUrlVisits(i IShortUrlVisits) {
	localShortUrlVisits = i
}
