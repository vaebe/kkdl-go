package v1

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/shortUrl/getList" method:"post" summary:"获取短链列表" tags:"短链"`
	model.PageParams
	Title  string `json:"title"   dc:"短链标题"`
	RawUrl string `json:"rawUrl"   dc:"跳转链接"`
}
type GetListRes struct {
	List []entity.ShortUrl `json:"list" dc:"短链数据"`
	model.PageParams
}
