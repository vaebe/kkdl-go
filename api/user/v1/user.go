package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta   `path:"/user/create" method:"post" summary:"创建用户" tags:"用户"`
	Email    string `json:"email"    dc:"邮箱"`
	WxId     string `json:"wxId"    dc:"小程序id"`
	Password string `json:"password"   dc:"密码"`
	Nickname string `json:"nickname"   dc:"昵称"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"用户id"`
}

type RemoveReq struct {
	g.Meta `path:"/user/remove" method:"delete" summary:"删除用户" tags:"用户"`
	UserId int64 `json:"userId"  v:"required#请输入用户 id"   dc:"用户id"`
}
type RemoveRes struct {
}

type UpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" summary:"更新用户信息" tags:"用户"`
	UserId   int64  `json:"userId"  v:"required#请输入用户 id"   dc:"用户id"`
	Password string `json:"password"   dc:"密码"`
	Nickname string `json:"nickname"   dc:"昵称"`
}
type UpdateRes struct {
}

type GetUserInfoReq struct {
	g.Meta `path:"/user/getUserInfo" method:"get" summary:"获取用户信息" tags:"用户"`
	UserId int64 `json:"userId"  v:"required#请输入用户 id"   dc:"用户id"`
}
type GetUserInfoRes struct {
	Id          int64  `json:"id" dc:"用户id"`
	Email       string `json:"email" dc:"邮箱"`
	WxId        string `json:"wxId" dc:"小程序id"`
	Nickname    string `json:"nickname" dc:"昵称"`
	AccountType string `json:"accountType" dc:"账号类型"`
	Role        string `json:"role" dc:"角色"`
}
