// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"compressURL/api/common/v1"
)

type ICommonV1 interface {
	GetVerificationCodeEmail(ctx context.Context, req *v1.GetVerificationCodeEmailReq) (res *v1.GetVerificationCodeEmailRes, err error)
	UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error)
}
