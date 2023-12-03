// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	UserId       string // 用户唯一标识，自增长整数
	Username     string // 用户名，唯一，不能为空
	Email        string // 用户的电子邮件地址，唯一，不能为空
	PasswordHash string // 存储密码的哈希值，不能为空
	FullName     string // 用户的全名
	CreatedAt    string // 用户账号创建时间，默认为当前时间戳
	IsAdmin      string // 表示用户是否是管理员，默认为 FALSE
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	UserId:       "user_id",
	Username:     "username",
	Email:        "email",
	PasswordHash: "password_hash",
	FullName:     "full_name",
	CreatedAt:    "created_at",
	IsAdmin:      "is_admin",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
