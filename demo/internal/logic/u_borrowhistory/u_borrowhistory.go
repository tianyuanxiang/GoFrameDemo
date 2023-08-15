package u_borrowhistory

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type uBorrowHistory struct{}

func New() *uBorrowHistory {
	return &uBorrowHistory{}
}

func init() {
	service.RegisterUBorrowHistory(New())
}

func (u *uBorrowHistory) UborrowHistory(ctx context.Context, req *v1.UHistoryBorrowReq) (res *v1.UHistoryBorrowRes, err error) {
	// 直接查询“图书信息借阅表”中用户名和用户ID=当前用户名的所有图书信息。
	OwnBorrowInformation, err := g.Model("bookborrowinformation").Ctx(ctx).Where("UserIP", req.UserIP).WhereOr("UserName", req.UserName).All()
	if err != nil {
		return
	}
	MessageCarrier := make([]v1.UHistoryinformation, 0)
	for _, record := range OwnBorrowInformation {
		net := v1.UHistoryinformation{
			ID:             gconv.Int(record["ID"]),
			BookName:       gconv.String(record["BookName"]),
			ISBN:           gconv.String(record["ISBN"]),
			UserIP:         gconv.String(record["UserIP"]),
			UserName:       gconv.String(record["UserName"]),
			BorrowingOrder: gconv.Int(record["BorrowingOrder"]),
		}
		MessageCarrier = append(MessageCarrier, net)
	}
	res = &v1.UHistoryBorrowRes{
		Message:       "您的借阅历史如下：",
		HistoryBorrow: &MessageCarrier,
	}
	return
}
