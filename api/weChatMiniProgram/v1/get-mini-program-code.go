package v1

import "github.com/gogf/gf/v2/frame/g"

type GetMiniProgramCodeReq struct {
	g.Meta `path:"/weChatMiniProgram/getMiniProgramCode" method:"get" summary:"获取小程序码" tags:"微信小程序"`
	Scene  string `json:"scene"    dc:"额外传递的信息"`
	Page   string `json:"page"    dc:"跳转小程序的页面"`
}

type GetMiniProgramCodeRes struct {
	ErrCode int64  `json:"errCode"   dc:"错误码"`
	ErrMsg  string `json:"errMsg" dc:"错误信息"`
}
