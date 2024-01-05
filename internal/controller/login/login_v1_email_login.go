package login

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
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
		UserInfo: entity.User{
			Id:          userInfo.Id,
			Email:       userInfo.Email,
			WxId:        userInfo.WxId,
			NickName:    userInfo.NickName,
			AccountType: userInfo.AccountType,
			Role:        userInfo.Role,
		},
	}, nil
}
