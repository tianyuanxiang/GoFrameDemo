package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 比较特殊，因为原本是一张新表，需要动态添加数据。
// 管理员侧查询图书借阅信息

type BorrowInformationReq struct {
	g.Meta   `path:"/BookBorrowInformation/Query" tags:"BorrowInfromation"  method:"get" summary:"查询图书借阅信息"`
	BookName string `json:"BookName" v:"required"`
	UserIP   string `json:"UserIP"`
	UserName string `json:"UserName"`
}

// 返回flag = 1的所有信息

type BorrowInformationRes struct {
	Message     string           `json:"message"`
	BBorrowData *[]BBinformation `json:"BBorrowInformation"`
}

// 图书信息

type BBinformation struct {
	BookName       string `json:"BookName"`
	ISBN           string `json:"BookISBN"`
	UserIP         string ` json:"UserIP"`
	UserName       string `json:"UserName"`
	BorrowDate     string `json:"BorrowDate"`
	ReturnDate     string `json:"ReturnDate"`
	BorrowingOrder int    `json:"BorrowingOrder"`
}

// 管理员侧查询归还信息接口

type ReturnInformationReq struct {
	g.Meta   `path:"/BookReturnInformation/Query" tags:"BorrowInfromation"  method:"get" summary:"查询图书归还信息"`
	BookName string `json:"BookName" v:"required"`
	UserIP   string `json:"UserIP"`
	UserName string `json:"UserName"`
}

// 返回flag = 1的所有信息

type BReturInformationRes struct {
	Message     string           `json:"message"`
	HistoryDate *[]BBinformation `json:"historyInformation"`
}

// 还书接口
// 在该表上删除该条借阅信息,在图书信息表上的该条记录+1

type ReturnBookReq struct {
	g.Meta         `path:"/BookReturn/Revert" tags:"BookReturn"  method:"get" summary:"还书"`
	BookName       string `json:"BookName" `
	ISBN           string `json:"BookISBN"`
	UserIP         string `json:"UserIP"`
	UserName       string `json:"UserName"`
	BorrowingOrder int    `json:"BorrowingOrder" v:"required"`
}

type ReturnBookRes struct {
	Message            string        `json:"message"`
	ValidateReturnDate BBinformation `json:"ReturnDate"`
}
