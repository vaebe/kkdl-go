// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package shortUrlCode

import (
	"context"

	"compressURL/api/shortUrlCode/v1"
)

type IShortUrlCodeV1 interface {
	BatchCreate(ctx context.Context, req *v1.BatchCreateReq) (res *v1.BatchCreateRes, err error)
}
