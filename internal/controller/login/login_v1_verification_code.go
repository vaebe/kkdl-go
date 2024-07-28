package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) VerificationCode(ctx context.Context, req *v1.VerificationCodeReq) (res *v1.VerificationCodeRes, err error) {
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

	// todo 判断是否是数据库的错误 在抛出去
	if err != nil {
		userInfo = entity.User{
			Email:       req.Email,
			Role:        "01",
			AccountType: "01",
		}

		if _, err = service.User().Create(ctx, userInfo); err != nil {
			return nil, err
		}
	}

	// 设置登录用户信息
	g.RequestFromCtx(ctx).SetCtxVar("loginInfo", userInfo)
	token, expire := service.Auth().AuthInstance().LoginHandler(ctx)
	tokenExpire := gtime.NewFromTime(expire).Format("Y-m-d H:i:s")

	return &v1.VerificationCodeRes{
		Token:       token,
		TokenExpire: tokenExpire,
		UserInfo: entity.User{
			Id:          userInfo.Id,
			Email:       userInfo.Email,
			NickName:    userInfo.NickName,
			AccountType: userInfo.AccountType,
			Role:        userInfo.Role,
			Avatar:      userInfo.Avatar,
		},
	}, nil
}
