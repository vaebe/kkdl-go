package shortURL

import (
	"compressURL/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/shortURL/v1"
)

func (c *ControllerV1) BatchCreateShortUrlCode(ctx context.Context, req *v1.BatchCreateShortUrlCodeReq) (res *v1.BatchCreateShortUrlCodeRes, err error) {
	// 获取当前登陆用户信息
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	if loginUserInfo.Role != "00" {
		return nil, gerror.New("无权限!")
	}

	if req.Num > 1000 {
		return nil, gerror.New("每次最多生成 1000 个!")
	}

	err = service.ShortURL().BatchCreateShortUrlCode(ctx, req.Num)
	return nil, err
}
