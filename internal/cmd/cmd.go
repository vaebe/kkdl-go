package cmd

import (
	v1 "compressURL/api/shortURL/v1"
	"compressURL/internal/controller/shortURL"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
)

// 注册获取短链路由
func registerGetShortUrlRouter(s *ghttp.Server, ctx context.Context) {
	s.BindHandler("/:id", func(r *ghttp.Request) {
		shortUrl := gstr.SubStr(r.Request.RequestURI, 1, len(r.Request.RequestURI))

		req, err := shortURL.NewV1().GetUrl(ctx, &v1.GetUrlReq{ShortUrl: shortUrl})

		if err != nil {
			r.Response.WriteJson(&ghttp.DefaultHandlerResponse{
				Code:    -1,
				Message: "",
				Data:    nil,
			})
		}

		r.Response.WriteJson(&ghttp.DefaultHandlerResponse{
			Code:    0,
			Message: "",
			Data:    req.RawUrl,
		})
	})

}

func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()

	registerGetShortUrlRouter(s, ctx)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			shortURL.NewV1(),
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
