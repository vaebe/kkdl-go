package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"compressURL/internal/dao"
	"golang.org/x/net/context"
)

// GetVisitsByRegion 根据区域统计访问数据
func (s *sAnalytics) GetVisitsByRegion(ctx context.Context, req v1.ClicksRegionReq, userId string) (list v1.ClicksRegionRes, err error) {
	err = dao.ShortUrlVisits.Ctx(ctx).OmitEmptyWhere().
		Where(dao.ShortUrlVisits.Columns().ShortUrl, req.Code).
		Where(dao.ShortUrlVisits.Columns().UserId, userId).
		Fields(dao.ShortUrlVisits.Columns().CountryCode, dao.ShortUrlVisits.Columns().Country).
		FieldCount(dao.ShortUrlVisits.Columns().ShortUrl, "clicks").
		Group(dao.ShortUrlVisits.Columns().CountryCode, dao.ShortUrlVisits.Columns().Country).
		Scan(&list)

	if err != nil {
		return nil, err
	}

	return
}
