package user

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	info, err := service.User().GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return info, nil
}
