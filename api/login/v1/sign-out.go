package v1

import "github.com/gogf/gf/v2/frame/g"

// SignOutReq 退出登录 req
type SignOutReq struct {
	g.Meta `path:"/login/signOut" method:"get" summary:"退出登录" tags:"登录"`
}
type SignOutRes struct {
}
