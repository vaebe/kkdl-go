package shortURL

import (
	"compressURL/api/shortURL/v1"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
	"time"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	code, err := service.ShortURL().CreateShortURL(ctx, model.ShortURLCreateInput{
		RawUrl: req.RawUrl,
		// 默认七天
		ExpirationTime: time.Now().AddDate(0, 0, 7).UTC().Format("2006-01-02"),
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{ShortUrl: code}, nil
}
