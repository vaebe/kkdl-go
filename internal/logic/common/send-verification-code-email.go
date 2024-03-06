package common

import (
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jordan-wright/email"
	"golang.org/x/net/context"
	"net/smtp"
)

// SendVerificationCodeEmail 发送邮箱验证码
func (s *sCommon) SendVerificationCodeEmail(ctx context.Context, VerificationCode string, emailAddress string) (err error) {
	mailUserName, _ := g.Cfg().Get(ctx, "emailConfig.email") // 邮箱账号
	mailPassword, _ := g.Cfg().Get(ctx, "emailConfig.key")   // 邮箱授权码
	addr := "smtp.qq.com:465"                                // TLS地址
	host := "smtp.qq.com"                                    // 邮件服务器地址
	Subject := "KKDL验证码"                                     // 发送的主题

	e := email.NewEmail()
	e.From = fmt.Sprintf("KKDL <%s>", mailUserName)
	e.To = []string{emailAddress}
	e.Subject = Subject
	e.HTML = []byte("您的验证码为：<h1>" + VerificationCode + "</h1>")
	return e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName.String(), mailPassword.String(), host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}
