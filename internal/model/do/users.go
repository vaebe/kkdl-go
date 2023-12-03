// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta       `orm:"table:users, do:true"`
	UserId       interface{} // 用户唯一标识，自增长整数
	Username     interface{} // 用户名，唯一，不能为空
	Email        interface{} // 用户的电子邮件地址，唯一，不能为空
	PasswordHash interface{} // 存储密码的哈希值，不能为空
	FullName     interface{} // 用户的全名
	CreatedAt    *gtime.Time // 用户账号创建时间，默认为当前时间戳
	IsAdmin      interface{} // 表示用户是否是管理员，默认为 FALSE
}
