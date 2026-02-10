package common

import (
	"compressURL/internal/dao"
	"compressURL/internal/model/do"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/jordan-wright/email"
	"golang.org/x/net/context"
)

// CreateVCode 创建验证码
func (s *sCommon) CreateVCode(ctx context.Context, email string) (string, error) {
	// 生成6位随机验证码
	code := grand.N(100000, 999999)
	codeStr := gconv.String(code)

	// 保存到数据库
	_, err := dao.VerificationCodes.Ctx(ctx).Data(do.VerificationCodes{
		Email:     email,
		Type:      1,
		Code:      codeStr,
		Used:      0,
		ExpiredAt: gtime.Now().Add(10 * time.Minute),
	}).Insert()

	return codeStr, err
}

// VerifyTheVCode 校验验证码
func (s *sCommon) VerifyTheVCode(ctx context.Context, email string, code string) (bool, error) {
	count, err := dao.VerificationCodes.Ctx(ctx).Where(do.VerificationCodes{
		Email: email,
		Code:  code,
		Used:  0,
	}).Where("expired_at > ?", gtime.Now()).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckVCodeCooldown 检查验证码发送冷却期
func (s *sCommon) CheckVCodeCooldown(ctx context.Context, email string) (inCoolDown bool, remaining int64, err error) {
	// 查找最近创建的验证码
	record, err := dao.VerificationCodes.Ctx(ctx).Where(do.VerificationCodes{Email: email}).OrderDesc("created_at").One()
	if err != nil {
		return false, 0, err
	}

	if record.IsEmpty() {
		return false, 0, nil
	}

	createdAt := record["created_at"].Time()
	elapsed := time.Since(createdAt)

	// 如果创建时间小于1分钟，则在冷却期
	if elapsed < 60*time.Second {
		remaining = 60 - int64(elapsed.Seconds())
		return true, remaining, nil
	}

	return false, 0, nil
}

// ConsumeVCode 消费验证码（标记为已使用）
func (s *sCommon) ConsumeVCode(ctx context.Context, email string, code string) error {
	_, err := dao.VerificationCodes.Ctx(ctx).Where(do.VerificationCodes{
		Email: email,
		Code:  code,
	}).Update(do.VerificationCodes{Used: 1})
	return err
}

// DeleteExpiredVCodes 删除过期的验证码
func (s *sCommon) DeleteExpiredVCodes(ctx context.Context) error {
	_, err := dao.VerificationCodes.Ctx(ctx).Where("expired_at < ?", gtime.Now()).Delete()
	return err
}

// SendEmailVCode 发送邮箱验证码
func (s *sCommon) SendEmailVCode(ctx context.Context, VerificationCode int, emailAddress string) (err error) {
	mailUserName, _ := g.Cfg().Get(ctx, "emailConfig.email") // 邮箱账号
	mailPassword, _ := g.Cfg().Get(ctx, "emailConfig.key")   // 邮箱授权码
	addr := "smtp.qq.com:465"                                // TLS 地址
	host := "smtp.qq.com"                                    // 邮件服务器地址
	Subject := "KKDL 验证码"                                    // 发送的主题

	e := email.NewEmail()
	e.From = fmt.Sprintf("KKDL <%s>", mailUserName)
	e.To = []string{emailAddress}
	e.Subject = Subject

	// 美化后的 HTML 内容
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
