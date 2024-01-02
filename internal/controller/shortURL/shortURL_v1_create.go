package shortURL

import (
	"compressURL/api/shortURL/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	info := entity.ShortUrl{
		RawUrl:         req.RawUrl,
		Title:          req.Title,
		GroupId:        req.GroupId,
		ExpirationTime: gtime.NewFromStrFormat(req.ExpirationTime, "Y-m-d H:i:s"),
	}

	code, err := service.ShortURL().CreateShortURL(ctx, info)

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{ShortUrl: code}, nil
}
