package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
