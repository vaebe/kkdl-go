package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserRegCheckReq 检查用户是否已经注册 req
type UserRegCheckReq struct {
	g.Meta `path:"/login/userRegCheck" method:"get" summary:"检查用户是否已经注册" tags:"登录"`
	Email  string `json:"email"    dc:"邮箱"`
}

type UserRegCheckRes struct {
	IsRegistered bool `json:"isRegistered" dc:"用户是否注册"`
}
