// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package login

import (
	"context"

	"compressURL/api/login/v1"
)

type ILoginV1 interface {
	EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error)
	WeChatMiniProgramLogin(ctx context.Context, req *v1.WeChatMiniProgramLoginReq) (res *v1.WeChatMiniProgramLoginRes, err error)
	SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
	GetMiniProgramCode(ctx context.Context, req *v1.GetMiniProgramCodeReq) (res *v1.GetMiniProgramCodeRes, err error)
}
