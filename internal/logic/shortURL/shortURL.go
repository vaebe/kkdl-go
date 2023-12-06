package shortURL

import (
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

type sShortURL struct {
}

func init() {
	url := New()
	service.RegisterShortURL(url)
}

func New() *sShortURL {
	return &sShortURL{}
}

// CreateShortURL 创建短链
func (s *sShortURL) CreateShortURL(ctx context.Context, in model.ShortURLCreateInput) error {
	_, err := dao.ShortUrl.Ctx(ctx).Save(g.Map{"ShortUrl": in.ShortUrl, "RawUrl": in.RawUrl})
	return err
}

// GetShortURL 获取短链
func (s *sShortURL) GetShortURL(ctx context.Context, url string) (string, error) {
	one, err := dao.ShortUrl.Ctx(ctx).Where("ShortUrl", url).One()
	if err != nil {
		return "", err
	}

	rawUrl := fmt.Sprintf("%s", one.GMap().Get("rawUrl"))
	return rawUrl, nil
}
