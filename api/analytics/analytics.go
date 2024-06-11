// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package analytics

import (
	"context"

	"compressURL/api/analytics/v1"
)

type IAnalyticsV1 interface {
	ClicksDevices(ctx context.Context, req *v1.ClicksDevicesReq) (res *v1.ClicksDevicesRes, err error)
	ClicksRegion(ctx context.Context, req *v1.ClicksRegionReq) (res *v1.ClicksRegionRes, err error)
	ClicksTime(ctx context.Context, req *v1.ClicksTimeReq) (res *v1.ClicksTimeRes, err error)
}
