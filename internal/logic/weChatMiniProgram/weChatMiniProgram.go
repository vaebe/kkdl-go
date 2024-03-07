package weChatMiniProgram

import "compressURL/internal/service"

type sWeChatMiniProgram struct {
}

func init() {
	service.RegisterWeChatMiniProgram(New())
}

func New() *sWeChatMiniProgram {
	return &sWeChatMiniProgram{}
}
