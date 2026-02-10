package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"compressURL/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

// GetVisitsByDevice 根据设备统计访问数据
func (s *sAnalytics) GetVisitsByDevice(ctx context.Context, req v1.ClicksDevicesReq, userId string) (list v1.ClicksDevicesRes, err error) {
	db := dao.ShortUrlVisits.Ctx(ctx).OmitEmptyWhere().
		Where(dao.ShortUrlVisits.Columns().ShortUrl, req.Code).
		Where(dao.ShortUrlVisits.Columns().UserId, userId)

	// 定义一个通用的处理函数
	queryAndLog := func(fieldName, groupBy string) error {
		return db.Fields(fieldName).
			Group(groupBy).
			Scan(&list)
	}

	switch req.Type {
	case "devices":
		err = queryAndLog("device_model AS name, COUNT(`short_url`) AS clicks", dao.ShortUrlVisits.Columns().DeviceModel)
	case "browsers":
		err = queryAndLog("browser_name AS name, COUNT(`short_url`) AS clicks", dao.ShortUrlVisits.Columns().BrowserName)
	case "os":
		err = queryAndLog("os_name AS name, COUNT(`short_url`) AS clicks", dao.ShortUrlVisits.Columns().OsName)
	default:
		g.Log().Warning(ctx, "Unsupported type: ", req.Type)
		return list, nil
	}

	if err != nil {
		g.Log().Error(ctx, "Error executing query: ", err)
	}

	return
}
