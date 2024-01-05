package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" summary:"更新用户信息" tags:"用户"`
	Id       string `json:"id"  v:"required#请输入用户 id"   dc:"用户 id"`
	Email    string `json:"email"    dc:"邮箱"`
	Password string `json:"password"   dc:"密码"`
	NickName string `json:"nickName"   dc:"昵称"`
	Role     string `json:"role"   dc:"角色"`
	Avatar   string `json:"avatar"   dc:"头像"`
}
type UpdateRes struct {
}
