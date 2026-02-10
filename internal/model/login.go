package model

import "compressURL/internal/model/entity"

type LoginInput struct {
	Email       string // 邮箱,唯一
	Password    string // 密码
	AccountType string // 账号类型: 01 邮箱
}

type LoginRes struct {
	Token       string      `json:"token" dc:"jwt token"`
	TokenExpire string      `json:"tokenExpire" dc:"token 过期时间"`
	UserInfo    entity.User `json:"userInfo"   dc:"用户信息"`
}

type JWTPayloadInfo struct {
	Id          string `json:"id"`
	AccountType string `json:"accountType"`
	Role        string `json:"role"`
	Token       string `json:"Token"`
}
