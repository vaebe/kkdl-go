package model

type UserCreateInput struct {
	Email       string // 邮箱,唯一
	WxId        string // 小程序id,唯一
	Password    string // 密码, 小程序登录无密码
	Nickname    string // 昵称, 创建默认生成
	AccountType string // 账号类型: 01 邮箱 02 小程序
	Role        string // 角色: 00 admin 01 普通用户 02 vip
	Salt        string // 用户盐值
	Avatar      string // 用户头像
}

type UserUpdateInput struct {
	Id string // 唯一标识，自增长整数
	UserCreateInput
}
