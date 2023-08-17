package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type readermanagementCtl struct{}

var ReaderManagementCtl = new(readermanagementCtl)

// 查询读者信息

func (r *readermanagementCtl) QueryReaderCtl(ctx context.Context, req *v1.ReadershipQueryReq) (res *v1.ReadershipQueryRes, err error) {
	ret, Err := service.ReaderShipManagement().ReaderShipQuery(ctx, req)
	if Err != nil {
		return
	}
	res = &v1.ReadershipQueryRes{
		Message:          ret.Message,
		UserProfileGroup: ret.UserProfileGroup,
	}
	return
}

// 修改读者信息

func (r *readermanagementCtl) ModifyReaderCtl(ctx context.Context, req *v1.ReadershipModifyReq) (res *v1.ReadershipModifyRes, err error) {
	ret, Err := service.ReaderShipManagement().ReaderShipModify(ctx, req)
	if Err != nil {
		return
	}
	res = &v1.ReadershipModifyRes{
		Message:  ret.Message,
		Modified: ret.Modified,
	}
	return
}
