package shortUrl

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/shortUrl/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.ShortUrl().Delete(ctx, req.Id)
	return
}
