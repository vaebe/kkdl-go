package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"compressURL/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

// GetVisitsByRegion 根据区域统计访问数据
func (s *sAnalytics) GetVisitsByRegion(ctx context.Context, req v1.ClicksRegionReq, userId string) (list v1.ClicksRegionRes, err error) {
	db := dao.ShortUrlVisits.Ctx(ctx).OmitEmptyWhere().
		Where(dao.ShortUrlVisits.Columns().ShortUrl, req.Code).
		Where(dao.ShortUrlVisits.Columns().UserId, userId)

	// 定义一个通用的处理函数
	queryAndLog := func(fields, groupBy string) error {
		return db.Fields(fields).
			Group(groupBy).
			Scan(&list)
	}

	switch req.Type {
	case "countries":
		err = queryAndLog(
			"country_code, country_code AS code, country AS name, COUNT(short_url) AS clicks",
			"country_code, country",
		)
	case "cities":
		err = queryAndLog(
			"country_code, region AS code, city AS name, COUNT(short_url) AS clicks",
			"country_code, region, city",
		)
	default:
		g.Log().Warning(ctx, "Unsupported type: ", req.Type)
		return list, nil
	}

	if err != nil {
		g.Log().Error(ctx, "Error executing query: ", err)
	}

	return
}
