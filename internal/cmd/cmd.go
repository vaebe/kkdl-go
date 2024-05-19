package cmd

import (
	"compressURL/internal/controller/common"
	"compressURL/internal/controller/login"
	"compressURL/internal/controller/short_url"
	"compressURL/internal/controller/short_url_code"
	"compressURL/internal/controller/short_url_visits"
	"compressURL/internal/controller/user"
	"compressURL/internal/controller/weChatMiniProgram"
	"compressURL/internal/middlewares"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
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

		// 获取请求的IP地址
		g.Log().Info(ctx, "Client IP", r.GetClientIp())

		// 获取请求的User-Agent头部信息，通常包含浏览器信息
		userAgent := r.Header.Get("User-Agent")
		g.Log().Info(ctx, "User-Agent", userAgent)

		if url == "favicon.ico" {
			return
		}

		req, err := service.ShortUrl().GenOne(ctx, url)
		if err != nil {
			r.Response.Write("未获取到对应的地址，请检查链接是否正确！")
			return
		}

		err = service.ShortUrlVisits().Create(ctx, entity.ShortUrlVisits{
			UserId:          req.UserId,
			ShortUrl:        req.ShortUrl,
			RawUrl:          req.RawUrl,
			Ip:              r.GetClientIp(),
			UserAgent:       r.Header.Get("User-Agent"),
			SecChUa:         r.Header.Get("Sec-Ch-Ua"),
			SecChUaMobile:   r.Header.Get("Sec-Ch-Ua-Mobile"),
			SecChUaPlatform: r.Header.Get("Sec-Ch-Ua-Platform"),
			SecFetchUser:    r.Header.Get("Sec-Fetch-User"),
		})

		if err != nil {
			g.Log().Error(ctx, "创建短链访问信息失败:", err)
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
				login.NewV1().Ws,
				common.NewV1().GetCaptcha,
				weChatMiniProgram.NewV1(),
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
