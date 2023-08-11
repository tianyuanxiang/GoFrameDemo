package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type typeCtl struct{}

var TypeCtl = new(typeCtl)

func (t *typeCtl) QueryBooktype(ctx context.Context, req *v1.BookTypeReq) (res *v1.BookTypeRes, err error) {
	ret, Err := service.BookType().BookTypeQuery(ctx, req)
	if err != nil {
		err = Err
		return
	}
	res = &v1.BookTypeRes{
		Message: ret.Message,
		Type:    ret.Type,
	}
	return
}

func (t *typeCtl) QueryBooktypeNum(ctx context.Context, req *v1.BookTypeNumReq) (res *v1.BookTypeNumRes, err error) {
	// 表的连接查询
	ret2, Err2 := service.BookType().BookTypeNumQuery(ctx, req.BookTypeID)
	if err != nil {
		err = Err2
		return
	}
	res = &v1.BookTypeNumRes{
		Message:     ret2.Message,
		Information: ret2.Information,
	}
	return
}
