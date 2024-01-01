// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	IUser interface {
		Create(ctx context.Context, in model.UserCreateInput) (userid string, error error)
		Remove(ctx context.Context, id string) error
		Update(ctx context.Context, in model.UserUpdateInput) error
		GetUserInfo(ctx context.Context, id string) (info *v1.GetUserInfoRes, error error)
		// Detail 获取用户详情
		Detail(ctx context.Context, id string) (info entity.User, error error)
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
