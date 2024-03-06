package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Registration(ctx context.Context, req *v1.RegistrationReq) (res *v1.RegistrationRes, err error) {

	// 获取缓存的验证码
	cacheVerificationCode, err := g.Redis().Get(ctx, req.Email)
	if err != nil {
		return nil, gerror.New("获取缓存验证码失败!")
	}

	// 验证是否正确
	if cacheVerificationCode.String() != req.VerificationCode {
		return nil, gerror.New("验证码不正确!")
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
