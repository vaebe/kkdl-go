package cmd

import (
	v1 "compressURL/api/shortURL/v1"
	"compressURL/internal/controller/shortURL"
	"compressURL/internal/controller/user"
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
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			shortURL.NewV1(),
			user.NewV1(),
		)
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
