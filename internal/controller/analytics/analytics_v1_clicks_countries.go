package analytics

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"compressURL/api/analytics/v1"
)

func (c *ControllerV1) ClicksCountries(ctx context.Context, req *v1.ClicksCountriesReq) (res *v1.ClicksCountriesRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
