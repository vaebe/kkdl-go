package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksCountries(ctx context.Context, req *v1.ClicksCountriesReq) (res *v1.ClicksCountriesRes, err error) {
	list, err := service.Analytics().GetVisitsByCountries(ctx, *req)
	if err != nil {
		return nil, err
	}
	return &list, nil
}
