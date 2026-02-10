// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrlVisits is the golang structure of table short_url_visits for DAO operations like Where/Data.
type ShortUrlVisits struct {
	g.Meta         `orm:"table:short_url_visits, do:true"`
	Id             any         //
	UserId         any         // 用户id
	ShortUrl       any         // 短链,唯一，不能为空
	RawUrl         any         // 原始 url 不能为空
	UserAgent      any         // 用户代理字符串，存储提供的完整用户代理
	BrowserName    any         // 浏览器名称
	BrowserVersion any         // 浏览器版本
	DeviceModel    any         // 设备型号
	EngineName     any         // 浏览器引擎名称
	EngineVersion  any         // 浏览器引擎版本
	OsName         any         // 操作系统名称
	OsVersion      any         // 操作系统版本
	Ip             any         // ip 不能为空
	Continent      any         // 大洲名称
	ContinentCode  any         // 大洲代码
	Country        any         // 国家名称
	CountryCode    any         // 国家代码
	Region         any         // 地区或州的短代码（FIPS或ISO）
	RegionName     any         // 地区或州名称
	City           any         // 城市名称
	District       any         // 位置的区（郡）
	Lat            any         // 纬度
	Lon            any         // 经度
	CreatedAt      *gtime.Time // 创建时间，默认为当前时间戳
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
