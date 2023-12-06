package shortURL

import (
	"compressURL/api/shortURL/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
	"fmt"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 获取一条未使用短链 code
	codeData, err := dao.ShortUrlCode.Ctx(ctx).Where("status", 0).One()
	if err != nil {
		return nil, err
	}

	shortUrl := fmt.Sprintf("%s", codeData.GMap().Get("code"))

	err = service.ShortURL().CreateShortURL(ctx, model.ShortURLCreateInput{
		ShortUrl: shortUrl,
		RawUrl:   req.RawUrl,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		ShortUrl: shortUrl,
	}, nil
}
