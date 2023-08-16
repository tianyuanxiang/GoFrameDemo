package u_borrowhistory

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBorrowHistory struct{}

func New() *sBorrowHistory {
	return &sBorrowHistory{}
}

func init() {
	service.RegisterBorrowHistory(New())
}

func (u *sBorrowHistory) UborrowHistory(ctx context.Context, req *v1.UHistoryBorrowReq) (res *v1.UHistoryBorrowRes, err error) {
	// 直接查询“图书信息借阅表”中用户名和用户ID=当前用户名的所有图书信息。
	object := g.Model("bookborrowinformation").Ctx(ctx).Where("UserIP", req.UserIP)
	if req.UserName != "" {
		object = object.Where("UserName", req.UserName)
	}
	OwnBorrowInformation, err := object.All()
	if err != nil {
		return
	}
	MessageCarrier := make([]v1.UHistoryinformation, 0)
	messages := ""
	for _, record := range OwnBorrowInformation {
		if gconv.Int(record["Flag"]) == 1 {
			messages = "已借阅"
		} else {
			messages = "已归还"
		}
		net := v1.UHistoryinformation{
			ID:             gconv.Int(record["ID"]),
			BookName:       gconv.String(record["BookName"]),
			ISBN:           gconv.String(record["ISBN"]),
			UserIP:         gconv.String(record["UserIP"]),
			UserName:       gconv.String(record["UserName"]),
			ReturnedOrNot:  messages,
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
