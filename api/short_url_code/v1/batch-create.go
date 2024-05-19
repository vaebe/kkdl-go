package v1

import "github.com/gogf/gf/v2/frame/g"

type BatchCreateReq struct {
	g.Meta `path:"/shortUrlCode/batchCreate" method:"get" summary:"批量生成短链code" tags:"短链code"`
	Num    int `json:"num" v:"required#请输入生成数量"   dc:"num"`
}
type BatchCreateRes struct {
}
