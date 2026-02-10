package common

import (
	"compressURL/api/common/v1"
	"compressURL/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (res *v1.GetCaptchaRes, err error) {
	// 检查冷却期（1分钟内只能请求一次）
	inCoolDown, remaining, err := service.Common().CheckVCodeCooldown(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if inCoolDown {
		return nil, gerror.Newf("请勿重复请求,请等待 %d 秒后再进行操作!", remaining)
	}

	// 清理过期验证码
	_ = service.Common().DeleteExpiredVCodes(ctx)

	// 生成并保存验证码
	codeStr, err := service.Common().CreateVCode(ctx, req.Email)
	if err != nil {
		return nil, gerror.New("保存验证码失败!")
	}

	// 发送邮件
	code := gconv.Int(codeStr)
	err = service.Common().SendEmailVCode(ctx, code, req.Email)
	return nil, err
}
