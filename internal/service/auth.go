// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"compressURL/internal/model"

	jwt "github.com/gogf/gf-jwt/v2"
	"golang.org/x/net/context"
)

type (
	IAuth interface {
		AuthInstance() *jwt.GfJWTMiddleware
		// GetLoginUserInfo 获取当前登录用户的信息
		GetLoginUserInfo(ctx context.Context) (model.JWTPayloadInfo, error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
