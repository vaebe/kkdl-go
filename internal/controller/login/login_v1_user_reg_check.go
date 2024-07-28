package login

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) UserRegCheck(ctx context.Context, req *v1.UserRegCheckReq) (res *v1.UserRegCheckRes, err error) {
	res = &v1.UserRegCheckRes{}

	info, err := service.User().GetOne(ctx, model.UserQueryInput{Email: req.Email})

	if info == nil || err != nil {
		res.IsRegistered = false
		return res, nil
	}

	res.IsRegistered = true

	return res, nil
}
