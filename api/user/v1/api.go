package v1

type CommonPaginationReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

type CommonPaginationRes struct {
	Total int `dc:"总数"`
}

type ContentGetListCommonReq struct {
	Type       string `json:"type"   in:"query" dc:"内容模型"`
	CategoryId uint   `json:"cate"   in:"query" dc:"栏目ID"`
	Sort       int    `json:"sort"   in:"query" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`
	CommonPaginationReq
}
