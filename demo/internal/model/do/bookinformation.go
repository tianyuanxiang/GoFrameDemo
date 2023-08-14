// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Bookinformation is the golang structure of table bookinformation for DAO operations like Where/Data.
type Bookinformation struct {
	g.Meta     `orm:"table:bookinformation, do:true"`
	ID         interface{} //
	BookName   interface{} //
	ISBN       interface{} //
	Author     interface{} //
	Publishers interface{} //
	BookTypeID interface{} //
	Amount     interface{} //
}
