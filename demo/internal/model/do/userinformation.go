// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Userinformation is the golang structure of table userinformation for DAO operations like Where/Data.
type Userinformation struct {
	g.Meta     `orm:"table:userinformation, do:true"`
	ID         interface{} //
	UserIP     interface{} //
	UserName   interface{} //
	Email      interface{} //
	CurrentNum interface{} //
	HistoryNum interface{} //
}
