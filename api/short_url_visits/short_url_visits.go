// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package short_url_visits

import (
	"context"

	"compressURL/api/short_url_visits/v1"
)

type IShortUrlVisitsV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
