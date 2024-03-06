package v1

import "github.com/gogf/gf/v2/frame/g"

type RegistrationReq struct {
	g.Meta           `path:"/user/registration" method:"post" summary:"邮箱注册" tags:"登录"`
	Email            string `json:"email"  v:"required#请输入邮箱"   dc:"邮箱"`
	Password         string `json:"password" v:"required#请输入密码"   dc:"密码"`
	NickName         string `json:"nickName"   dc:"昵称"`
	VerificationCode string `json:"verificationCode"  v:"required#请输入验证码"   dc:"验证码"`
}
type RegistrationRes struct {
	Id string `json:"id" dc:"用户 id"`
}
