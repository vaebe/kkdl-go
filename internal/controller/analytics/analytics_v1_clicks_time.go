package analytics

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksTime(ctx context.Context, req *v1.ClicksTimeReq) (res *v1.ClicksTimeRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)

	userId := ""
	// 普通用户只能查询自己的数据, 未登录和管理员查询全部数据
	if loginUserInfo.Id != "" && loginUserInfo.Role != "00" {
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
