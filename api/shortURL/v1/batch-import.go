package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// BatchImportReq 批量导入
type BatchImportReq struct {
	g.Meta `path:"/shortURL/batchImport" method:"post" summary:"批量导入短链" tags:"短链"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"请选择文件"`
}

type BatchImportRes struct {
}
