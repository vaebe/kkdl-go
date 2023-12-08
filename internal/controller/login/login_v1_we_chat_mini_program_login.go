package login

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) WeChatMiniProgramLogin(ctx context.Context, req *v1.WeChatMiniProgramLoginReq) (res *v1.WeChatMiniProgramLoginRes, err error) {

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
