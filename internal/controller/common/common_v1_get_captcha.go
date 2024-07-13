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

func (c *ControllerV1) GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (res *v1.GetCaptchaRes, err error) {

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

	// 一分钟内只能请求一次
	countdown := ttl - 540

	/**
	ttl -1 没有设置过期时间,-2 键已经过期或不存在
	且剩余时间大于 1
	*/
	if ttl != -1 && ttl != -2 && countdown > 1 {
		return nil, gerror.Newf("请勿重复请求,请等待 %d 秒后在进行操作!", countdown)
	}

	// 发送验证码
	code := grand.N(100000, 999999)
	err = g.Redis().SetEX(ctx, rdsKey, code, 60*10)
	if err != nil {
		return nil, gerror.New("redis 缓存邮箱验证码失败!")
	}

	err = service.Common().SendVerificationCodeEmail(ctx, code, req.Email)
	return nil, err
}
