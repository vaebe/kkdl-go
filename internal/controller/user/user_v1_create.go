package user

import (
	"compressURL/api/user/v1"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	userinfo := model.UserCreateInput{
		Email:       req.Email,
		WxId:        req.WxId,
		Password:    req.Password,
		Nickname:    req.Nickname,
		AccountType: "01",
		Role:        "01",
		Salt:        "",
	}

	if req.WxId != "" {
		userinfo.AccountType = "02"
	}

	id, err := service.User().Create(ctx, userinfo)

	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id}, nil
}
