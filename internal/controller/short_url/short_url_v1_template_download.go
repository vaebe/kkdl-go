package short_url

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"compressURL/api/short_url/v1"
)

func (c *ControllerV1) TemplateDownload(ctx context.Context, _ *v1.TemplateDownloadReq) (res *v1.TemplateDownloadRes, err error) {
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return nil, gerror.New("获取请求对象失败")
	}

	r.Response.ServeFileDownload("resource/template/shortUrl/template.xlsx", "批量生成短链模版.xlsx")
	return
}
