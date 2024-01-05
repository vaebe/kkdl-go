package user

import (
	"compressURL/api/user/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	userinfo := entity.User{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
		Role:     req.Role,
		Avatar:   req.Avatar,
	}

	id, err := service.User().Create(ctx, userinfo)

	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id}, nil
}
