// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/model/entity"

	"golang.org/x/net/context"
)

type (
	IUser interface {
		Create(ctx context.Context, in entity.User) (string, error)
		// Update 更新用户信息
		Update(ctx context.Context, in entity.User) error
		// Remove 删除用户
		Remove(ctx context.Context, id string) error
		GetUserInfo(ctx context.Context, id string) (*v1.GetUserInfoRes, error)
		// Detail 获取用户详情
		Detail(ctx context.Context, id string) (entity.User, error)
		// GetUserList 获取用户列表
		GetUserList(ctx context.Context, in v1.GetUserListReq) ([]entity.User, int, error)
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
