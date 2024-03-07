// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	ILogin interface {
		// GetUserInfo 获取用户信息
		GetUserInfo(ctx context.Context, in model.LoginInput) (entity.User, error)
		// UserLogin 用户登录
		UserLogin(ctx context.Context, in model.LoginInput) (userInfo entity.User, token string, tokenExpire string, err error)
		// SignOutLogin 退出登录
		SignOutLogin(ctx context.Context) error
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
