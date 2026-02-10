// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrlVisits is the golang structure for table short_url_visits.
type ShortUrlVisits struct {
	Id             uint64      `json:"id"             orm:"id"              ` //
	UserId         string      `json:"userId"         orm:"user_id"         ` // 用户id
	ShortUrl       string      `json:"shortUrl"       orm:"short_url"       ` // 短链,唯一，不能为空
	RawUrl         string      `json:"rawUrl"         orm:"raw_url"         ` // 原始 url 不能为空
	UserAgent      string      `json:"userAgent"      orm:"user_agent"      ` // 用户代理字符串，存储提供的完整用户代理
	BrowserName    string      `json:"browserName"    orm:"browser_name"    ` // 浏览器名称
	BrowserVersion string      `json:"browserVersion" orm:"browser_version" ` // 浏览器版本
	DeviceModel    string      `json:"deviceModel"    orm:"device_model"    ` // 设备型号
	EngineName     string      `json:"engineName"     orm:"engine_name"     ` // 浏览器引擎名称
	EngineVersion  string      `json:"engineVersion"  orm:"engine_version"  ` // 浏览器引擎版本
	OsName         string      `json:"osName"         orm:"os_name"         ` // 操作系统名称
	OsVersion      string      `json:"osVersion"      orm:"os_version"      ` // 操作系统版本
	Ip             string      `json:"ip"             orm:"ip"              ` // ip 不能为空
	Continent      string      `json:"continent"      orm:"continent"       ` // 大洲名称
	ContinentCode  string      `json:"continentCode"  orm:"continent_code"  ` // 大洲代码
	Country        string      `json:"country"        orm:"country"         ` // 国家名称
	CountryCode    string      `json:"countryCode"    orm:"country_code"    ` // 国家代码
	Region         string      `json:"region"         orm:"region"          ` // 地区或州的短代码（FIPS或ISO）
	RegionName     string      `json:"regionName"     orm:"region_name"     ` // 地区或州名称
	City           string      `json:"city"           orm:"city"            ` // 城市名称
	District       string      `json:"district"       orm:"district"        ` // 位置的区（郡）
	Lat            float64     `json:"lat"            orm:"lat"             ` // 纬度
	Lon            float64     `json:"lon"            orm:"lon"             ` // 经度
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      ` // 创建时间，默认为当前时间戳
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      ` // 更新时间
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      ` // 删除时间
}
