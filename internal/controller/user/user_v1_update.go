package user

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	err = service.User().Update(ctx, model.UserUpdateInput{
		Id:       req.UserId,
		Password: req.Password,
		Nickname: req.Nickname,
	})

	return nil, err
}
