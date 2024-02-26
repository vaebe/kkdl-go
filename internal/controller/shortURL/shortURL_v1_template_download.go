package shortURL

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"compressURL/api/shortURL/v1"
)

func (c *ControllerV1) TemplateDownload(ctx context.Context, _ *v1.TemplateDownloadReq) (res *v1.TemplateDownloadRes, err error) {
	g.RequestFromCtx(ctx).Response.ServeFileDownload("resource/template/shortURL/template.xlsx", "批量生成短链模版.xlsx")
	return
}
