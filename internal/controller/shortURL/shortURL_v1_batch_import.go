package shortURL

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/shortURL/v1"
)

func (c *ControllerV1) BatchImport(ctx context.Context, req *v1.BatchImportReq) (res *v1.BatchImportRes, err error) {
	err = service.ShortURL().BatchImport(ctx, req.File)
	return
}
