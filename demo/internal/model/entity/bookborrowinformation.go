// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Bookborrowinformation is the golang structure for table bookborrowinformation.
type Bookborrowinformation struct {
	ID         int         `json:"iD"         ` //
	BookName   string      `json:"bookName"   ` //
	ISBN       string      `json:"iSBN"       ` //
	UserIP     string      `json:"userIP"     ` //
	UserName   string      `json:"userName"   ` //
	BorrowDate *gtime.Time `json:"borrowDate" ` //
	ReturnDate *gtime.Time `json:"returnDate" ` //
	Flag       int         `json:"flag"       ` //
}
