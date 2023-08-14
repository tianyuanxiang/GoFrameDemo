package bbinformation

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBBInformation struct{}

func New() *sBBInformation {
	return &sBBInformation{}
}

func init() {
	service.RegisterBBInformation(New())
}

func (s *sBBInformation) BorrowInformationQuery(ctx context.Context, req *v1.BorrowInformationReq) (res *v1.BorrowInformationRes, err error) {

	object := g.Model("bookborrowinformation").Ctx(ctx).Where("Flag", 1).WhereLike("BookName", "%"+req.BookName+"%")
	if req.UserIP != "" {
		object.WhereLike("UserIP", req.UserIP+"%")
	}
	if req.UserName != "" {
		object.WhereLike("UserName", req.UserName+"%")
	}
	all, err := object.All()
	if err != nil {
		return
	}
	CarrierArray := make([]v1.BBinformation, 0)
	for _, element := range all {
		Carrier := v1.BBinformation{
			BookName:   gconv.String(element["BookName"]),
			ISBN:       gconv.String(element["ISBN"]),
			UserIP:     gconv.String(element["UserIP"]),
			UserName:   gconv.String(element["UserName"]),
			BorrowDate: gconv.String(element["BorrowDate"]),
			ReturnDate: gconv.String(element["ReturnDate"]),
		}
		CarrierArray = append(CarrierArray, Carrier)
	}
	res = &v1.BorrowInformationRes{
		Message:     "图书借阅信息如下：",
		BBorrowData: &CarrierArray,
	}
	return
}

// 还书接口ReturnBooks
func (s *sBBInformation) ReturnBooks(ctx context.Context, req *v1.ReturnBookReq) (res *v1.ReturnBookRes, err error) {
	// 事务开启
	db := g.DB()
	errTransaction := db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 图书借阅信息表中该记录的flag = 0、图书归还日期更新至当日
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			//_, err = db.Ctx(ctx).Update(ctx, "bookborrowinformation",updateData, "ISBN", req.ISBN)
			// updated_at字段自动更新
			_, err = g.Model("bookborrowinformation").Ctx(ctx).Data("Flag", 0).Where("ISBN", req.ISBN).Update()
			if err != nil {
				return err
			}
			return err
		})
		// 图书信息表的该书数量 +1
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			updateData := g.Map{
				"Amount": &gdb.Counter{
					Field: "Amount",
					Value: 1,
				},
			}
			_, err = db.Ctx(ctx).Update(ctx, "bookinformation", updateData, "ISBN", req.ISBN)
			if err != nil {
				return err
			}
			return err
		})
		//读者信息表的该读者借阅数量 -1
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			num, err := g.Model("userinformation").Ctx(ctx).Fields("CurrentNum").Where("UserIP", req.UserIP).Value()
			if gconv.Int(num) <= 0 {
				//res = &v1.UBookBorrowRes{
				//	Message:     "图书库存不足！",
				//	Information: v1.UBBInformation{},
				//}
				panic("error")
				return err
			}
			_, err = g.Model("userinformation").Ctx(ctx).Where("UserIP", req.UserIP).Decrement("CurrentNum", 1)
			return err
		})
		return nil
	})
	if errTransaction != nil {
		return
	}
	res = &v1.ReturnBookRes{
		Message: "还书信息如下：",
		ReturnDate: v1.BBinformation{
			BookName: req.BookName,
			ISBN:     req.ISBN,
			UserIP:   req.UserIP,
			UserName: req.UserName,
		},
	}
	return
}
