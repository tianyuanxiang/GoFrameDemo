package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 图书类别查询接口
type BookTypeReq struct {
	g.Meta     `path:"/bookType/query" method:"get" summary:"查询图书的类别总数"`
	BookTypeID int    `json:"BookTypeID"`
	TypeName   string `json:"TypeName"`
}

type BookTypeRes struct {
	Message    string `json:"message"` // 消息
	ID         int    `json:"ID"`
	BookTypeID int    `json:"booktypeID"`
	TypeName   string `json:"TypeName"`
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
