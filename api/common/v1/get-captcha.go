package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetCaptchaReq 获取验证码
type GetCaptchaReq struct {
	g.Meta `path:"/common/getCaptcha" method:"get" summary:"获取验证码" tags:"公共"`
	Email  string `json:"email" v:"required|email#请输入邮箱|邮箱格式不正确"    dc:"邮箱"`
}

type GetCaptchaRes struct {
}
