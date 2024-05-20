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
	g.Meta          `orm:"table:short_url_visits, do:true"`
	Id              interface{} //
	UserId          interface{} // 用户id
	ShortUrl        interface{} // 短链,唯一，不能为空
	RawUrl          interface{} // 原始 url 不能为空
	Ip              interface{} // ip 不能为空
	UserAgent       interface{} // 客户端软件的类型、版本和其他相关信息
	SecChUa         interface{} // 客户端使用的浏览器和版本
	SecChUaMobile   interface{} // 请求是否来自移动设备。?0 表示不是移动设备，?1 表示是移动设备
	SecChUaPlatform interface{} // 客户端所运行的平台
	SecFetchUser    interface{} // 请求是否是用户发起的,?1 表示是用户发起的请求
	Continent       interface{} // 大洲名称
	ContinentCode   interface{} // 大洲代码
	Country         interface{} // 国家名称
	CountryCode     interface{} // 国家代码
	Region          interface{} // 地区或州的短代码（FIPS或ISO）
	RegionName      interface{} // 地区或州名称
	City            interface{} // 城市名称
	District        interface{} // 位置的区（郡）
	Lat             interface{} // 纬度
	Lon             interface{} // 经度
	CreatedAt       *gtime.Time // 创建时间，默认为当前时间戳
	UpdatedAt       *gtime.Time // 更新时间
	DeletedAt       *gtime.Time // 删除时间
}
