package analytics

import "compressURL/internal/service"

type sAnalytics struct {
}

func init() {
	service.RegisterAnalytics(New())
}

func New() *sAnalytics { return &sAnalytics{} }
