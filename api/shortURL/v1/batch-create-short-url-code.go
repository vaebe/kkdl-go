package v1

import "github.com/gogf/gf/v2/frame/g"

type BatchCreateShortUrlCodeReq struct {
	g.Meta `path:"/shortURL/batchCreateShortUrlCode" method:"get" summary:"批量生成短链" tags:"短链"`
	Num    int `json:"num" v:"required#请输入生成数量"   dc:"num"`
}
type BatchCreateShortUrlCodeRes struct {
}
