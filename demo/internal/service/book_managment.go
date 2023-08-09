// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"demo/api/v1"
)

type (
	IUser interface {
		// 新增
		Insert(ctx context.Context, in v1.InsertReq) (out *v1.InsertRes, err error)
		// 得到书名和出版日期，判断有没有
		// 都隶属于sUser这个结构体
		Query(ctx context.Context, name string, ISBN string, PublisherId int) (out *v1.Q_Res, err error)
		// 修改
		Update(ctx context.Context, in v1.UpdateReq) (outUpdated *v1.UpdateRes, err error)
		// 删除
		Delete(ctx context.Context, in v1.DeleteReq) (OutDeleted *v1.DeleteRes, err error)
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
