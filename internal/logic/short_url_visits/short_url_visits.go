package short_url_visits

import (
	v1 "compressURL/api/short_url_visits/v1"
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

// GetList 短链列表
func (s *sShortUrlVisits) GetList(ctx context.Context, in v1.GetListReq, userId string) ([]entity.ShortUrlVisits, int, error) {
	var list []entity.ShortUrlVisits

	db := dao.ShortUrlVisits.Ctx(ctx).OmitEmptyWhere().
		Where(dao.ShortUrlVisits.Columns().ShortUrl, in.Code)

	// 用户 id 存在只查询当前用户的数据
	if userId != "" {
		db = db.Where(dao.ShortUrlVisits.Columns().UserId, userId)
	}

	total, _ := db.Count()

	err := db.OrderDesc(dao.ShortUrlVisits.Columns().CreatedAt).Page(in.PageNo, in.PageSize).Scan(&list)

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
