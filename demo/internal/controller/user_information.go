package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type userCtl struct{}

var UserCtl = new(userCtl)

func (u *userCtl) QueryUserInformation(ctx context.Context, req *v1.ReaderInformationReq) (res *v1.ReaderInformationRes, err error) {
	ret, err := service.UserInformation().UserInformationQuery(ctx, req)
	if err != nil {
		return
	}
	res = &v1.ReaderInformationRes{
		Message: ret.Message,
		User:    ret.User,
	}
	return
}

func (u *userCtl) AddUserInformation(ctx context.Context, req *v1.ReaderAddReq) (res *v1.ReaderAddRes, err error) {
	ret, err := service.UserInformation().UserInformationAdd(ctx, req)
	if err != nil {
		return
	}
	res = &v1.ReaderAddRes{
		Message:   ret.Message,
		UserAdded: ret.UserAdded,
	}
	return
}
