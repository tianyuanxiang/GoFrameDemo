package v1

import "github.com/gogf/gf/v2/frame/g"

// 图书借阅信息查询接口
// 查看当前用户在借阅信息表上的借阅信息

type UserBookBorrowReq struct {
	g.Meta   `path:"/UserBorrowedReq/Query" tags:"UserBorrowReq"  method:"get" summary:"用户查询借阅的图书信息"`
	UserIP   string `json:"UserIP" v:"required"`
	UserName string `json:"UserName"`
}

type UserBookBorrowRes struct {
	Message          string              `json:"message"`
	UBookBorrowGroup *[]UBookBorrowGroup `json:"UBookBorrowGroup"`
}

type UBookBorrowGroup struct {
	ID             int    `json:"ID"`
	BookName       string `json:"BookName"`
	ISBN           string `json:"ISBN"`
	BorrowTime     string `json:"BorrowTime"`
	ReturnTime     string `json:"ReturnTime"`
	BorrowingOrder int    `json:"BorrowingOrder"`
}

// 用户侧的还书接口
// 与管理员侧共用一个还书接口
