package v1

import "github.com/gogf/gf/v2/frame/g"

// 查询个人资料

type PersonalProfileQueryReq struct {
	g.Meta   `path:"/PersonalProfileQuery/Query" tags:"UserPersonalProfile"  method:"post" summary:"用户修改个人信息"`
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
