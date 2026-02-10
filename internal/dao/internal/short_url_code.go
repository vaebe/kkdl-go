// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortUrlCodeDao is the data access object for the table short_url_code.
type ShortUrlCodeDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ShortUrlCodeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ShortUrlCodeColumns defines and stores column names for the table short_url_code.
type ShortUrlCodeColumns struct {
	Id        string // 唯一标识，自增长整数
	Code      string // 短链,唯一，不能为空
	Status    string // 是否使用
	CreatedAt string // 创建时间，默认为当前时间戳
}

// shortUrlCodeColumns holds the columns for the table short_url_code.
var shortUrlCodeColumns = ShortUrlCodeColumns{
	Id:        "id",
	Code:      "code",
	Status:    "status",
	CreatedAt: "created_at",
}

// NewShortUrlCodeDao creates and returns a new DAO object for table data access.
func NewShortUrlCodeDao(handlers ...gdb.ModelHandler) *ShortUrlCodeDao {
	return &ShortUrlCodeDao{
		group:    "default",
		table:    "short_url_code",
		columns:  shortUrlCodeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShortUrlCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShortUrlCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShortUrlCodeDao) Columns() ShortUrlCodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShortUrlCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShortUrlCodeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ShortUrlCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
