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
		Create(ctx context.Context, in entity.User) (string, error)
		// Update 更新用户信息
		Update(ctx context.Context, in entity.User) error
		// Delete 删除用户
		Delete(ctx context.Context, id string) error
		// Detail 获取用户详情
		Detail(ctx context.Context, in model.UserQueryInput) (entity.User, error)
		// GetOne 根据 id 获取用户信息,隐藏关键信息
		GetOne(ctx context.Context, in model.UserQueryInput) (*v1.GetOneRes, error)
		// GetUserInfoByWxId 根据 wxIdOpenId 获取用户信息 todo WxId 后期会合并到 id 中到时删除
		GetUserInfoByWxId(ctx context.Context, wxId string) (*v1.GetOneRes, error)
		// GetUserList 获取用户列表
		GetUserList(ctx context.Context, in v1.GetListReq) ([]entity.User, int, error)
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
