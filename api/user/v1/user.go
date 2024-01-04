package v1

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

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

type RemoveReq struct {
	g.Meta `path:"/user/remove" method:"delete" summary:"删除用户" tags:"用户"`
	UserId string `json:"userId"  v:"required#请输入用户 id"   dc:"用户 id"`
}
type RemoveRes struct {
}

type UpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" summary:"更新用户信息" tags:"用户"`
	UserId   string `json:"userId"  v:"required#请输入用户 id"   dc:"用户 id"`
	Password string `json:"password"   dc:"密码"`
	Nickname string `json:"nickname"   dc:"昵称"`
}
type UpdateRes struct {
}

type GetUserInfoReq struct {
	g.Meta `path:"/user/getUserInfo" method:"get" summary:"获取用户信息" tags:"用户"`
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

type GetUserListReq struct {
	g.Meta `path:"/user/getUserList" method:"post" summary:"获取用户列表" tags:"用户"`
	model.PageParams
	Nickname string `json:"nickname" dc:"昵称"`
	Email    string `json:"email"   dc:"邮箱"`
	WxId     string `json:"wxId"   dc:"微信 id"`
}
type GetUserListRes struct {
	List []entity.User `json:"list" dc:"用户数据"`
	model.PageParams
}
