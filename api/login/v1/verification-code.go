package v1

import (
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// VerificationCodeReq 验证码登录 req
type VerificationCodeReq struct {
	g.Meta `path:"/login/verificationCode" method:"post" summary:"验证码登录" tags:"登录"`
	Email  string `json:"email"    dc:"邮箱"`
	Code   string `json:"code"   dc:"验证码"`
}

type VerificationCodeRes struct {
	Token       string      `json:"token" dc:"jwt token"`
	TokenExpire string      `json:"tokenExpire" dc:"token 过期时间"`
	UserInfo    entity.User `json:"userInfo"   dc:"用户信息"`
}
