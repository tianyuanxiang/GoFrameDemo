package controller

import (
	"context"
	v1 "demo/api/v1"
)

type uHistory struct{}

var UHistory = new(uHistory)

func (u *uHistory) UBorrowHistory(ctx context.Context, in *v1.UHistoryBorrowReq) (out *v1.UHistoryBorrowRes, err error) {
	return
	// num, err := service.H
}
