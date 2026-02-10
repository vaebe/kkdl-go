// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/common/v1"

	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/net/context"
)

type (
	ICommon interface {
		UploadFile(ctx context.Context, file *ghttp.UploadFile) (out v1.UploadFileRes, err error)
		// CreateVCode 创建验证码
		CreateVCode(ctx context.Context, email string) (string, error)
		// VerifyTheVCode 校验验证码
		VerifyTheVCode(ctx context.Context, email string, code string) (bool, error)
		// CheckVCodeCooldown 检查验证码发送冷却期
		CheckVCodeCooldown(ctx context.Context, email string) (inCoolDown bool, remaining int64, err error)
		// ConsumeVCode 消费验证码（标记为已使用）
		ConsumeVCode(ctx context.Context, email string, code string) error
		// DeleteVCode 删除指定邮箱的验证码（用于回滚）
		DeleteVCode(ctx context.Context, email string) error
		// DeleteExpiredVCodes 删除过期的验证码
		DeleteExpiredVCodes(ctx context.Context) error
		// SendEmailVCode 发送邮箱验证码
		SendEmailVCode(ctx context.Context, VerificationCode string, emailAddress string) (err error)
	}
)

var (
	localCommon ICommon
)

func Common() ICommon {
	if localCommon == nil {
		panic("implement not found for interface ICommon, forgot register?")
	}
	return localCommon
}

func RegisterCommon(i ICommon) {
	localCommon = i
}
