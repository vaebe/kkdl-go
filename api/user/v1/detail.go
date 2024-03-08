package v1

import "github.com/gogf/gf/v2/frame/g"

type GetOneReq struct {
	g.Meta `path:"/user/detail" method:"get" summary:"获取用户信息详情" tags:"用户"`
	UserId string `json:"userId"  v:"required#请输入用户 id"   dc:"用户 id"`
}
type GetOneRes struct {
	Id          string `json:"id" dc:"用户 id"`
	Email       string `json:"email" dc:"邮箱"`
	WxId        string `json:"wxId" dc:"小程序 id"`
	NickName    string `json:"nickName" dc:"昵称"`
	AccountType string `json:"accountType" dc:"账号类型"`
	Role        string `json:"role" dc:"角色"`
	Avatar      string `json:"avatar" dc:"头像"`
}
