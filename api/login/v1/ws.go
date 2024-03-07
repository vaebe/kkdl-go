package v1

import "github.com/gogf/gf/v2/frame/g"

type WsReq struct {
	g.Meta   `path:"/login/ws" method:"get" summary:"webSocket" tags:"登录"`
	UserCode string `json:"userCode"    dc:"用户code"`
}

type WsRes struct {
}
