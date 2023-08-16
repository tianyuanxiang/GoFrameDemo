package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type uHistory struct{}

var UHistory = new(uHistory)

func (u *uHistory) UBorrowHistory(ctx context.Context, in *v1.UHistoryBorrowReq) (out *v1.UHistoryBorrowRes, err error) {
	num, err := service.BorrowHistory().UborrowHistory(ctx, in)
	if err != nil {
		return
	}
	out = &v1.UHistoryBorrowRes{
		Message:       num.Message,
		HistoryBorrow: num.HistoryBorrow,
	}
	return
}
