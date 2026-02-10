// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrl is the golang structure of table short_url for DAO operations like Where/Data.
type ShortUrl struct {
	g.Meta         `orm:"table:short_url, do:true"`
	Id             any         // 唯一标识，自增长整数
	ShortUrl       any         // 短链,唯一，不能为空
	RawUrl         any         // 原始 url 不能为空
	ExpirationTime *gtime.Time // 过期时间
	UserId         any         // 用户id
	CreatedAt      *gtime.Time // 创建时间，默认为当前时间戳
	Title          any         // 短链标题
	GroupId        any         // 短链分组id
}
