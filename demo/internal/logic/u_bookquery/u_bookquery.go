package u_bookquery

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/logic/book_managment"
	"demo/internal/service"
	"fmt"
)

type sUBookQuery struct{}

func New() *sUBookQuery {
	return &sUBookQuery{}
}

func init() {
	service.RegisterUBookQuery(New())
}
func (s *sUBookQuery) UBookQuery(ctx context.Context, in *v1.UBookQueryReq) (out *v1.UBookQueryRes, err error) {
	result, err := book_managment.New().Query(ctx, in.BookName, in.ISBN)
	if err != nil {
		return
	}
	fmt.Println(result)
	// bookArr := make([]v1.UBookInformation, 0)

	//out = &v1.UBookQueryRes{
	//	Message: "图书信息如下：",
	//	Information: result.Information,
	//
	//	}
	//}
	return
}
