package common

import "compressURL/internal/service"

type sCommon struct {
}

func init() {
	service.RegisterCommon(New())
}

func New() *sCommon {
	return &sCommon{}
}
