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
func (s *sCommon) SendVerificationCodeEmail(ctx context.Context, VerificationCode int, emailAddress string) (err error) {
	mailUserName, _ := g.Cfg().Get(ctx, "emailConfig.email") // 邮箱账号
	mailPassword, _ := g.Cfg().Get(ctx, "emailConfig.key")   // 邮箱授权码
	addr := "smtp.qq.com:465"                                // TLS地址
	host := "smtp.qq.com"                                    // 邮件服务器地址
	Subject := "KKDL验证码"                                     // 发送的主题

	e := email.NewEmail()
	e.From = fmt.Sprintf("KKDL <%s>", mailUserName)
	e.To = []string{emailAddress}
	e.Subject = Subject

	// 美化后的HTML内容
	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>%s</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f4f4f4;
				color: #333;
				padding: 20px;
			}
			.container {
				max-width: 600px;
				margin: 0 auto;
				background-color: #fff;
				padding: 20px;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}
			h1 {
				color: #007BFF;
			}
			p {
				font-size: 16px;
				line-height: 1.5;
			}
			.footer {
				margin-top: 20px;
				font-size: 12px;
				color: #777;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>验证码</h1>
			<p>您好，</p>
			<p>您的验证码为：<strong>%d</strong></p>
			<p>请在10分钟内使用此验证码。</p>
			<p>谢谢！</p>
			<div class="footer">
				<p>此邮件由系统自动发送，请勿回复。</p>
				<p>如果您有任何问题，请联系我们的客服支持。</p>
			</div>
		</div>
	</body>
	</html>
	`, Subject, VerificationCode)

	e.HTML = []byte(htmlContent)
	return e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName.String(), mailPassword.String(), host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}
