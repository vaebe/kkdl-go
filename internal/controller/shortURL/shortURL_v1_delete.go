package shortURL

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/shortURL/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.ShortURL().Delete(ctx, req.Id)
	return
}
