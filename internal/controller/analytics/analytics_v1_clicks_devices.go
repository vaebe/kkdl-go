package analytics

import (
	"compressURL/api/analytics/v1"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) ClicksDevices(ctx context.Context, req *v1.ClicksDevicesReq) (res *v1.ClicksDevicesRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)

	// 获取到用户id 就赋值,没有就默认查询全部数据
	userId := ""
	if loginUserInfo.Id != "" {
		userId = loginUserInfo.Id
	}

	list, err := service.Analytics().GetVisitsByDevice(ctx, *req, userId)
	if err != nil {
		return res, err
	}

	if list == nil {
		list = v1.ClicksDevicesRes{}
	}

	return &list, nil
}
