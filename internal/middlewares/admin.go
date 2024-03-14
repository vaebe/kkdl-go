package middlewares

import (
	"compressURL/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UserIsAdmin 是否是 admin 用户
func UserIsAdmin(r *ghttp.Request) {
	// 获取当前登陆用户信息
	loginUserInfo, err := service.Auth().GetLoginUserInfo(r.GetCtx())
	if err != nil {
		r.Response.WriteJsonExit(map[string]interface{}{
			"code":    500,
			"message": "用户未登录不可操作!",
		})
		return
	}

	// 非管理员不能创建用户
	if loginUserInfo.Role != "00" {
		r.Response.WriteJsonExit(map[string]interface{}{
			"code":    500,
			"message": "非管理员不能操作!",
		})

		return
	}

	r.Middleware.Next()
}
