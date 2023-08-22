package v1

import "github.com/gogf/gf/v2/frame/g"

// 读者信息查询接口

type ReadershipQueryReq struct {
	g.Meta `path:"/ReadershipInformation/Query" tags:"ReadershipInfromation"  method:"get" summary:"查询用户信息"`
}

type ReadershipQueryRes struct {
	Message          string         `json:"message"`
	UserProfileGroup *[]UserProfile `json:"userProfileGroup"`
}

type UserProfile struct {
	UserIP     string `json:"userIP"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	CurrentNum int    `json:"currentNum"`
	HistoryNum int    `json:"historyNum"`
}

// 读者信息修改接口

type ReadershipModifyReq struct {
	g.Meta   `path:"/ReadershipInformation/Modify" tags:"ReadershipInfromation"  method:"post" summary:"修改用户信息"`
	UserIP   string `json:"userIP" `
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type ReadershipModifyRes struct {
	Message  string      `json:"message"`
	Modified UserProfile `json:"modified"`
}

// 读者信息删除接口
