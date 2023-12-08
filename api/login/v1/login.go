package v1

import "github.com/gogf/gf/v2/frame/g"

// EmailLoginReq 邮箱登录 req
type EmailLoginReq struct {
	g.Meta   `path:"/login/emailLogin" method:"post" summary:"邮箱登录" tags:"登录"`
	Email    string `json:"email"    dc:"邮箱"`
	Password string `json:"password"   dc:"密码"`
}

type UserInfo struct {
	Id          int64  `json:"id" dc:"用户id"`
	Email       string `json:"email" dc:"邮箱"`
	WxId        string `json:"wxId" dc:"小程序id"`
	Nickname    string `json:"nickname" dc:"昵称"`
	AccountType string `json:"accountType" dc:"账号类型"`
	Role        string `json:"role" dc:"角色"`
}

type EmailLoginRes struct {
	Token       string   `json:"token" dc:"jwt token"`
	TokenExpire string   `json:"tokenExpire" dc:"token 过期时间"`
	UserInfo    UserInfo `json:"userInfo"   dc:"用户信息"`
}

// WeChatMiniProgramLoginReq 微信小程序登录 req
type WeChatMiniProgramLoginReq struct {
	g.Meta   `path:"/login/weChatMiniProgramLogin" method:"post" summary:"微信小程序登录" tags:"登录"`
	WxId     string `json:"wxId"    dc:"小程序id"`
	Password string `json:"password"   dc:"密码"`
}

type WeChatMiniProgramLoginRes struct {
	Token       string   `json:"token" dc:"jwt token"`
	TokenExpire string   `json:"tokenExpire" dc:"token 过期时间"`
	UserInfo    UserInfo `json:"userInfo"   dc:"用户信息"`
}

// SignOutReq 退出登录 req
type SignOutReq struct {
	g.Meta `path:"/login/signOut" method:"get" summary:"退出登录" tags:"登录"`
}
type SignOutRes struct {
}
