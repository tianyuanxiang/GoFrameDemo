// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BookborrowinformationDao is the data access object for table bookborrowinformation.
type BookborrowinformationDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns BookborrowinformationColumns // columns contains all the column names of Table for convenient usage.
}

// BookborrowinformationColumns defines and stores column names for table bookborrowinformation.
type BookborrowinformationColumns struct {
	ID             string //
	BookName       string //
	ISBN           string //
	UserIP         string //
	UserName       string //
	CreatedAt      string //
	ReturnDate     string //
	Flag           string //
	BorrowingOrder string //
}

// bookborrowinformationColumns holds the columns for table bookborrowinformation.
var bookborrowinformationColumns = BookborrowinformationColumns{
	ID:             "ID",
	BookName:       "BookName",
	ISBN:           "ISBN",
	UserIP:         "UserIP",
	UserName:       "UserName",
	CreatedAt:      "created_at",
	ReturnDate:     "ReturnDate",
	Flag:           "Flag",
	BorrowingOrder: "BorrowingOrder",
}

// NewBookborrowinformationDao creates and returns a new DAO object for table data access.
func NewBookborrowinformationDao() *BookborrowinformationDao {
	return &BookborrowinformationDao{
		group:   "default",
		table:   "bookborrowinformation",
		columns: bookborrowinformationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BookborrowinformationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BookborrowinformationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BookborrowinformationDao) Columns() BookborrowinformationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BookborrowinformationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BookborrowinformationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BookborrowinformationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
