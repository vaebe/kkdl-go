package user

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.User().Remove(ctx, req.UserId)
	return nil, err
}
