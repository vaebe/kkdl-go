// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id          string      `json:"id"          orm:"id"          ` // 唯一标识
	Email       string      `json:"email"       orm:"email"       ` // 邮箱,唯一
	Password    string      `json:"password"    orm:"password"    ` // 密码, 小程序登录无密码
	NickName    string      `json:"nickName"    orm:"nickName"    ` // 昵称, 创建默认生成
	AccountType string      `json:"accountType" orm:"accountType" ` // 账号类型: 01 邮箱 02 小程序
	Role        string      `json:"role"        orm:"role"        ` // 角色: 00 admin 01 普通用户 02 vip
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  ` // 删除时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  ` // 创建时间
	Salt        string      `json:"salt"        orm:"salt"        ` // 用户盐值
	Avatar      string      `json:"avatar"      orm:"avatar"      ` // 用户头像
}
