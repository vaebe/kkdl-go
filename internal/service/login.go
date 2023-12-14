// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"

	context0 "golang.org/x/net/context"
)

type (
	ILogin interface {
		GetUserInfo(ctx context0.Context, in model.LoginInput) (entity.User, error)
		UserLogin(ctx context0.Context, in model.LoginInput) (userInfo entity.User, token string, tokenExpire string, err error)
		SignOutLogin(ctx context0.Context) error
		GetWeChatToken(ctx context0.Context) (string, error)
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
