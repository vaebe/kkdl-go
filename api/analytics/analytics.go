// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package analytics

import (
	"context"

	"compressURL/api/analytics/v1"
)

type IAnalyticsV1 interface {
	ClicksTime(ctx context.Context, req *v1.ClicksTimeReq) (res *v1.ClicksTimeRes, err error)
	ClicksCountries(ctx context.Context, req *v1.ClicksCountriesReq) (res *v1.ClicksCountriesRes, err error)
}
