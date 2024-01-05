package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUrlReq struct {
	g.Meta   `path:"/shortURL/GetUrl" method:"get" summary:"根据短链获取 url" tags:"短链"`
	ShortUrl string `json:"shortUrl" v:"required#请输入 url"   dc:"shortUrl"`
}
type GetUrlRes struct {
	RawUrl string `json:"rawUrl" dc:"原始 url"`
}
