package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksRegion(ctx context.Context, req *v1.ClicksRegionReq) (res *v1.ClicksRegionRes, err error) {
	list, err := service.Analytics().GetVisitsByRegion(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &list, nil
}
