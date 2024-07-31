package v1

import (
	"compressURL/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// WxMiniProgramLoginReq 微信小程序登录 req
type WxMiniProgramLoginReq struct {
	g.Meta   `path:"/login/wxMiniProgramLogin" method:"post" summary:"微信小程序登录" tags:"登录"`
	Code     string `json:"code" v:"required#请输入code"    dc:"登录时获取的 code"`
	UserCode string `json:"user_code" v:"required#请输入用户 code"    dc:"登录时生成的用户 code"`
}

type WxMiniProgramLoginRes model.LoginRes
