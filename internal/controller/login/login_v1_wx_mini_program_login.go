package login

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) WxMiniProgramLogin(ctx context.Context, req *v1.WxMiniProgramLoginReq) (res *v1.WxMiniProgramLoginRes, err error) {
	token, err := service.WeChatMiniProgram().GetWeChatToken(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.WxMiniProgramLoginRes{Token: token}, nil
}
