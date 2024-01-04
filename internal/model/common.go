package model

// PageParams 分页参数
type PageParams struct {
	Total    int `json:"total" `
	PageSize int `json:"pageSize" v:"required#请输入 PageSize" dc:"分页条数" d:"1"`
	PageNo   int `json:"pageNo"  v:"required#请输入用户 PageNo" dc:"当前页" d:"10"`
}
