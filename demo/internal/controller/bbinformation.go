package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

type bbInformationCtl struct{}

var BBInformation = new(bbInformationCtl)

// 借阅信息查询
func (b *bbInformationCtl) QueryBorrowInformation(ctx context.Context, req *v1.BorrowInformationReq) (res *v1.BorrowInformationRes, err error) {
	QueryOutcome, err := service.BBInformation().BorrowInformationQuery(ctx, req)
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(500, nil)
		return
	}
	res = &v1.BorrowInformationRes{
		Message:     QueryOutcome.Message,
		BBorrowData: QueryOutcome.BBorrowData,
	}
	return
}

// 还书接口
func (b *bbInformationCtl) ReturnBook(ctx context.Context, req *v1.ReturnBookReq) (res *v1.ReturnBookRes, err error) {
	ReturnResult, err := service.BBInformation().ReturnBooks(ctx, req)
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(500, nil)
		return
	}
	res = &v1.ReturnBookRes{
		Message:            ReturnResult.Message,
		ValidateReturnDate: ReturnResult.ValidateReturnDate,
	}
	return
}
