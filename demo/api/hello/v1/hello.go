package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api -> controller -> logic
// 在数据库中查询，并返回出来

// 查询
// 描述了客户端在调用接口时所需提供的数据格式和字段(调用接口需要提供的东西)。
type Q_Req struct {
	g.Meta      `path:"/Query" tags:"Query_infromation" method:"get" summary:"查询图书信息"`
	Name        string `v:"require" json:"book_name"`
	ISBN        string `json:"Book_ISBN"`
	PublisherId int    `json:"PublisherId"`
}

type Q_Res struct {
	// mime	接口的MIME类型，例如multipart/form-data一般是全局设置，默认为application/json。
	g.Meta  `mime:"text/html" example:"string"`
	Message string            `json:"message"`
	Date    []BookInformation `json:"date"`
	Flag    bool              `json:"IsOrNullExisit"`
}

// 图书信息
type BookInformation struct {
	Id           int    ` dc:"book_id" json:"id"`
	Name         string ` json:"book_name"`
	ISBN         string `json:"Book_ISBN"`
	Translator   string `json:"translator"`
	Date         string ` json:"publish_date"`
	Publisher_id int    `json:"publisher_Id"`
}

// 新增
type InsertReq struct {
	g.Meta `path:"/Insert" tags:"Insert_Data" method:"post" summary:"添加一条图书信息"`
	Date   BookInformation `json:"InsertInformation"`
}
type InsertRes struct {
	Message string          `json:"Message"`
	Date    BookInformation `json:"date"`
}

// 修改
// 说明里面已经存在了，我需要先查询出来，再修改
type UpdateReq struct {
	g.Meta      `path:"/Update" tags:"Update_Data" method:"get" summary:"修改图书信息"`
	Information BookInformation `json:"Book_Information"`
}
type UpdateRes struct {
	Message string          `json:"message"`
	Date    BookInformation `json:"Updated_information"`
}

// 删除
// 只需要提供其中一个字段（name或者ISBN）
type DeleteReq struct {
	g.Meta `path:"/Delete" tags:"DeleteData" method:"get" summary:"删除一条图书信息"`
	// 假设数据表中Name、ISBN和publisher_id无重复
	Name        string `json:"bookName"`
	ISBN        string `json:"bookISBN"`
	PublisherID int    `json:"PublisherID"`
}
type DeleteRes struct {
	Date    []BookInformation `json:"DeleteInformation"`
	Message string            `json:"message"`
}
