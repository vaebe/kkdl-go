package user

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	info, err := service.User().GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return info, nil
}
