package middlewares

import (
	"compressURL/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

// MiddlewareAuth jwt 鉴权
func MiddlewareAuth(r *ghttp.Request) {
	service.Auth().AuthInstance().MiddlewareFunc()(r)
	r.Middleware.Next()
}
