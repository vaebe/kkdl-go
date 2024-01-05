package v1

import "github.com/gogf/gf/v2/frame/g"

type RemoveReq struct {
	g.Meta `path:"/user/remove" method:"delete" summary:"删除用户" tags:"用户"`
	UserId string `json:"userId"  v:"required#请输入用户 id"   dc:"用户 id"`
}
type RemoveRes struct {
}
