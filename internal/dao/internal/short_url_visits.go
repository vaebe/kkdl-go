// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortUrlVisitsDao is the data access object for the table short_url_visits.
type ShortUrlVisitsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  ShortUrlVisitsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// ShortUrlVisitsColumns defines and stores column names for the table short_url_visits.
type ShortUrlVisitsColumns struct {
	Id             string //
	UserId         string // 用户id
	ShortUrl       string // 短链,唯一，不能为空
	RawUrl         string // 原始 url 不能为空
	UserAgent      string // 用户代理字符串，存储提供的完整用户代理
	BrowserName    string // 浏览器名称
	BrowserVersion string // 浏览器版本
	DeviceModel    string // 设备型号
	EngineName     string // 浏览器引擎名称
	EngineVersion  string // 浏览器引擎版本
	OsName         string // 操作系统名称
	OsVersion      string // 操作系统版本
	Ip             string // ip 不能为空
	Continent      string // 大洲名称
	ContinentCode  string // 大洲代码
	Country        string // 国家名称
	CountryCode    string // 国家代码
	Region         string // 地区或州的短代码（FIPS或ISO）
	RegionName     string // 地区或州名称
	City           string // 城市名称
	District       string // 位置的区（郡）
	Lat            string // 纬度
	Lon            string // 经度
	CreatedAt      string // 创建时间，默认为当前时间戳
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
}

// shortUrlVisitsColumns holds the columns for the table short_url_visits.
var shortUrlVisitsColumns = ShortUrlVisitsColumns{
	Id:             "id",
	UserId:         "user_id",
	ShortUrl:       "short_url",
	RawUrl:         "raw_url",
	UserAgent:      "user_agent",
	BrowserName:    "browser_name",
	BrowserVersion: "browser_version",
	DeviceModel:    "device_model",
	EngineName:     "engine_name",
	EngineVersion:  "engine_version",
	OsName:         "os_name",
	OsVersion:      "os_version",
	Ip:             "ip",
	Continent:      "continent",
	ContinentCode:  "continent_code",
	Country:        "country",
	CountryCode:    "country_code",
	Region:         "region",
	RegionName:     "region_name",
	City:           "city",
	District:       "district",
	Lat:            "lat",
	Lon:            "lon",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewShortUrlVisitsDao creates and returns a new DAO object for table data access.
func NewShortUrlVisitsDao(handlers ...gdb.ModelHandler) *ShortUrlVisitsDao {
	return &ShortUrlVisitsDao{
		group:    "default",
		table:    "short_url_visits",
		columns:  shortUrlVisitsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShortUrlVisitsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShortUrlVisitsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShortUrlVisitsDao) Columns() ShortUrlVisitsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShortUrlVisitsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShortUrlVisitsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ShortUrlVisitsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
