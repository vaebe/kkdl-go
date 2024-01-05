package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" summary:"更新用户信息" tags:"用户"`
	UserId   string `json:"userId"  v:"required#请输入用户 id"   dc:"用户 id"`
	Password string `json:"password"   dc:"密码"`
	Nickname string `json:"nickname"   dc:"昵称"`
}
type UpdateRes struct {
}
