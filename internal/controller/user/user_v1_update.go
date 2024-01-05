package user

import (
	"compressURL/api/user/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	userInfo := entity.User{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		NickName: req.NickName,
		Role:     req.Role,
		Avatar:   req.Avatar,
	}

	err = service.User().Update(ctx, userInfo)

	return nil, err
}
