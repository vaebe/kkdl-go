package shortURL

import (
	"context"

	"compressURL/api/shortURL/v1"
)

func (c *ControllerV1) GetUrl(ctx context.Context, req *v1.GetUrlReq) (res *v1.GetUrlRes, err error) {

	return &v1.GetUrlRes{RawUrl: req.ShortUrl}, nil
}
