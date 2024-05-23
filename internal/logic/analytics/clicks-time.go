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
	startDate := gtime.NewFromStr("2024-05-20 00:00:00")
	endDate := gtime.NewFromStr("2024-05-20 23:59:59")

	// 生成唯一的临时表名
	hoursTableName := fmt.Sprintf("temp_hours_%s", grand.Digits(6))
	g.Log().Debug(ctx, hoursTableName)

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

// GetVisitsByDate 根据时间统计访问数据
func (s *sAnalytics) GetVisitsByDate(ctx context.Context, req v1.ClicksTimeReq) (list v1.ClicksTimeRes, err error) {
	if req.DateType == "24h" {
		list, err = statisticsByHour(ctx, req.Code)
		if err != nil {
			return nil, err
		}

		return list, err
	}

	return
}
