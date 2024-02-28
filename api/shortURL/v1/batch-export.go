package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BatchExportReq 批量导出
type BatchExportReq struct {
	g.Meta `path:"/common/batchExport" method:"post" summary:"批量导出" tags:"短链"`
	Title  string `json:"title"   dc:"短链标题"`
	RawUrl string `json:"rawUrl"   dc:"跳转链接"`
}

type BatchExportRes struct {
}
