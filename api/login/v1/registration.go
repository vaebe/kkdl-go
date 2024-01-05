package v1

import "github.com/gogf/gf/v2/frame/g"

type RegistrationReq struct {
	g.Meta   `path:"/user/registration" method:"post" summary:"用户注册" tags:"登录"`
	Email    string `json:"email"    dc:"邮箱"`
	WxId     string `json:"wxId"    dc:"小程序 id"`
	Password string `json:"password"   dc:"密码"`
	Nickname string `json:"nickname"   dc:"昵称"`
}
type RegistrationRes struct {
	Id string `json:"id" dc:"用户 id"`
}
