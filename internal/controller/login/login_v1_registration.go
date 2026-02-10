package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Registration(ctx context.Context, req *v1.RegistrationReq) (res *v1.RegistrationRes, err error) {
	// 验证验证码
	valid, err := service.Common().VerifyTheVCode(ctx, req.Email, req.VerificationCode)
	if err != nil {
		return nil, gerror.New("验证验证码失败!")
	}
	if !valid {
		return nil, gerror.New("验证码不正确或已过期!")
	}

	// 标记验证码为已使用
	_ = service.Common().ConsumeVCode(ctx, req.Email, req.VerificationCode)

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
