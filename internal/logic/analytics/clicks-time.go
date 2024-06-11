package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"compressURL/internal/dao"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/net/context"
)

// 按小时统计
func statisticsByHour(ctx context.Context, shortUrl string, userId string) (list v1.ClicksTimeRes, err error) {
	now := gtime.Now()
	startDate := now.StartOfDay()
	endDate := now.EndOfDay()

	// 创建查询对象
	db := g.DB().Model("short_url_visits").
		Fields("HOUR(created_at) AS hour, COUNT(id) AS visit_count").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Where(dao.ShortUrlVisits.Columns().UserId, userId).
		Group("HOUR(created_at)").
		Order("HOUR(created_at)")

	// 如果 shortUrl 不为空，添加条件
	if shortUrl != "" {
		db = db.Where("short_url", shortUrl)
	}

	// 执行查询
	rows, err := db.All()
	if err != nil {
		g.Log().Error(ctx, "短链访问根据时间统计 查询数据失败", err)
		return
	}

	// 初始化每小时点击数为0
	hourlyClicks := make(map[int]int)
	for i := 0; i < 24; i++ {
		hourlyClicks[i] = 0
	}

	// 处理结果
	for _, row := range rows {
		hour := row["hour"].Int()
		visitCount := row["visit_count"].Int()
		hourlyClicks[hour] = visitCount
	}

	// 将结果转换为列表形式并按时间排序
	for hour := 0; hour < 24; hour++ {
		list = append(list, v1.ClicksTimeItem{
			Clicks: hourlyClicks[hour],
			Time:   fmt.Sprintf("%02d:00", hour),
		})
	}

	return
}

// statisticsByDays 根据传入天数进行统计
func statisticsByDays(ctx context.Context, shortUrl string, days int, userId string) (list v1.ClicksTimeRes, err error) {
	startDate := gtime.Now().AddDate(0, 0, -days).StartOfDay()
	endDate := gtime.Now().EndOfDay()

	// 创建查询对象
	db := g.DB().Model("short_url_visits").
		Fields("DATE(created_at) AS date, COUNT(id) AS visit_count").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Where(dao.ShortUrlVisits.Columns().UserId, userId).
		Group("DATE(created_at)").
		Order("DATE(created_at)")

	// 如果 shortUrl 不为空，添加条件
	if shortUrl != "" {
		db = db.Where("short_url", shortUrl)
	}

	// 执行查询
	rows, err := db.All()
	if err != nil {
		g.Log().Error(ctx, "根据日期统计失败", err)
		return
	}

	// 初始化每天点击数为 0
	dateClicks := make(map[string]int)
	for i := 0; i <= days; i++ {
		date := gtime.Now().AddDate(0, 0, -i).Format("Y-m-d")
		dateClicks[date] = 0
	}

	// 处理结果
	for _, row := range rows {
		date := row["date"].String()
		visitCount := row["visit_count"].Int()
		dateClicks[date] = visitCount
	}

	// 将结果转换为列表形式并按日期排序
	for i := days; i >= 0; i-- {
		date := gtime.Now().AddDate(0, 0, -i).Format("Y-m-d")
		list = append(list, v1.ClicksTimeItem{
			Clicks: dateClicks[date],
			Time:   date,
		})
	}

	return
}

// GetVisitsByDate 根据时间统计访问数据
func (s *sAnalytics) GetVisitsByDate(ctx context.Context, req v1.ClicksTimeReq, userId string) (list v1.ClicksTimeRes, err error) {
	if req.DateType == "24h" {
		return statisticsByHour(ctx, req.Code, userId)
	}

	if req.DateType == "7d" {
		return statisticsByDays(ctx, req.Code, 7, userId)
	}

	if req.DateType == "30d" {
		return statisticsByDays(ctx, req.Code, 30, userId)
	}

	return
}
