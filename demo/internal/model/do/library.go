// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Library is the golang structure of table library for DAO operations like Where/Data.
type Library struct {
	g.Meta      `orm:"table:library, do:true"`
	Id          interface{} //
	Name        interface{} //
	ISBN        interface{} //
	Translator  interface{} //
	Date        *gtime.Time //
	PublisherId interface{} //
}
