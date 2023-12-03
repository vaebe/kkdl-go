package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) UserTopic(ctx context.Context, req *v1.UserTopicReq) (res *v1.UserTopicRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
