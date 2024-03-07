package shortUrl

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/shortUrl/v1"
)

func (c *ControllerV1) GetUrl(ctx context.Context, req *v1.GetUrlReq) (res *v1.GetUrlRes, err error) {
	rawUrl, err := service.ShortUrl().GetShortUrl(ctx, req.ShortUrl)

	println(rawUrl)
	if err != nil {
		return nil, err
	}

	return &v1.GetUrlRes{RawUrl: rawUrl}, nil
}
