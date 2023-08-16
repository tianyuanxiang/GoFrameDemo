package u_bbinformation

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUserBBinformation struct{}

func New() *sUserBBinformation {
	return &sUserBBinformation{}
}

func init() {
	service.RegisterUserBBinformation(New())
}

// 用户查询借阅信息

func (s *sUserBBinformation) UserQueryBorrowBook(ctx context.Context, req *v1.UserBookBorrowReq) (res *v1.UserBookBorrowRes, err error) {
	arr, err := g.Model("bookborrowinformation").Ctx(ctx).Where("UserIP", req.UserIP).Where("Flag", 1).WhereOr("UserName", req.UserName).All()
	if err != nil {
		return
	}
	ResultsArray := make([]v1.UBookBorrowGroup, 0)
	for _, ele := range arr {
		net := v1.UBookBorrowGroup{
			ID:             gconv.Int(ele["ID"]),
			BookName:       gconv.String(ele["BookName"]),
			ISBN:           gconv.String(ele["ISBN"]),
			BorrowTime:     gconv.String(ele["created_at"]),
			ReturnTime:     gconv.String(ele["ReturnDate"]),
			BorrowingOrder: gconv.Int(ele["BorrowingOrder"]),
		}
		ResultsArray = append(ResultsArray, net)
	}
	res = &v1.UserBookBorrowRes{
		Message:          "您的借阅信息如下：",
		UBookBorrowGroup: &ResultsArray,
	}
	return
}
