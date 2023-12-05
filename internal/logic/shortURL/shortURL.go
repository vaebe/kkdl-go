package shortURL

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
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
	_, err := g.Model("short_url").Save(g.Map{"ShortUrl": in.ShortUrl, "RawUrl": in.RawUrl})
	return err
}

// GetShortURL 获取短链
func (s *sShortURL) GetShortURL(ctx context.Context, in model.ShortURLCreateInput) (string, error) {
	return "", nil
}
