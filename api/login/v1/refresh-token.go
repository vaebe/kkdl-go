package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RefreshTokenReq 刷新 token
type RefreshTokenReq struct {
	g.Meta `path:"/login/refreshToken" method:"post" summary:"刷新 token" tags:"登录"`
}

type RefreshTokenRes struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}
