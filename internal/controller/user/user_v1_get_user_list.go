package user

import (
	"compressURL/api/user/v1"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	list, total, err := service.User().GetUserList(ctx, *req)

	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].Password = ""
		list[i].Salt = ""
	}

	return &v1.GetUserListRes{
		List: list,
		PageParams: model.PageParams{
			Total:    total,
			PageSize: req.PageSize,
			PageNo:   req.PageNo,
		},
	}, nil
}
