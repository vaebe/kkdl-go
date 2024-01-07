package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, _ *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	token, expire := service.Auth().AuthInstance().RefreshHandler(ctx)
	expireStr := gtime.NewFromTime(expire).Format("Y-m-d H:i:s")

	return &v1.RefreshTokenRes{
		Token:  token,
		Expire: expireStr,
	}, nil
}
