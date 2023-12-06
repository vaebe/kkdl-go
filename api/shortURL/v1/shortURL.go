package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta         `path:"/shortURL/create" method:"post" summary:"创建短链" tags:"短链"`
	RawUrl         string `json:"rawUrl" v:"required#请输入url"   dc:"rawUrl"`
	ExpirationTime string `json:"ExpirationTime"    dc:"过期时间"`
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
