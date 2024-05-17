package shortUrl

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/shortUrl/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	userId := ""
	if loginUserInfo.Role == "01" {
		userId = loginUserInfo.Id
	}

	err = service.ShortUrl().Delete(ctx, req.Id, userId)
	return
}
