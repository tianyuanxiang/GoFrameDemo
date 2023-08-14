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
	IUBookQuery interface {
		UBookBorrow(ctx context.Context, req *v1.UBookBorrowReq) (res *v1.UBookBorrowRes, err error)
	}
)

var (
	localUBookQuery IUBookQuery
)

func UBookQuery() IUBookQuery {
	if localUBookQuery == nil {
		panic("implement not found for interface IUBookQuery, forgot register?")
	}
	return localUBookQuery
}

func RegisterUBookQuery(i IUBookQuery) {
	localUBookQuery = i
}
