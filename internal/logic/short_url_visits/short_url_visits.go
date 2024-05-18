package short_url_visits

import (
	"compressURL/internal/dao"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"golang.org/x/net/context"
)

type sShortUrlVisits struct {
}

func init() {
	service.RegisterShortUrlVisits(New())
}

func New() *sShortUrlVisits {
	return &sShortUrlVisits{}
}

func (s *sShortUrlVisits) Create(ctx context.Context, in entity.ShortUrlVisits) error {
	_, err := dao.ShortUrlVisits.Ctx(ctx).Data(in).Save()
	return err
}
