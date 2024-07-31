package v1

import (
	"compressURL/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// EmailLoginReq 邮箱登录 req
type EmailLoginReq struct {
	g.Meta   `path:"/login/emailLogin" method:"post" summary:"邮箱登录" tags:"登录"`
	Email    string `json:"email"    dc:"邮箱"`
	Password string `json:"password"   dc:"密码"`
}

type EmailLoginRes model.LoginRes
