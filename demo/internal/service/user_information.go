// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/v1"
)

type (
	IUserInformation interface {
		// 查询读者信息方法
		UserInformationQuery(ctx context.Context, req *v1.ReaderInformationReq) (res *v1.ReaderInformationRes, err error)
		// 修改读者信息方法
		UserInformationModify(ctx context.Context, req *v1.ReaderModifyReq) (res *v1.ReaderModifyRes, err error)
		// 添加读者信息方法
		UserInformationAdd(ctx context.Context, req *v1.ReaderAddReq) (res *v1.ReaderAddRes, err error)
	}
)

var (
	localUserInformation IUserInformation
)

func UserInformation() IUserInformation {
	if localUserInformation == nil {
		panic("implement not found for interface IUserInformation, forgot register?")
	}
	return localUserInformation
}

func RegisterUserInformation(i IUserInformation) {
	localUserInformation = i
}
