package v1

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/user/getList" method:"post" summary:"获取用户列表" tags:"用户"`
	model.PageParams
	NickName string `json:"nickName"   dc:"昵称"`
	Email    string `json:"email"   dc:"邮箱"`
	WxId     string `json:"wxId"   dc:"微信 id"`
}
type GetListRes struct {
	List []entity.User `json:"list" dc:"用户数据"`
	model.PageParams
}
