package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) VerificationCodeLogin(ctx context.Context, req *v1.VerificationCodeLoginReq) (res *v1.VerificationCodeLoginRes, err error) {
	// 验证验证码
	valid, err := service.Common().VerifyTheVCode(ctx, req.Email, req.Code)
	if err != nil {
		return nil, gerror.New("验证验证码失败!")
	}
	if !valid {
		return nil, gerror.New("验证码不正确或已过期!")
	}

	// 标记验证码为已使用
	_ = service.Common().ConsumeVCode(ctx, req.Email, req.Code)

	userInfo, err := service.User().Detail(ctx, model.UserQueryInput{Email: req.Email})

	if err != nil {
		return nil, err
	}

	// 用户不存在则创建
	if userInfo == nil {
		userInfo = &entity.User{
			Email:       req.Email,
			Role:        "01",
			AccountType: "01",
		}

		if _, err = service.User().Create(ctx, *userInfo); err != nil {
			return nil, err
		}

		if userInfo, err = service.User().Detail(ctx, model.UserQueryInput{Email: userInfo.Email}); err != nil {
			return nil, err
		}
	}

	info := getLoginRes(ctx, *userInfo)
	return (*v1.VerificationCodeLoginRes)(info), nil
}
