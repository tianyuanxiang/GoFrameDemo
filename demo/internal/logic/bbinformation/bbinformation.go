package bbinformation

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sBBInformation struct{}

func New() *sBBInformation {
	return &sBBInformation{}
}

func init() {
	service.RegisterBBInformation(New())
}

// 查询某本图书的借阅信息
// 问题：
// "ReturnDate": "", 归还日期为空

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
			BookName:       gconv.String(element["BookName"]),
			ISBN:           gconv.String(element["ISBN"]),
			UserIP:         gconv.String(element["UserIP"]),
			UserName:       gconv.String(element["UserName"]),
			BorrowDate:     gconv.String(element["created_at"]),
			ReturnDate:     gconv.String(element["ReturnDate"]),
			BorrowingOrder: gconv.Int(element["BorrowingOrder"]),
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
// 问题：
// 1. 点击还书功能，会归还所有借阅的同名图书，改为归还指定图书
// 2. 请求返回时BorrowDate和ReturnDate不显示
// 3. BorrowDate未更新

func (s *sBBInformation) ReturnBooks(ctx context.Context, req *v1.ReturnBookReq) (res *v1.ReturnBookRes, err error) {
	// 事务开启
	db := g.DB()
	errTransaction := db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 图书借阅信息表中该记录的flag = 0、图书归还日期更新至当日
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 传入图书ISBN和BorrowingOrder；
			_, err = g.Model("bookborrowinformation").Ctx(ctx).Data(g.Map{"Flag": 0, "ReturnDate": gtime.New(time.Now())}).Where("BookName", req.BookName).Where("BorrowingOrder", req.BorrowingOrder).Update()
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
	ModifyTime, err := g.Model("bookborrowinformation").Fields("created_at", "ReturnDate").Where("BorrowingOrder", req.BorrowingOrder).One()
	if err != nil {
		return
	}
	res = &v1.ReturnBookRes{
		Message: "还书信息如下：",
		ValidateReturnDate: v1.BBinformation{
			BookName:       req.BookName,
			ISBN:           req.ISBN,
			UserIP:         req.UserIP,
			UserName:       req.UserName,
			BorrowDate:     gconv.String(ModifyTime["created_at"]),
			ReturnDate:     gconv.String(ModifyTime["ReturnDate"]),
			BorrowingOrder: req.BorrowingOrder,
		},
	}
	return
}
