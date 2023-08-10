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
	IUser interface {
		// 新增
		Insert(ctx context.Context, in v1.BookInsertReq) (out *v1.BookInsertRes, err error)
		// 得到书名和出版日期，判断有没有
		// 都隶属于sUser这个结构体
		Query(ctx context.Context, name string, ISBN string) (out *v1.BookQueryRes, err error)
		// 修改
		Update(ctx context.Context, in v1.BookUpdateReq) (outUpdated *v1.BookUpdateRes, err error)
		// 删除
		Delete(ctx context.Context, in v1.BookDeleteReq) (OutDeleted *v1.BookDeleteRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
