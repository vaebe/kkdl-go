package common

import (
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/common/v1"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}

	fileInfo, err := service.Common().UploadFile(ctx, req.File)
	if err != nil {
		return nil, err
	}

	return &v1.UploadFileRes{
		Name: fileInfo.Name,
		Url:  fileInfo.Url,
	}, err
}
