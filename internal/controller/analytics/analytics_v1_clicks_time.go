package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksTime(ctx context.Context, req *v1.ClicksTimeReq) (res *v1.ClicksTimeRes, err error) {
	date, err := service.Analytics().GetVisitsByDate(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &date, err
}
