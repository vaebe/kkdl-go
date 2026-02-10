package model

type UserCreateInput struct {
	Email       string // 邮箱,唯一
	Password    string // 密码
	NickName    string // 昵称, 创建默认生成
	AccountType string // 账号类型: 01 邮箱
	Role        string // 角色: 00 admin 01 普通用户 02 vip
	Salt        string // 用户盐值
	Avatar      string // 用户头像
}

type UserUpdateInput struct {
	Id string // 唯一标识，自增长整数
	UserCreateInput
}

type UserQueryInput struct {
	Email string // 邮箱,唯一
	Id    string // id 唯一
}
