package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 图书类别查询接口
// 无参数，点击即可

type BookTypeReq struct {
	g.Meta `path:"/bookType/class/query" method:"get" summary:"查询图书的类别总数"`
}

type BookType struct {
	ID         int    `json:"ID"`
	BookTypeID int    `json:"booktypeID"`
	TypeName   string `json:"TypeName"`
}
type BookTypeRes struct {
	Message string      `json:"message"` // 消息
	Type    *[]BookType `json:"bookType"`
}

// 图书信息
type BookTypeInformation struct {
	Id         int    ` dc:"book_id" json:"id"`
	Name       string ` json:"book_name"`
	ISBN       string `json:"Book_ISBN"`
	Author     string `json:"Author"`
	Publishers string `json:"Publishers"`
	BookTypeID int    ` json:"BookTypeID"`
	Amount     int    `json:"Amount"`
}

// 图书类别数量查询接口
type BookTypeNumReq struct {
	g.Meta     `path:"/bookType/number/query" method:"get" summary:"查询该类别图书的总数"`
	BookTypeID int `json:"BookTypeID"`
}

type BookTypeNumRes struct {
	Message     string                 `json:"message"` // 消息
	Information *[]BookTypeInformation `json:"Information"`
}

// 图书类别添加接口
type BookClassAddReq struct {
	g.Meta     `path:"/bookType/Add" method:"post" summary:"添加图书类型"`
	BookTypeID int    `json:"BookTypeID"`
	TypeName   string `json:"TypeName"`
}

type BookClassAddRes struct {
	Message    string `json:"message"` // 返回是否插入成功，成功则返回插入的信息，失败则返回状态码
	ID         int    `json:"ID"`
	BookTypeID int    `json:"booktypeID"`
	TypeName   string `json:"TypeName"`
}

// 图书类别名修改接口
type BookTypeModifyReq struct {
	g.Meta     `path:"/bookType/Modify" method:"get" summary:"修改图书类型名称"`
	BookTypeID int    `json:"BookTypeID"`
	TypeName   string `json:"TypeName"` // 如果啥也没输入，原有数据会被覆盖吗？
}
type BookTypeModifyRes struct {
	Message    string `json:"message"` // 返回是否插入成功，成功则返回插入的信息，失败则返回状态码
	ID         int    `json:"ID"`
	BookTypeID int    `json:"booktypeID"`
	TypeName   string `json:"TypeName"`
}
