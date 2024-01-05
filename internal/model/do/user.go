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
	Id          interface{} // 唯一标识
	Email       interface{} // 邮箱,唯一
	WxId        interface{} // 小程序id,唯一
	Password    interface{} // 密码, 小程序登录无密码
	Nickname    interface{} // 昵称, 创建默认生成
	AccountType interface{} // 账号类型: 01 邮箱 02 小程序
	Role        interface{} // 角色: 00 admin 01 普通用户 02 vip
	DeletedAt   *gtime.Time // 删除时间
	UpdatedAt   *gtime.Time // 更新时间
	CreatedAt   *gtime.Time // 创建时间
	Salt        interface{} // 用户盐值
	Avatar      interface{} // 用户头像
}
