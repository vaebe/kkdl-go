// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrlCode is the golang structure for table short_url_code.
type ShortUrlCode struct {
	Id        int         `json:"id"        ` // 唯一标识，自增长整数
	Code      string      `json:"code"      ` // 短链,唯一，不能为空
	Status    int         `json:"status"    ` // 是否使用
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间，默认为当前时间戳
}
