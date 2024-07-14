package analytics

import (
	"compressURL/api/analytics/v1"
	"compressURL/internal/service"
	"context"
)

func (c *ControllerV1) ClicksDevices(ctx context.Context, req *v1.ClicksDevicesReq) (res *v1.ClicksDevicesRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)

	userId := ""
	// 普通用户只能查询自己的数据, 未登录和管理员查询全部数据
	if loginUserInfo.Id != "" && loginUserInfo.Role != "00" {
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
