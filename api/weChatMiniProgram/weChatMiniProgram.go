// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package weChatMiniProgram

import (
	"context"

	"compressURL/api/weChatMiniProgram/v1"
)

type IWeChatMiniProgramV1 interface {
	GetWeChatMiniProgramOpenId(ctx context.Context, req *v1.GetWeChatMiniProgramOpenIdReq) (res *v1.GetWeChatMiniProgramOpenIdRes, err error)
}
