package v1

import (
	"compressURL/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// VerificationCodeLoginReq 验证码登录 req
type VerificationCodeLoginReq struct {
	g.Meta `path:"/login/verificationCodeLogin" method:"post" summary:"验证码登录" tags:"登录"`
	Email  string `json:"email" v:"required#请输入邮箱"   dc:"邮箱"`
	Code   string `json:"code" v:"required#请输入验证码"   dc:"验证码"`
}

type VerificationCodeLoginRes model.LoginRes
