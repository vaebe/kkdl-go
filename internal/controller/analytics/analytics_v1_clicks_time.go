package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksTime(ctx context.Context, req *v1.ClicksTimeReq) (res *v1.ClicksTimeRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	list, err := service.Analytics().GetVisitsByDate(ctx, *req, loginUserInfo.Id)
	if err != nil {
		return nil, err
	}

	if list == nil {
		list = v1.ClicksTimeRes{}
	}

	return &list, err
}
