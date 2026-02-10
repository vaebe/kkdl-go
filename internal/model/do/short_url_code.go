// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrlCode is the golang structure of table short_url_code for DAO operations like Where/Data.
type ShortUrlCode struct {
	g.Meta    `orm:"table:short_url_code, do:true"`
	Id        any         // 唯一标识，自增长整数
	Code      any         // 短链,唯一，不能为空
	Status    any         // 是否使用
	CreatedAt *gtime.Time // 创建时间，默认为当前时间戳
}
