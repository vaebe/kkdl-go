// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	IShortUrlVisits interface {
		Create(ctx context.Context, in entity.ShortUrlVisits) error
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
