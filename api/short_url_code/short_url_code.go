// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package short_url_code

import (
	"context"

	"compressURL/api/short_url_code/v1"
)

type IShortUrlCodeV1 interface {
	BatchCreate(ctx context.Context, req *v1.BatchCreateReq) (res *v1.BatchCreateRes, err error)
}
