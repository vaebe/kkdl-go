package middlewares

import (
	"compressURL/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Auth jwt 鉴权
func Auth(r *ghttp.Request) {
	service.Auth().AuthInstance().MiddlewareFunc()(r)
	r.Middleware.Next()
}
