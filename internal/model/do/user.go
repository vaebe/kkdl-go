// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:user, do:true"`
	Id          any         // 唯一标识
	Email       any         // 邮箱,唯一
	Password    any         // 密码
	NickName    any         // 昵称, 创建默认生成
	AccountType any         // 账号类型: 01 邮箱
	Role        any         // 角色: 00 admin 01 普通用户 02 vip
	DeletedAt   *gtime.Time // 删除时间
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
	Salt        any         // 用户盐值
	Avatar      any         // 用户头像
}
