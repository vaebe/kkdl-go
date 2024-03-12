// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package weChatMiniProgram

import (
	"context"

	"compressURL/api/weChatMiniProgram/v1"
)

type IWeChatMiniProgramV1 interface {
	GetMiniProgramCode(ctx context.Context, req *v1.GetMiniProgramCodeReq) (res *v1.GetMiniProgramCodeRes, err error)
	GetOpenId(ctx context.Context, req *v1.GetOpenIdReq) (res *v1.GetOpenIdRes, err error)
}
