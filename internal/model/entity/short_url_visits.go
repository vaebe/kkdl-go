// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrlVisits is the golang structure for table short_url_visits.
type ShortUrlVisits struct {
	Id              uint64      `json:"id"              ` //
	UserId          string      `json:"userId"          ` // 用户id
	ShortUrl        string      `json:"shortUrl"        ` // 短链,唯一，不能为空
	RawUrl          string      `json:"rawUrl"          ` // 原始 url 不能为空
	Ip              string      `json:"ip"              ` // ip 不能为空
	UserAgent       string      `json:"userAgent"       ` // 客户端软件的类型、版本和其他相关信息
	SecChUa         string      `json:"secChUa"         ` // 客户端使用的浏览器和版本
	SecChUaMobile   string      `json:"secChUaMobile"   ` // 请求是否来自移动设备。?0 表示不是移动设备，?1 表示是移动设备
	SecChUaPlatform string      `json:"secChUaPlatform" ` // 客户端所运行的平台
	SecFetchUser    string      `json:"secFetchUser"    ` // 请求是否是用户发起的,?1 表示是用户发起的请求
	Continent       string      `json:"continent"       ` // 大洲名称
	ContinentCode   string      `json:"continentCode"   ` // 大洲代码
	Country         string      `json:"country"         ` // 国家名称
	CountryCode     string      `json:"countryCode"     ` // 国家代码
	Region          string      `json:"region"          ` // 地区或州的短代码（FIPS或ISO）
	RegionName      string      `json:"regionName"      ` // 地区或州名称
	City            string      `json:"city"            ` // 城市名称
	District        string      `json:"district"        ` // 位置的区（郡）
	Lat             float64     `json:"lat"             ` // 纬度
	Lon             float64     `json:"lon"             ` // 经度
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间，默认为当前时间戳
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       ` // 删除时间
}
