package v1

import "github.com/gogf/gf/v2/frame/g"

// 用户侧查询图书信息接口
// 与管理员侧的图书查询功能共用一个接口

// 还得再加一个借阅订单号，用来区分同一用户借阅同一本书的区别

// 用户侧查询图书信息接口
// 与管理员侧共用一个查询图书信息接口

// 用户侧借阅图书接口
type UBookBorrowReq struct {
	g.Meta            `path:"/UBookQuery/Borrow" tags:"BookBorrow"  method:"get" summary:"用户借阅图书"`
	BorrowInformation UBBInformation `json:"BorrowInformation"`
}

type UBookBorrowRes struct {
	Message     string         `json:"message"`
	Information UBBInformation `json:"Information"`
}

// 图书信息
type UBBInformation struct {
	BookName string `json:"BookName"`
	ISBN     string `json:"ISBN"`
	UserIP   string `json:"UserIP"`
	UserName string `json:"UserName"`
}
