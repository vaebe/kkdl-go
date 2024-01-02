package cmd

import (
	v1 "compressURL/api/shortURL/v1"
	"compressURL/internal/controller/login"
	"compressURL/internal/controller/shortURL"
	"compressURL/internal/controller/user"
	"compressURL/internal/middlewares"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

// 注册获取短链路由
func registerGetShortUrlRouter(s *ghttp.Server, ctx context.Context) {
	s.BindHandler("/:id", func(r *ghttp.Request) {
		shortUrl := gstr.SubStr(r.Request.RequestURI, 1, len(r.Request.RequestURI))

		if shortUrl == "favicon.ico" {
			return
		}

		req, err := shortURL.NewV1().GetUrl(ctx, &v1.GetUrlReq{ShortUrl: shortUrl})
		if err != nil {
			r.Response.Write("未获取到对应的地址，请检查链接是否正确！")
			return
		}

		http.Redirect(r.Response.ResponseWriter, r.Request, req.RawUrl, http.StatusFound)
	})

}

func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {

	s := g.Server()

	err = GenerateShortLinkScheduledTask(ctx)

	if err != nil {
		return err
	}

	registerGetShortUrlRouter(s, ctx)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)

		// 不需要权限
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				login.NewV1().EmailLogin,
				login.NewV1().WeChatMiniProgramLogin,
				login.NewV1().GetMiniProgramCode,
				user.NewV1().Create,
				shortURL.NewV1().Create,
			)
		})

		// 需要权限验证
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middlewares.MiddlewareAuth)
			group.Bind(
				shortURL.NewV1().GetUrl,
				user.NewV1().GetUserInfo,
				user.NewV1().Remove,
				user.NewV1().Update,
				login.NewV1().SignOut,
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
