// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VerificationCodes is the golang structure of table verification_codes for DAO operations like Where/Data.
type VerificationCodes struct {
	g.Meta    `orm:"table:verification_codes, do:true"`
	Id        any         //
	Email     any         // 邮箱
	Type      any         // 验证码类型: 目前只有一种用途
	Code      any         // 验证码
	Used      any         // 是否已使用: 0未使用 1已使用
	ExpiredAt *gtime.Time // 过期时间
	CreatedAt *gtime.Time // 创建时间
}
