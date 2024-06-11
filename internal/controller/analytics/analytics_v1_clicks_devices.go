package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksDevices(ctx context.Context, req *v1.ClicksDevicesReq) (res *v1.ClicksDevicesRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	list, err := service.Analytics().GetVisitsByDevice(ctx, *req, loginUserInfo.Id)
	if err != nil {
		return res, err
	}

	if list == nil {
		list = v1.ClicksDevicesRes{}
	}

	return &list, nil
}
