package user

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) Remove(ctx context.Context, req *v1.RemoveReq) (res *v1.RemoveRes, err error) {
	err = service.User().Remove(ctx, req.UserId)
	return nil, err
}
