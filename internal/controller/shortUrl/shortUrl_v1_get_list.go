package shortUrl

import (
	"compressURL/api/shortUrl/v1"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	userId := ""
	if loginUserInfo.Role == "01" {
		userId = loginUserInfo.Id
	}

	list, total, err := service.ShortURL().GetList(ctx, *req, userId)

	if err != nil {
		return nil, err
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
