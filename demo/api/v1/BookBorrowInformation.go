package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 比较特殊，因为原本是一张新表，需要动态添加数据。
// 查询借阅信息和归还信息，都通过这一个接口完成
type BorrowInformationReq struct {
	g.Meta   `path:"/BookBorrowInformation/Query" tags:"BorrowInfromation"  method:"get" summary:"查询图书借阅信息"`
	BookName string `json:"BookName" v:"required"`
	UserIP   string `json:"UserIP"`
	UserName string
}

// 返回flag = 1的所有信息
type BorrowInformationRes struct {
	Message     string           `json:"message"`
	HistoryDate []*BBinformation `json:"historyInformation"`
}

// 图书信息
type BBinformation struct {
	Id         int    ` dc:"bookId" json:"ID"`
	ISBN       string `json:"BookISBN"`
	BookName   string `json:"BookName"`
	UserIP     string ` json:"UserIP"`
	UserName   string `json:"UserName"`
	BorrowDate string `json:"BorrowDate"`
	ReturnDate string `json:"ReturnDate"`
}

// 查询归还信息接口
type ReturnInformationReq struct {
	g.Meta   `path:"/BookBorrowInformation/Query" tags:"BorrowInfromation"  method:"get" summary:"查询图书借阅信息"`
	BookName string `json:"BookName" v:"required"`
	UserIP   string `json:"UserIP"`
	UserName string `json:"UserName"`
}

// 返回flag = 1的所有信息
type BReturInformationRes struct {
	Message     string           `json:"message"`
	HistoryDate []*BBinformation `json:"historyInformation"`
}

// 还书接口
// 在该表上删除该条借阅信息,在图书信息表上的该条记录+1
type ReturnBookReq struct {
	BookName string `json:"BookName" v:"required"`
	ISBN     string `json:"BookISBN"`
}

type ReturnBookRes struct {
	Message    string           `json:"message"`
	ReturnDate []*BBinformation `json:"ReturnDate"`
}
