package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta   `path:"/user/create" method:"post" summary:"创建用户" tags:"用户"`
	Email    string `json:"email"    dc:"邮箱"`
	WxId     string `json:"wxId"    dc:"小程序 id"`
	Password string `json:"password"   dc:"密码"`
	Nickname string `json:"nickname"   dc:"昵称"`
}
type CreateRes struct {
	Id string `json:"id" dc:"用户 id"`
}
