package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) UserArticle(ctx context.Context, req *v1.UserArticleReq) (res *v1.UserArticleRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
