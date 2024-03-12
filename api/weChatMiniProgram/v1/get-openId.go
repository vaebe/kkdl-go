package v1

import "github.com/gogf/gf/v2/frame/g"

type GetOpenIdReq struct {
	g.Meta   `path:"/weChatMiniProgram/getOpenId" method:"get" summary:"获取小程序用户openId" tags:"微信小程序"`
	Code     string `json:"code" v:"required#请输入code"    dc:"登录时获取的 code"`
	UserCode string `json:"user_code" v:"required#请输入用户 code"    dc:"登录时生成的用户 code"`
}

type GetOpenIdRes struct {
	SessionKey string `json:"session_key" dc:"会话密钥"`
	Unionid    string `json:"unionid" dc:"用户在开放平台的唯一标识符"`
	Openid     string `json:"openid" dc:"用户唯一标识"`
	ErrCode    int64  `json:"errCode"   dc:"错误码"`
	ErrMsg     string `json:"errMsg" dc:"错误信息"`
}
