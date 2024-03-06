package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GetVerificationCodeEmailReq 获取邮箱验证码
type GetVerificationCodeEmailReq struct {
	g.Meta `path:"/common/getVerificationCodeEmail" method:"get" summary:"获取邮箱验证码" tags:"公共"`
	Email  string `json:"email"    dc:"邮箱"`
}

type GetVerificationCodeEmailRes struct {
}
