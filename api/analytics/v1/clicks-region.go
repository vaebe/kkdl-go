package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ClicksRegionReq 统计分析接口 按区域统计
type ClicksRegionReq struct {
	g.Meta   `path:"/analytics/clicksRegion" method:"get" summary:"根据区域统计短链访问次数" tags:"统计分析"`
	Code     string `json:"code"   dc:"短链 code"`
	DateType string `json:"date" dc:"日期类型"`
	Type     string `json:"type" v:"required|in:countries,cities" dc:"类型"`
}

type ClicksRegionItem struct {
	Clicks      int    `json:"clicks"   dc:"访问次数"`
	Code        string `json:"code" dc:"区域 code"`
	Name        string `json:"name" dc:"区域名称"`
	CountryCode string `json:"countryCode" dc:"国家 code"`
}

type ClicksRegionRes []ClicksRegionItem
