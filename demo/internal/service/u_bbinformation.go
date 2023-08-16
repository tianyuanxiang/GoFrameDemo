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
	IUserBBinformation interface {
		UserQueryBorrowBook(ctx context.Context, req *v1.UserBookBorrowReq) (res *v1.UserBookBorrowRes, err error)
	}
)

var (
	localUserBBinformation IUserBBinformation
)

func UserBBinformation() IUserBBinformation {
	if localUserBBinformation == nil {
		panic("implement not found for interface IUserBBinformation, forgot register?")
	}
	return localUserBBinformation
}

func RegisterUserBBinformation(i IUserBBinformation) {
	localUserBBinformation = i
}
