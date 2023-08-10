package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

var TypeCtl = new(typeCtl)

type typeCtl struct{}

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
