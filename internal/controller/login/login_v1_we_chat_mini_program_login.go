package login

import (
	"compressURL/internal/service"
	"context"

	"compressURL/api/login/v1"
)

func (c *ControllerV1) WeChatMiniProgramLogin(ctx context.Context, req *v1.WeChatMiniProgramLoginReq) (res *v1.WeChatMiniProgramLoginRes, err error) {
	token, err := service.Login().GetWeChatToken(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.WeChatMiniProgramLoginRes{Token: token}, nil
}
