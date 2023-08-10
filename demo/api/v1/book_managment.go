package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api -> controller -> logic
// 在数据库中查询，并返回出来

// 查询
// 描述了客户端在调用接口时所需提供的数据格式和字段(调用接口需要提供的东西)。
type BookQueryReq struct {
	g.Meta `path:"/book_managment/Query" tags:"BookQueryReq" method:"get" summary:"查询图书信息"`
	Name   string `v:"require" json:"BookName"`
	ISBN   string `json:"Book_ISBN"`
}

type BookQueryRes struct {
	// mime	接口的MIME类型，例如multipart/form-data一般是全局设置，默认为application/json。
	g.Meta      `mime:"text/html" example:"string"`
	Message     string            `json:"message"`
	Information []BookInformation `json:"Information"`
	Flag        bool              `json:"IsOrNullExisit"`
}

// 图书信息
type BookInformation struct {
	Id         int    ` dc:"book_id" json:"id"`
	Name       string ` json:"book_name"`
	ISBN       string `json:"Book_ISBN"`
	Author     string `json:"Author"`
	Publishers string `json:"Publishers"`
	BookTypeID int    ` json:"BookTypeID"`
	Amount     int    `json:"Amount"`
}

// 新增
type BookInsertReq struct {
	g.Meta `path:"/book_managment/Insert" tags:"Insert_Data" method:"post" summary:"添加一条图书信息"`
	Date   BookInformation `json:"InsertInformation"`
}
type BookInsertRes struct {
	Message string          `json:"Message"`
	Date    BookInformation `json:"date"`
}

// 修改
// 说明里面已经存在了，我需要先查询出来，再修改
type BookUpdateReq struct {
	g.Meta      `path:"/book_managment/Update" tags:"Update_Data" method:"get" summary:"修改图书信息"`
	Information BookInformation `json:"Book_Information"`
	Ret         int             `json:"updateID"`
}
type BookUpdateRes struct {
	Message string          `json:"message"`
	Date    BookInformation `json:"Updated_information"`
}

// 删除
// 只需要提供其中一个字段（name或者ISBN）
type BookDeleteReq struct {
	g.Meta `path:"/book_managment/Delete" tags:"DeleteData" method:"get" summary:"删除一条图书信息"`
	// 假设数据表中Name、ISBN和publisher_id无重复
	Name string `json:"bookName"`
	ISBN string `json:"bookISBN"`
}

// 删除是指对该条图书信息直接执行删除，不论库存，返回被删除的图书信息
type BookDeleteRes struct {
	Message     string            `json:"message"`
	Information []BookInformation `json:"DeleteInformation"`
}
