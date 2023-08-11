package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 读者信息查询接口
type ReaderInformationReq struct {
	g.Meta   `path:"/ReaderInformationReq/Query" tags:"ReaderInformation"  method:"get" summary:"查询读者信息"`
	ID       int    `json:"ReaderID"`
	UserIP   string `json:"UserIP"`
	UserName string `json:"UserName" v:"required"`
}

// 直接返回完整的读者信息，当前借阅数和历史借阅数怎么搞？
type ReaderInformationRes struct {
	Message string             `json:"message"`
	User    *[]UserInformation `json:"userInformation"`
}

type UserInformation struct {
	Id         int    ` dc:"UserId" json:"ID"`
	UserIP     string ` json:"UserIP"`
	UserName   string `json:"UserName"`
	Email      string `json:"Email"`
	CurrentNum int    `json:"CurrentNum"`
	HistoryNum int    `json:"HistoryNum"`
}

// 读者信息修改接口
// 如果改了“读者管理表”中的读者名字，那么在“借阅信息表”中该读者的名字<在查询的时候>也要随之而变。
// 需要通过一个中间表来构建多对多关系。
type ReaderModifyReq struct {
	g.Meta     `path:"/ReaderModifyReq/Modify" tags:"ReaderModifyReq"  method:"post" summary:"修改读者信息"`
	UserIP     string `json:"UserIP"`
	UserName   string `json:"UserName" v:"require"`
	CurrentNum string `json:"CurrentNum"`
	HistoryNum string `json:"HistoryNum"`
}

// 返回修改后的信息
type ReaderModifyRes struct {
	Message      string          `json:"message"`
	UserModified UserInformation `json:"UserModified"`
}

// 读者信息添加接口
type ReaderAddReq struct {
	g.Meta     `path:"/ReaderInformationAdd/Add" tags:"ReaderAddReq" method:"post" summary:"添加读者信息"`
	UserIP     string ` json:"UserIP"`
	UserName   string `json:"UserName"`
	Email      string `json:"Email"`
	CurrentNum int    `json:"CurrentNum"`
	HistoryNum int    `json:"HistoryNum"`
}

type ReaderAddRes struct {
	Message   string          `json:"message"`
	UserAdded UserInformation `json:"UserAdded"`
}
