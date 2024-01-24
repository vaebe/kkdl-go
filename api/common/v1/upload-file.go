package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadFileReq 文件上传 req
type UploadFileReq struct {
	g.Meta `path:"/common/uploadFile" method:"post" summary:"文件上传" tags:"公共"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type UploadFileRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"文件 url"`
}
