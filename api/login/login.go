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
	GetMiniProgramCode(ctx context.Context, req *v1.GetMiniProgramCodeReq) (res *v1.GetMiniProgramCodeRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	Registration(ctx context.Context, req *v1.RegistrationReq) (res *v1.RegistrationRes, err error)
	SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
	WeChatMiniProgramLogin(ctx context.Context, req *v1.WeChatMiniProgramLoginReq) (res *v1.WeChatMiniProgramLoginRes, err error)
}
