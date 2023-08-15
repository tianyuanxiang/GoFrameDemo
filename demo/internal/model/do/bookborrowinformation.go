// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Bookborrowinformation is the golang structure of table bookborrowinformation for DAO operations like Where/Data.
type Bookborrowinformation struct {
	g.Meta         `orm:"table:bookborrowinformation, do:true"`
	ID             interface{} //
	BookName       interface{} //
	ISBN           interface{} //
	UserIP         interface{} //
	UserName       interface{} //
	CreatedAt      *gtime.Time //
	ReturnDate     *gtime.Time //
	Flag           interface{} //
	BorrowingOrder interface{} //
}
