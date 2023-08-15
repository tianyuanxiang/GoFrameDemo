package book_type

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBookType struct{}

func New() *sBookType {
	return &sBookType{}
}

func init() {
	service.RegisterBookType(New())
}

// 图书类别查询方法

func (s *sBookType) BookTypeQuery(ctx context.Context, req *v1.BookTypeReq) (res *v1.BookTypeRes, err error) {
	pub, err := g.Model("booktype").Ctx(ctx).All()
	if err != nil {
		return
	}
	ArrType1 := make([]v1.BookType, 0)
	for _, j := range pub {
		net := v1.BookType{
			ID:         gconv.Int(j["ID"]),
			BookTypeID: gconv.Int(j["BookTypeID"]),
			TypeName:   gconv.String(j["TypeName"]),
		}
		ArrType1 = append(ArrType1, net)
	}
	res = &v1.BookTypeRes{
		Message: "图书类别信息如下:",
		Type:    &ArrType1,
	}
	return
}
func (s *sBookType) BookTypeNumQuery(ctx context.Context, BookTypeID int) (res *v1.BookTypeNumRes, err error) {
	val, err := g.Model("booktype t").Ctx(ctx).RightJoin("bookinformation b", "t.BookTypeID = b.BookTypeID").
		Fields("b.*").Where("b.BookTypeID", BookTypeID).All()
	if err != nil {
		return
	}
	ArrType2 := make([]v1.BookTypeInformation, 0)
	for _, flag := range val {
		SingleRecord := v1.BookTypeInformation{
			Id:         gconv.Int(flag["ID"]),
			Name:       gconv.String(flag["BookName"]),
			ISBN:       gconv.String(flag["ISBN"]),
			Author:     gconv.String(flag["Author"]),
			Publishers: gconv.String(flag["Publishers"]),
			BookTypeID: gconv.Int(flag["BookTypeID"]),
			Amount:     gconv.Int(flag["Amount"]),
		}
		ArrType2 = append(ArrType2, SingleRecord)
	}
	res = &v1.BookTypeNumRes{
		Message:     "该类型的图书如下：",
		Information: &ArrType2,
	}
	return
}
