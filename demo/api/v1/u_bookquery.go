package v1

import "github.com/gogf/gf/v2/frame/g"

// 用户侧查询图书借阅信息
type UBookQueryReq struct {
	g.Meta   `path:"/uBookQuery/Query" tags:"BookQuery"  method:"get" summary:"用户侧查询图书信息"`
	BookName string `json:"BookName"`
	ISBN     string `json:"ISBN"`
}

type UBookQueryRes struct {
	Message     string              `json:"message"`
	Information *[]UBookInformation `json:"Information"`
}

// 图书信息
type UBookInformation struct {
	Id         int    ` dc:"book_id" json:"id"`
	Name       string ` json:"book_name"`
	ISBN       string `json:"Book_ISBN"`
	Author     string `json:"Author"`
	Publishers string `json:"Publishers"`
	BookTypeID int    ` json:"BookTypeID"`
	Amount     int    `json:"Amount"`
}
