package common

import (
	"compressURL/api/common/v1"
	v12 "compressURL/api/user/v1"
	"compressURL/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

func (c *ControllerV1) GetVerificationCodeEmail(ctx context.Context, req *v1.GetVerificationCodeEmailReq) (res *v1.GetVerificationCodeEmailRes, err error) {
	// 检查用户是否已经注册
	_, total, err := service.User().GetUserList(ctx, v12.GetListReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	if total != 0 {
		return nil, gerror.New("用户已注册请直接登录!")
	}

	// 检查验证码是否过期
	rdsKey := fmt.Sprintf("verificationCode-%s", req.Email)
	ttl, err := g.Redis().TTL(ctx, rdsKey)
	if err != nil {
		return nil, err
	}

	// ttl -1 没有设置过期时间,-2 键已经过期或不存在
	if ttl != -1 && ttl != -2 {
		return nil, gerror.Newf("请勿重复请求,有效期还剩余 %d 秒!", ttl)
	}

	// 发送验证码
	code := grand.S(6)
	err = g.Redis().SetEX(ctx, rdsKey, code, 60*2)
	if err != nil {
		return nil, gerror.New("redis 缓存邮箱验证码失败!")
	}

	err = service.Common().SendVerificationCodeEmail(ctx, code, req.Email)
	return nil, err
}
