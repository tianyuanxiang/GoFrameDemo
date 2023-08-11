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
	IBookType interface {
		// 图书类别查询方法
		BookTypeQuery(ctx context.Context, req *v1.BookTypeReq) (res *v1.BookTypeRes, err error)
		BookTypeNumQuery(ctx context.Context, BookTypeID int) (res *v1.BookTypeNumRes, err error)
	}
)

var (
	localBookType IBookType
)

func BookType() IBookType {
	if localBookType == nil {
		panic("implement not found for interface IBookType, forgot register?")
	}
	return localBookType
}

func RegisterBookType(i IBookType) {
	localBookType = i
}
