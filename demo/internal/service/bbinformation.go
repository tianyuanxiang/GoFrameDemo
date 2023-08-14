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
	IBBInformation interface {
		BorrowInformationQuery(ctx context.Context, req *v1.BorrowInformationReq) (res *v1.BorrowInformationRes, err error)
		// 还书接口ReturnBooks
		ReturnBooks(ctx context.Context, req *v1.ReturnBookReq) (res *v1.ReturnBookRes, err error)
	}
)

var (
	localBBInformation IBBInformation
)

func BBInformation() IBBInformation {
	if localBBInformation == nil {
		panic("implement not found for interface IBBInformation, forgot register?")
	}
	return localBBInformation
}

func RegisterBBInformation(i IBBInformation) {
	localBBInformation = i
}
