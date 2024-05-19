package cmd

import (
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

func saveVisitsInfo(ctx context.Context, r *ghttp.Request, shortUrlInfo entity.ShortUrl) {
	err := service.ShortUrlVisits().Create(ctx, entity.ShortUrlVisits{
		UserId:          shortUrlInfo.UserId,
		ShortUrl:        shortUrlInfo.ShortUrl,
		RawUrl:          shortUrlInfo.RawUrl,
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
}

// 注册短链服务
func registerShortURLService(s *ghttp.Server, ctx context.Context) {
	s.BindHandler("/:id", func(r *ghttp.Request) {
		url := gstr.SubStr(r.Request.RequestURI, 1, len(r.Request.RequestURI))

		if url == "favicon.ico" {
			return
		}

		req, err := service.ShortUrl().GenOne(ctx, url)

		g.Log().Debug(ctx, url, err)

		if err != nil {
			r.Response.Write("未获取到对应的地址，请检查链接是否正确！")
			return
		}

		go saveVisitsInfo(ctx, r, req)

		http.Redirect(r.Response.ResponseWriter, r.Request, req.RawUrl, http.StatusFound)
	})
}
