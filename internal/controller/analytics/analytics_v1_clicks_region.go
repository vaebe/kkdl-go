package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksRegion(ctx context.Context, req *v1.ClicksRegionReq) (res *v1.ClicksRegionRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)

	// 获取到用户id 就赋值,没有就默认查询全部数据
	userId := ""
	if loginUserInfo.Id != "" {
		userId = loginUserInfo.Id
	}

	list, err := service.Analytics().GetVisitsByRegion(ctx, *req, userId)
	if err != nil {
		return nil, err
	}

	if list == nil {
		list = v1.ClicksRegionRes{}
	}

	return &list, nil
}
