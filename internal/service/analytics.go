// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "compressURL/api/analytics/v1"

	"golang.org/x/net/context"
)

type (
	IAnalytics interface {
		// GetVisitsByDevice 根据设备统计访问数据
		GetVisitsByDevice(ctx context.Context, req v1.ClicksDevicesReq, userId string) (list v1.ClicksDevicesRes, err error)
		// GetVisitsByRegion 根据区域统计访问数据
		GetVisitsByRegion(ctx context.Context, req v1.ClicksRegionReq, userId string) (list v1.ClicksRegionRes, err error)
		// GetVisitsByDate 根据时间统计访问数据
		GetVisitsByDate(ctx context.Context, req v1.ClicksTimeReq, userId string) (list v1.ClicksTimeRes, err error)
	}
)

var (
	localAnalytics IAnalytics
)

func Analytics() IAnalytics {
	if localAnalytics == nil {
		panic("implement not found for interface IAnalytics, forgot register?")
	}
	return localAnalytics
}

func RegisterAnalytics(i IAnalytics) {
	localAnalytics = i
}
