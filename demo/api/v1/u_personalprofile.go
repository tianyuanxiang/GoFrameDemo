package v1

import "github.com/gogf/gf/v2/frame/g"

// 查询个人资料

type PersonalProfileQueryReq struct {
	g.Meta   `path:"/PersonalProfileQuery/Query" tags:"UserPersonalProfile"  method:"get" summary:"用户查询个人资料"`
	UserIP   string `json:"UserIP" v:"required"`
	UserName string `json:"UserName" v:"required"`
}

type PersonalProfileQueryRes struct {
	Message                    string              `json:"message"`
	PersonalInformationDisplay PersonalInformation `json:"PersonalInformationDisplay" `
}

type PersonalInformation struct {
	UserIP     string `json:"userIP"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	CurrentNum int    `json:"currentNum"`
	// 图书名:该书数量
	Data       interface{} `json:"data"`
	HistoryNum int         `json:"historyNum"`
}

// 修改个人资料

type PersonalProfileModifyReq struct {
	g.Meta   `path:"/PersonalProfileModify/Modify" tags:"UserPersonalProfileModify"  method:"post" summary:"用户修改个人资料"`
	UserIP   string `json:"UserIP"`
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
}

type PersonalProfileModifyRes struct {
	Message          string `json:"message"`
	UserIP           string `json:"userIP"`
	ModifiedUserName string `json:"userName"`
	ModifiedEmail    string `json:"email"`
}
