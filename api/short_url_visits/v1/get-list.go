package v1

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/shortUrlVisits/getList" method:"post" summary:"获取短链访问信息列表" tags:"短链访问信息"`
	model.PageParams
	Code string `json:"code"   dc:"短链"`
}

type GetListRes struct {
	List []entity.ShortUrlVisits `json:"list" dc:"短链访问数据"`
	model.PageParams
}
