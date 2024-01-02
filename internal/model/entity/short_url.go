// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrl is the golang structure for table short_url.
type ShortUrl struct {
	Id             int         `json:"id"             ` // 唯一标识，自增长整数
	ShortUrl       string      `json:"shortUrl"       ` // 短链,唯一，不能为空
	RawUrl         string      `json:"rawUrl"         ` // 原始 url 不能为空
	ExpirationTime *gtime.Time `json:"expirationTime" ` // 过期时间
	UserId         int         `json:"userId"         ` // 用户id
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间，默认为当前时间戳
	Title          string      `json:"title"          ` // 短链标题
	GroupId        int         `json:"groupId"        ` // 短链分组id
}
