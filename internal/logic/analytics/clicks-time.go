package analytics

import (
	v1 "compressURL/api/analytics/v1"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"golang.org/x/net/context"
)

// 按小时统计
func statisticsByHour(ctx context.Context, shortUrl string) (list v1.ClicksTimeRes, err error) {
	now := gtime.Now()
	startDate := now.StartOfDay()
	endDate := now.EndOfDay()

	// 生成唯一的临时表名
	hoursTableName := fmt.Sprintf("temp_hours_%s", grand.Digits(6))

	_, err = g.DB().Exec(ctx, fmt.Sprintf(`CREATE TEMPORARY TABLE %s (hour INT PRIMARY KEY);`, hoursTableName))
	if err != nil {
		g.Log().Error(ctx, "短链访问根据时间统计创建临时表失败", err)
		return
	}

	// 创建临时表
	sql := fmt.Sprintf(`INSERT INTO %s (hour) VALUES (0), (1), (2), (3), (4), (5), (6), (7), (8), (9), (10), (11),(12), (13), (14), (15), (16), (17), (18), (19), (20), (21), (22), (23);`, hoursTableName)
	_, err = g.DB().Exec(ctx, sql)
	if err != nil {
		g.Log().Error(ctx, "短链访问根据时间统计 hours 临时表插入数据失败", err)
		return
	}

	// 构建查询
	query := fmt.Sprintf(`
	SELECT 
		h.hour,
		COUNT(suv.id) AS visit_count
	FROM 
		%s h
	LEFT JOIN 
		short_url_visits suv 
		ON HOUR(suv.created_at) = h.hour
		AND suv.created_at BETWEEN ? AND ?
	`, hoursTableName)
	args := []interface{}{startDate.String(), endDate.String()}

	// 如果shortUrl不为空，添加条件
	if shortUrl != "" {
		query += " AND suv.short_url = ?"
		args = append(args, shortUrl)
	}

	query += " GROUP BY h.hour ORDER BY h.hour;"

	// 执行查询
	rows, err := g.DB().GetAll(ctx, query, args...)
	if err != nil {
		g.Log().Error(ctx, "短链访问根据时间统计 查询数据失败", err)

		return
	}

	// 处理结果
	for _, row := range rows {
		list = append(list, v1.ClicksTimeItem{
			Clicks: row["visit_count"].Int(),
			Time:   row["hour"].String(),
		})
	}

	// 删除临时表
	_, err = g.DB().Exec(ctx, fmt.Sprintf("DROP TEMPORARY TABLE %s;", hoursTableName))
	if err != nil {
		g.Log().Error(ctx, "短链访问根据时间统计 删除临时表失败", err)
		return
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
