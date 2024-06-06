package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/net/context"
)

// 按小时统计
func statisticsByHour(ctx context.Context, shortUrl string) (list v1.ClicksTimeRes, err error) {
	now := gtime.Now()
	startDate := now.StartOfDay()
	endDate := now.EndOfDay()

	// 创建查询对象
	db := g.DB().Model("short_url_visits").
		Fields("HOUR(created_at) AS hour, COUNT(id) AS visit_count").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
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
func statisticsByDays(ctx context.Context, shortUrl string, day int) (list v1.ClicksTimeRes, err error) {
	startDate := gtime.Now().AddDate(0, 0, day)
	endDate := gtime.Now() // 今天

	// 原始SQL语句
	sql := `
		SELECT
			d.Date AS date,
			COUNT(suv.id) AS visit_count
		FROM (
			SELECT a.Date
			FROM (
				SELECT CURDATE() - INTERVAL (a.a + (10 * b.a) + (100 * c.a)) DAY AS Date
				FROM (SELECT 0 AS a UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9) a
				CROSS JOIN (SELECT 0 AS a UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9) b
				CROSS JOIN (SELECT 0 AS a UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9) c
			) a
			WHERE a.Date BETWEEN ? AND ?
			ORDER BY a.Date
		) d
		LEFT JOIN
		short_url_visits suv
		ON DATE(suv.created_at) = d.Date
		AND (? = '' OR suv.short_url = ?)
		GROUP BY
		d.Date
		ORDER BY
		d.Date;
	`

	// 执行原始SQL查询
	rows, err := g.DB().GetAll(ctx, sql, startDate.String(), endDate.String(), shortUrl, shortUrl)
	if err != nil {
		g.Log().Error(ctx, "根据日期统计失败", err)
	}

	// 处理结果
	for _, row := range rows {
		list = append(list, v1.ClicksTimeItem{
			Clicks: row["visit_count"].Int(),
			Time:   row["date"].String(),
		})
	}

	g.Log().Info(ctx, len(list))

	return
}

// GetVisitsByDate 根据时间统计访问数据
func (s *sAnalytics) GetVisitsByDate(ctx context.Context, req v1.ClicksTimeReq) (list v1.ClicksTimeRes, err error) {
	if req.DateType == "24h" {
		return statisticsByHour(ctx, req.Code)
	}

	if req.DateType == "7d" {
		return statisticsByDays(ctx, req.Code, -7)
	}

	if req.DateType == "30d" {
		return statisticsByDays(ctx, req.Code, -30)
	}

	return
}
