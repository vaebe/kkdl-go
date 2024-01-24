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
