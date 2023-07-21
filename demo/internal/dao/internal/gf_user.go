// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================
// 据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 `CURD` 方法
package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GfUserDao is the data access object for table gf_user.
type GfUserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns GfUserColumns // columns contains all the column names of Table for convenient usage.
}

// GfUserColumns defines and stores column names for table gf_user.
type GfUserColumns struct {
	Id        string // UID
	Passport  string // 账号
	Password  string // MD5密码
	Nickname  string // 昵称
	Avatar    string // 头像地址
	Status    string // 状态 0:启用 1:禁用
	Gender    string // 性别 0: 未设置 1: 男 2: 女
	CreatedAt string // 注册时间
	UpdatedAt string // 更新时间
}

// gfUserColumns holds the columns for table gf_user.
var gfUserColumns = GfUserColumns{
	Id:        "id",
	Passport:  "passport",
	Password:  "password",
	Nickname:  "nickname",
	Avatar:    "avatar",
	Status:    "status",
	Gender:    "gender",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewGfUserDao creates and returns a new DAO object for table data access.
func NewGfUserDao() *GfUserDao {
	return &GfUserDao{
		group:   "default",
		table:   "gf_user",
		columns: gfUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GfUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GfUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GfUserDao) Columns() GfUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GfUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GfUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GfUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
