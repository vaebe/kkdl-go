package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta   `path:"/user/create" method:"post" summary:"创建用户" tags:"用户"`
	Email    string `json:"email"    dc:"邮箱"`
	Password string `json:"password"   dc:"密码"`
	NickName string `json:"nickName"   dc:"昵称"`
	Role     string `json:"role"   dc:"角色"`
	Avatar   string `json:"avatar"   dc:"头像"`
}

type CreateRes struct {
	Id string `json:"id" dc:"用户 id"`
}
