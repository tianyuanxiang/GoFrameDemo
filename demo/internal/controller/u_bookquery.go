package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type uBorrow struct{}

var UBorrow = new(uBorrow)

func (u *uBorrow) UBorrowBook(ctx context.Context, req *v1.UBookBorrowReq) (res *v1.UBookBorrowRes, err error) {
	ret, Err := service.UBookQuery().UBookBorrow(ctx, req)
	if Err != nil {
		return
	}
	res = &v1.UBookBorrowRes{
		Message:     ret.Message,
		Information: ret.Information,
	}
	return
}
