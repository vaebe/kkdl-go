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
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间，默认为当前时间戳
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       ` // 删除时间
}
