package shortURL

import (
	"compressURL/api/shortURL/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 获取一条未使用短链 code
	codeData, err := dao.ShortUrlCode.Ctx(ctx).Where("status", 0).One()
	if err != nil {
		return nil, err
	}

	code := fmt.Sprintf("%s", codeData.GMap().Get("code"))

	err = service.ShortURL().CreateShortURL(ctx, model.ShortURLCreateInput{
		ShortUrl: code,
		RawUrl:   req.RawUrl,
		// 默认七天
		ExpirationTime: time.Now().AddDate(0, 0, 7).UTC().Format("2006-01-02"),
	})

	if err != nil {
		return nil, err
	}

	// 修改已使用的 code 状态
	_, err = dao.ShortUrlCode.Ctx(ctx).Data(g.Map{"status": 1}).Where("code", code).Update()
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		ShortUrl: code,
	}, nil
}
