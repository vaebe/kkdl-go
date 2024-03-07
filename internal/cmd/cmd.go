package cmd

import (
	v1 "compressURL/api/shortUrl/v1"
	"compressURL/internal/controller/common"
	"compressURL/internal/controller/login"
	"compressURL/internal/controller/shortUrl"
	"compressURL/internal/controller/shortUrlCode"
	"compressURL/internal/controller/user"
	"compressURL/internal/controller/weChatMiniProgram"
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
		url := gstr.SubStr(r.Request.RequestURI, 1, len(r.Request.RequestURI))

		if url == "favicon.ico" {
			return
		}

		req, err := shortUrl.NewV1().GetUrl(ctx, &v1.GetUrlReq{ShortUrl: url})
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
				login.NewV1().WxMiniProgramLogin,
				login.NewV1().Registration,
				shortUrl.NewV1().Create,
				common.NewV1().GetVerificationCodeEmail,
				weChatMiniProgram.NewV1(),
			)
		})

		// 需要权限验证
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middlewares.MiddlewareAuth)
			group.Bind(
				login.NewV1().SignOut,
				login.NewV1().RefreshToken,
				shortUrl.NewV1().GetUrl,
				shortUrl.NewV1().GetList,
				shortUrl.NewV1().Delete,
				shortUrl.NewV1().BatchImport,
				shortUrl.NewV1().TemplateDownload,
				shortUrl.NewV1().BatchExport,
				shortUrlCode.NewV1(),
				common.NewV1().UploadFile,
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
