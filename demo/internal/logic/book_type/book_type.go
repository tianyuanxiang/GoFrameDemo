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
	ArrType := make([]v1.BookType, 0)
	for _, j := range pub {
		net := v1.BookType{
			ID:         gconv.Int(j["ID"]),
			BookTypeID: gconv.Int(j["BookTypeID"]),
			TypeName:   gconv.String(j["TypeName"]),
		}
		ArrType = append(ArrType, net)
	}
	res = &v1.BookTypeRes{
		Message: "图书类别信息如下:",
		Type:    &ArrType,
	}
	return
}
