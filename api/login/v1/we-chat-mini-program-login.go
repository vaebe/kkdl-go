package v1

import (
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// WeChatMiniProgramLoginReq 微信小程序登录 req
type WeChatMiniProgramLoginReq struct {
	g.Meta   `path:"/login/weChatMiniProgramLogin" method:"post" summary:"微信小程序登录" tags:"登录"`
	WxId     string `json:"wxId"    dc:"小程序id"`
	Password string `json:"password"   dc:"密码"`
}

type WeChatMiniProgramLoginRes struct {
	Token       string      `json:"token" dc:"jwt token"`
	TokenExpire string      `json:"tokenExpire" dc:"token 过期时间"`
	UserInfo    entity.User `json:"userInfo"   dc:"用户信息"`
}
