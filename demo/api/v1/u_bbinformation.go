package v1

import "github.com/gogf/gf/v2/frame/g"

// 图书查询接口
type UserBookReq struct {
	g.Meta   `path:"/UserBookReq/Query" tags:"UserBookReq"  method:"get" summary:"用户查询图书信息"`
	BookName string `json:"BookName" v:"required"`
	UserIP   string `json:"UserIP"`
	UserName string
}
