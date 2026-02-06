package cmd

import (
	"compressURL/internal/controller/analytics"
	"compressURL/internal/controller/common"
	"compressURL/internal/controller/login"
	"compressURL/internal/controller/short_url"
	"compressURL/internal/controller/short_url_code"
	"compressURL/internal/controller/short_url_visits"
	"compressURL/internal/controller/user"
	"compressURL/internal/middlewares"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func mainFunc(ctx context.Context, _ *gcmd.Parser) (err error) {

	s := g.Server()

	err = GenerateShortLinkScheduledTask(ctx)

	if err != nil {
		return err
	}

	registerShortURLService(s, ctx)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)

		// 不需要权限
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				login.NewV1().EmailLogin,
				login.NewV1().Registration,
				login.NewV1().Ws,
				login.NewV1().GithubLogin,
				login.NewV1().UserRegCheck,
				login.NewV1().VerificationCodeLogin,
				common.NewV1().GetCaptcha,
				analytics.NewV1(),
			)
		})

		// 需要权限验证
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middlewares.Auth)
			group.Bind(
				login.NewV1().SignOut,
				login.NewV1().RefreshToken,
				short_url_code.NewV1(),
				common.NewV1().UploadFile,
				short_url.NewV1(),
				short_url_visits.NewV1(),
			)
		})

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middlewares.Auth)
			group.Middleware(middlewares.UserIsAdmin)

			group.Bind(
				user.NewV1(),
			)
		})
	})
	s.Run()
	return nil
}

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "start http server",
	Func:  mainFunc,
}
