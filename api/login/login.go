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
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	Registration(ctx context.Context, req *v1.RegistrationReq) (res *v1.RegistrationRes, err error)
	SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
	WxMiniProgramLogin(ctx context.Context, req *v1.WxMiniProgramLoginReq) (res *v1.WxMiniProgramLoginRes, err error)
	Ws(ctx context.Context, req *v1.WsReq) (res *v1.WsRes, err error)
}
