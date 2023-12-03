// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	UserId       int         `json:"userId"       ` // 用户唯一标识，自增长整数
	Username     string      `json:"username"     ` // 用户名，唯一，不能为空
	Email        string      `json:"email"        ` // 用户的电子邮件地址，唯一，不能为空
	PasswordHash string      `json:"passwordHash" ` // 存储密码的哈希值，不能为空
	FullName     string      `json:"fullName"     ` // 用户的全名
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 用户账号创建时间，默认为当前时间戳
	IsAdmin      int         `json:"isAdmin"      ` // 表示用户是否是管理员，默认为 FALSE
}
