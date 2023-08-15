package u_bookquery

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUBookQuery struct{}

func New() *sUBookQuery {
	return &sUBookQuery{}
}

func init() {
	service.RegisterUBookQuery(New())
}

// 借书接口
func (s *sUBookQuery) UBookBorrow(ctx context.Context, req *v1.UBookBorrowReq) (res *v1.UBookBorrowRes, err error) {
	mid := req.BorrowInformation
	// transaction 1.
	db := g.DB()
	err1 := db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 图书借阅信息表的记录 +1
			// 借阅日期和归还日期自动生成，其中归还日期为归还日期+70天
			_, err := g.Model("bookborrowinformation").Ctx(ctx).Data(g.Map{"BookName": mid.BookName, "ISBN": mid.ISBN, "UserIP": mid.UserIP,
				"UserName": mid.UserName, "Flag": 1}).Insert()
			Result, err := g.Model("bookborrowinformation").Ctx(ctx).Fields("MAX(ID)", "created_at").Where("ISBN", mid.ISBN).Group("created_at").One()
			var (
				year  = 0
				mouth = 0
				day   = 70
			)
			// updated_at更新时间被他们写死了，只要设置了updated_at字段，不管插入什么数据，都会被updated_at(当前修改时间所覆盖)
			date := gtime.New(gconv.String(Result["created_at"]))
			date2 := date.AddDate(year, mouth, day)
			_, err = g.Model("bookborrowinformation").Ctx(ctx).Data(g.Map{"ReturnDate": gconv.String(date2), "BorrowingOrder": &gdb.Counter{
				Field: "ID",
				Value: 10000}}).Where("ID", gconv.Int(Result["MAX(ID)"])).Update()
			return err
		})
		if err != nil {
			return err
		}
		// transaction 2.
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 判断图书数量是否>0
			num, err := g.Model("bookinformation").Ctx(ctx).Fields("Amount").Where("ISBN", mid.ISBN).Value()
			if gconv.Int(num) <= 0 {
				res = &v1.UBookBorrowRes{
					Message:     "图书库存不足！",
					Information: v1.UBBInformation{},
				}
				panic("error")
				return err
			}
			// 图书信息表的该书数量 -1
			// 数量自减
			_, err = g.Model("bookinformation").Ctx(ctx).Where("ISBN", mid.ISBN).Decrement("Amount", 1)
			return err
			// 如果后续发现库存数量可以为-1，那么加个panic("error")，然后用if语句来判断
		})
		if err != nil {
			return err
		}
		// transaction 3.
		err = db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 读者信息表的该读者借阅数量和历史数量同时 +1
			UpdateData := g.Map{
				"CurrentNum": &gdb.Counter{
					Field: "CurrentNum",
					Value: 1,
				},
				"HistoryNum": &gdb.Counter{
					Field: "HistoryNum",
					Value: 1,
				},
			}
			_, err := db.Ctx(ctx).Update(ctx, "userinformation", UpdateData, "UserIP", mid.UserIP)
			return err
		})
		if err != nil {
			return err
		}
		// 没毛病
		return nil
	})
	// 有毛病
	if err1 != nil {
		res = &v1.UBookBorrowRes{
			Message:     "操作错误",
			Information: v1.UBBInformation{},
		}
		return
	}
	// 没毛病
	res = &v1.UBookBorrowRes{
		Message: "借阅成功，图书信息为：",
		Information: v1.UBBInformation{
			BookName: mid.BookName,
			ISBN:     mid.ISBN,
			UserIP:   mid.UserIP,
			UserName: mid.UserName,
		},
	}
	return
}

// 测试：
// 图书信息表库存为0时，是否引发panic
// 借书日期为当日，还书日期自动计算

// 已写好，未调试
