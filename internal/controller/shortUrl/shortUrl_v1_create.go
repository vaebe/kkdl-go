package shortUrl

import (
	"compressURL/api/shortUrl/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 获取当前登陆用户信息,获取不到用户信息即视为未登录不赋值用户id
	loginUserInfo, _ := service.Auth().GetLoginUserInfo(ctx)

	info := entity.ShortUrl{
		RawUrl:         req.RawUrl,
		Title:          req.Title,
		GroupId:        req.GroupId,
		UserId:         loginUserInfo.Id,
		ExpirationTime: gtime.NewFromStrFormat(req.ExpirationTime, "Y-m-d H:i:s"),
	}

	code, err := service.ShortUrl().CreateShortUrl(ctx, info)

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{ShortUrl: code}, nil
}
