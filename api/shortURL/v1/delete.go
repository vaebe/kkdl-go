package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/shortURL/delete" method:"delete" summary:"根据 id 删除对应的短链" tags:"短链"`
	Id     string `json:"id" v:"required#请输入 id"   dc:"短链 id"`
}
type DeleteRes struct {
}
