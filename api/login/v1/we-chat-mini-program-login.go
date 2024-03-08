package v1

import (
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// WxMiniProgramLoginReq 微信小程序登录 req
type WxMiniProgramLoginReq struct {
	g.Meta   `path:"/login/wxMiniProgramLogin" method:"post" summary:"微信小程序登录" tags:"登录"`
	Code     string `json:"code" v:"required#请输入code"    dc:"登录时获取的 code"`
	UserCode string `json:"user_code" v:"required#请输入用户 code"    dc:"登录时生成的用户 code"`
}

type WxMiniProgramLoginRes struct {
	Token       string      `json:"token" dc:"jwt token"`
	TokenExpire string      `json:"tokenExpire" dc:"token 过期时间"`
	UserInfo    entity.User `json:"userInfo"   dc:"用户信息"`
}
