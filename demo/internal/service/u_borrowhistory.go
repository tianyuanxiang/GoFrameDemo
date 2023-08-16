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
	IBorrowHistory interface {
		UborrowHistory(ctx context.Context, req *v1.UHistoryBorrowReq) (res *v1.UHistoryBorrowRes, err error)
	}
)

var (
	localBorrowHistory IBorrowHistory
)

func BorrowHistory() IBorrowHistory {
	if localBorrowHistory == nil {
		panic("implement not found for interface IBorrowHistory, forgot register?")
	}
	return localBorrowHistory
}

func RegisterBorrowHistory(i IBorrowHistory) {
	localBorrowHistory = i
}
