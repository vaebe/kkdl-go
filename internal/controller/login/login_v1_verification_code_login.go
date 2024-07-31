package login

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) VerificationCodeLogin(ctx context.Context, req *v1.VerificationCodeLoginReq) (res *v1.VerificationCodeLoginRes, err error) {
	// 获取缓存的验证码
	rdsKey := fmt.Sprintf("verificationCode-%s", req.Email)
	cacheVerificationCode, err := g.Redis().Get(ctx, rdsKey)

	if err != nil {
		return nil, gerror.New("获取缓存验证码失败!")
	}

	// 验证是否正确
	if cacheVerificationCode.String() != req.Code {
		return nil, gerror.New("验证码不正确!")
	}

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
