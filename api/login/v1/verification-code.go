package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// VerificationCodeReq 验证码登录 req
type VerificationCodeReq struct {
	g.Meta `path:"/login/verificationCode" method:"post" summary:"验证码登录" tags:"登录"`
	Email  string `json:"email" v:"required#请输入邮箱"   dc:"邮箱"`
	Code   string `json:"code" v:"required#请输入验证码"   dc:"验证码"`
}

type VerificationCodeRes EmailLoginRes
