package shortUrl

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"compressURL/api/shortUrl/v1"
)

func (c *ControllerV1) TemplateDownload(ctx context.Context, _ *v1.TemplateDownloadReq) (res *v1.TemplateDownloadRes, err error) {
	g.RequestFromCtx(ctx).Response.ServeFileDownload("resource/template/shortUrl/template.xlsx", "批量生成短链模版.xlsx")
	return
}
