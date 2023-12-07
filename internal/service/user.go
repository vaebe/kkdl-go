// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/model"

	"golang.org/x/net/context"
)

type (
	IUser interface {
		Create(ctx context.Context, in model.UserCreateInput) (userid int64, error error)
		Remove(ctx context.Context, id int64) error
		Update(ctx context.Context, in model.UserUpdateInput) error
		GetUserInfo(ctx context.Context, id int64) (info *v1.GetUserInfoRes, error error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
