// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Library is the golang structure for table library.
type Library struct {
	Id          int         `json:"id"          ` //
	Name        string      `json:"name"        ` //
	ISBN        string      `json:"iSBN"        ` //
	Translator  string      `json:"translator"  ` //
	Date        *gtime.Time `json:"date"        ` //
	PublisherId int         `json:"publisherId" ` //
}
