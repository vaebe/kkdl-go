package common

import (
	"compressURL/api/common/v1"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

func (c *ControllerV1) GetVerificationCodeEmail(ctx context.Context, req *v1.GetVerificationCodeEmailReq) (res *v1.GetVerificationCodeEmailRes, err error) {
	code := grand.S(6)

	err = g.Redis().SetEX(ctx, req.Email, code, 60)
	if err != nil {
		return nil, gerror.New("redis 缓存邮箱验证码失败!")
	}

	err = service.Common().SendVerificationCodeEmail(ctx, code, req.Email)
	return nil, err
}
