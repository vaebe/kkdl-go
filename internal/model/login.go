package model

type LoginInput struct {
	Email       string // 邮箱,唯一
	WxId        string // 小程序id,唯一
	Password    string // 密码, 小程序登录无密码
	AccountType string // 账号类型: 01 邮箱 02 小程序
}

type JWTPayloadInfo struct {
	Id          string `json:"id"`
	AccountType string `json:"accountType"`
	Role        string `json:"role"`
	Token       string `json:"Token"`
}
