package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Registration(ctx context.Context, req *v1.RegistrationReq) (res *v1.RegistrationRes, err error) {
	// 原子性验证并消费验证码（防止 TOCTOU）
	valid, err := service.Common().VerifyAndConsumeVCode(ctx, req.Email, req.VerificationCode)
	if err != nil {
		return nil, gerror.Wrap(err, "验证验证码失败!")
	}
	if !valid {
		return nil, gerror.New("验证码不正确或已过期!")
	}

	userinfo := entity.User{
		Email:       req.Email,
		Password:    req.Password,
		NickName:    req.NickName,
		AccountType: "01",
		Role:        "01",
	}

	id, err := service.User().Create(ctx, userinfo)

	if err != nil {
		return nil, err
	}
	return &v1.RegistrationRes{Id: id}, nil
}
