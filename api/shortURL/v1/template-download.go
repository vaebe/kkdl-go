package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TemplateDownloadReq 导入模版下载
type TemplateDownloadReq struct {
	g.Meta `path:"/shortURL/templateDownload" method:"get" summary:"导入模版下载" tags:"短链"`
}

type TemplateDownloadRes struct {
}
