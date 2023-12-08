package middlewares

import (
	"compressURL/internal/logic/login"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddlewareAuth jwt 鉴权
func MiddlewareAuth(r *ghttp.Request) {
	login.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}
