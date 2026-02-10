// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VerificationCodes is the golang structure for table verification_codes.
type VerificationCodes struct {
	Id        int         `json:"id"        orm:"id"         ` //
	Email     string      `json:"email"     orm:"email"      ` // 邮箱
	Type      int         `json:"type"      orm:"type"       ` // 验证码类型: 目前只有一种用途
	Code      string      `json:"cODE"      orm:"CODE"       ` // 验证码
	Used      int         `json:"used"      orm:"used"       ` // 是否已使用: 0未使用 1已使用
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" ` // 过期时间
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
}
