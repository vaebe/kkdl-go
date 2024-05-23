package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 统计分析接口 按国家统计

type ClicksCountriesReq struct {
	g.Meta   `path:"/analytics/clicksCountries" method:"get" summary:"根据国家统计短链访问次数" tags:"统计分析"`
	Code     string `json:"code"   dc:"短链 code"`
	DateType string `json:"date" dc:"日期类型"`
}

type ClicksCountriesItem struct {
	Clicks      int    `json:"clicks"   dc:"访问次数"`
	Country     string `json:"country" dc:"国家 code"`
	CountryName string `json:"countryName" dc:"国家名称"`
}

type ClicksCountriesRes []ClicksCountriesItem
