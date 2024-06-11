package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ClicksDevicesReq 统计分析接口 按设备
type ClicksDevicesReq struct {
	g.Meta   `path:"/analytics/clicksDevices" method:"get" summary:"根据设备统计短链访问次数" tags:"统计分析"`
	Code     string `json:"code" dc:"短链 code"`
	DateType string `json:"date" dc:"日期类型"`
	Type     string `json:"type" v:"required|in:devices,browsers,os" dc:"类型"`
}

type ClicksDevicesItem struct {
	Clicks int    `json:"clicks"   dc:"访问次数"`
	Name   string `json:"name" dc:"统计项名称"`
}

type ClicksDevicesRes []ClicksDevicesItem
