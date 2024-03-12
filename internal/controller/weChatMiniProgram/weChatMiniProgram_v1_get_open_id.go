package weChatMiniProgram

import (
	"compressURL/internal/consts"
	"compressURL/internal/service"
	"compressURL/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/weChatMiniProgram/v1"
)

func (c *ControllerV1) GetOpenId(ctx context.Context, req *v1.GetOpenIdReq) (res *v1.GetOpenIdRes, err error) {
	info, err := service.WeChatMiniProgram().GetOpenId(ctx, req.Code)

	err = g.Redis().SetEX(ctx, utility.GetWeChatMiniProgramLoginCode(req.UserCode), consts.WeChatMiniProgramLoginStatusSuccess, 60*10)
	if err != nil {
		return nil, gerror.New("redis 缓存小程序用户登录状态失败!")
	}

	return info, err
}
