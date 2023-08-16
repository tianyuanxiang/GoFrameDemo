package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type uBooksBorrowed struct {
}

var UBooksBorrowed = new(uBooksBorrowed)

func (u *uBooksBorrowed) UBooksBorrowed(ctx context.Context, in *v1.UserBookBorrowReq) (out *v1.UserBookBorrowRes, err error) {
	ret, Err := service.UserBBinformation().UserQueryBorrowBook(ctx, in)
	if Err != nil {
		return
	}
	out = &v1.UserBookBorrowRes{
		Message:          ret.Message,
		UBookBorrowGroup: ret.UBookBorrowGroup,
	}
	return
}
