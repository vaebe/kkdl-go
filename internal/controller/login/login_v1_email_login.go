package login

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error) {
	userInfo, token, tokenExpire, err := service.Login().UserLogin(ctx, model.LoginInput{Email: req.Email, Password: req.Password, AccountType: "01"})
	if err != nil {
		return nil, err
	}

	return &v1.EmailLoginRes{
		Token:       token,
		TokenExpire: tokenExpire,
		UserInfo: v1.UserInfo{
			Id:          int64(userInfo.Id),
			Email:       userInfo.Email,
			WxId:        userInfo.WxId,
			Nickname:    userInfo.Nickname,
			AccountType: userInfo.AccountType,
			Role:        userInfo.Role,
		},
	}, nil
}