package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete" summary:"删除用户" tags:"用户"`
	Id     string `json:"id"  v:"required#请输入用户 id"   dc:"用户 id"`
}
type DeleteRes struct {
}
