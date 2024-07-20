package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GithubLoginReq github 登录授权
type GithubLoginReq struct {
	g.Meta `path:"/login/githubLogin" method:"post" summary:"github 登录授权" tags:"登录"`
	Code   string `json:"code"    dc:"授权 code"`
}

type GithubLoginRes EmailLoginRes
