// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package shortURL

import (
	"context"

	"compressURL/api/shortURL/v1"
)

type IShortURLV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetUrl(ctx context.Context, req *v1.GetUrlReq) (res *v1.GetUrlRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
