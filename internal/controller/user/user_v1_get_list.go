package user

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"

	"compressURL/api/user/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	list, total, err := service.User().GetUserList(ctx, *req)

	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].Password = ""
		list[i].Salt = ""
	}

	return &v1.GetListRes{
		List: list,
		PageParams: model.PageParams{
			Total:    total,
			PageSize: req.PageSize,
			PageNo:   req.PageNo,
		},
	}, nil
}
