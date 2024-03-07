// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/weChatMiniProgram/v1"

	"golang.org/x/net/context"
)

type (
	IWeChatMiniProgram interface {
		// GetOpenId 获取微信小程序用户openId
		GetOpenId(ctx context.Context, Code string) (resInfo *v1.GetWeChatMiniProgramOpenIdRes, err error)
		// GetWeChatToken 获取微信 api token
		GetWeChatToken(ctx context.Context) (string, error)
	}
)

var (
	localWeChatMiniProgram IWeChatMiniProgram
)

func WeChatMiniProgram() IWeChatMiniProgram {
	if localWeChatMiniProgram == nil {
		panic("implement not found for interface IWeChatMiniProgram, forgot register?")
	}
	return localWeChatMiniProgram
}

func RegisterWeChatMiniProgram(i IWeChatMiniProgram) {
	localWeChatMiniProgram = i
}
