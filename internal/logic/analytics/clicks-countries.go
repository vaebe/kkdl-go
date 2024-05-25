package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"compressURL/internal/dao"
	"golang.org/x/net/context"
)

// GetVisitsByCountries 根据国家统计访问数据
func (s *sAnalytics) GetVisitsByCountries(ctx context.Context, req v1.ClicksCountriesReq) (list v1.ClicksCountriesRes, err error) {
	err = dao.ShortUrlVisits.Ctx(ctx).OmitEmptyWhere().
		Where(dao.ShortUrlVisits.Columns().ShortUrl, req.Code).
		Fields(dao.ShortUrlVisits.Columns().CountryCode, dao.ShortUrlVisits.Columns().Country).
		FieldCount(dao.ShortUrlVisits.Columns().ShortUrl, "clicks").
		Group(dao.ShortUrlVisits.Columns().CountryCode, dao.ShortUrlVisits.Columns().Country).
		Scan(&list)

	if err != nil {
		return nil, err
	}

	return
}
