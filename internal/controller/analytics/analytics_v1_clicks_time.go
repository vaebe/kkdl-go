package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksTime(ctx context.Context, req *v1.ClicksTimeReq) (res *v1.ClicksTimeRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)

	// 获取到用户id 就赋值,没有就默认查询全部数据
	userId := ""
	if loginUserInfo.Id != "" {
		userId = loginUserInfo.Id
	}

	list, err := service.Analytics().GetVisitsByDate(ctx, *req, userId)
	if err != nil {
		return nil, err
	}

	if list == nil {
		list = v1.ClicksTimeRes{}
	}

	return &list, err
}
