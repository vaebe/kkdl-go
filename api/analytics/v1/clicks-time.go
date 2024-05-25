package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 统计分析接口 按时间统计

type ClicksTimeReq struct {
	g.Meta   `path:"/analytics/clicksTime" method:"get" summary:"根据时间统计短链访问次数" tags:"统计分析"`
	Code     string `json:"code" dc:"短链 code"`
	DateType string `json:"dateType" v:"required|in:24h,7d,30d" dc:"日期类型"`
}

type ClicksTimeItem struct {
	Clicks int    `json:"clicks"   dc:"访问次数"`
	Time   string `json:"time" dc:"日期"`
}

type ClicksTimeRes []ClicksTimeItem

func (c ClicksTimeRes) Error() string {
	//TODO implement me
	panic("implement me")
}
