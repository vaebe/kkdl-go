package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta         `path:"/shortURL/create" method:"post" summary:"创建短链" tags:"短链"`
	Title          string `json:"title" v:"required#请输入短链标题"   dc:"短链标题"`
	RawUrl         string `json:"rawUrl" v:"required#请输入跳转链接"   dc:"跳转链接"`
	ExpirationTime string `json:"expirationTime" v:"required#请输入过期时间"   dc:"过期时间"`
	GroupId        int    `json:"group_id"    dc:"分组id"`
}
type CreateRes struct {
	ShortUrl string `json:"shortUrl" dc:"短链"`
}

type GetUrlReq struct {
	g.Meta   `path:"/shortURL/GetUrl" method:"get" summary:"根据短链获取 url" tags:"短链"`
	ShortUrl string `json:"shortUrl" v:"required#请输入 url"   dc:"shortUrl"`
}
type GetUrlRes struct {
	RawUrl string `json:"rawUrl" dc:"原始 url"`
}
