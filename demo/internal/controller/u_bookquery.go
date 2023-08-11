package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type uQuery struct{}

var UQuery = new(uQuery)

func (u *uQuery) UqueryBook(ctx context.Context, req *v1.UBookQueryReq) (res *v1.UBookQueryRes, err error) {
	ret, Err := service.UBookQuery().UBookQuery(ctx, req)
	if Err != nil {
		return
	}
	res = &v1.UBookQueryRes{
		Message:     ret.Message,
		Information: ret.Information,
	}
	return
}
