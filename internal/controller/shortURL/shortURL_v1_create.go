package shortURL

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
	"compressURL/utility"
	"context"
	"math/rand"
	"time"

	"compressURL/api/shortURL/v1"
)

func generateRandomNumber() int {
	// 设置随机数种子，确保每次运行程序时都得到不同的结果
	rand.NewSource(time.Now().UnixNano())
	max := 9999999999
	min := 1
	return rand.Intn(max-min+1) + min
}

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	ShortUrl := utility.IntToBase62(generateRandomNumber())

	err = service.ShortURL().CreateShortURL(ctx, model.ShortURLCreateInput{
		ShortUrl: ShortUrl,
		RawUrl:   "www.baidu.com",
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		ShortUrl: ShortUrl,
	}, nil
}
