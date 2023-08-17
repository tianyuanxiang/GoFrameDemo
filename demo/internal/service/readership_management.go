// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/v1"
)

type (
	IReaderShipManagement interface {
		ReaderShipQuery(ctx context.Context, req *v1.ReadershipQueryReq) (res *v1.ReadershipQueryRes, err error)
		ReaderShipModify(ctx context.Context, req *v1.ReadershipModifyReq) (res *v1.ReadershipModifyRes, err error)
	}
)

var (
	localReaderShipManagement IReaderShipManagement
)

func ReaderShipManagement() IReaderShipManagement {
	if localReaderShipManagement == nil {
		panic("implement not found for interface IReaderShipManagement, forgot register?")
	}
	return localReaderShipManagement
}

func RegisterReaderShipManagement(i IReaderShipManagement) {
	localReaderShipManagement = i
}
