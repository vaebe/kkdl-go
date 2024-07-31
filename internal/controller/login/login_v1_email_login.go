package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"compressURL/utility"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error) {
	// 获取用户信息
	userInfo, err := service.User().Detail(ctx, model.UserQueryInput{Email: req.Email})

	if err != nil {
		return nil, err
	}

	// 用户不存在则创建用户
	if userInfo == nil {
		return nil, gerror.New("您还没有注册!")
	}

	if utility.EncryptPassword(req.Password, userInfo.Salt) != userInfo.Password {
		return nil, gerror.New("账号或者密码不正确!")
	}

	info := getLoginRes(ctx, *userInfo)
	return (*v1.EmailLoginRes)(info), nil
}
