package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUserInfoReq struct {
	g.Meta `path:"/user/detail" method:"get" summary:"获取用户信息" tags:"用户"`
	UserId string `json:"userId"  v:"required#请输入用户 id"   dc:"用户 id"`
}
type GetUserInfoRes struct {
	Id          string `json:"id" dc:"用户 id"`
	Email       string `json:"email" dc:"邮箱"`
	WxId        string `json:"wxId" dc:"小程序 id"`
	Nickname    string `json:"nickname" dc:"昵称"`
	AccountType string `json:"accountType" dc:"账号类型"`
	Role        string `json:"role" dc:"角色"`
}
