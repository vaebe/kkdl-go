package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta         `path:"/shortUrl/create" method:"post" summary:"创建短链" tags:"短链"`
	Title          string `json:"title" v:"required#请输入短链标题"   dc:"短链标题"`
	RawUrl         string `json:"rawUrl" v:"required#请输入跳转链接"   dc:"跳转链接"`
	ExpirationTime string `json:"expirationTime" v:"required#请输入过期时间"   dc:"过期时间"`
	GroupId        int    `json:"group_id"    dc:"分组id"`
}
type CreateRes struct {
	ShortUrl string `json:"shortUrl" dc:"短链"`
}
