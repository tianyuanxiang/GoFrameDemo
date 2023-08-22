// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddle interface {
		// 中间件,用于作返回处理
		ResponseHandler(r *ghttp.Request)
		// 开启CORS权限
		MiddlewareCORS(r *ghttp.Request)
		// 登录中间件
		Auth(r *ghttp.Request)
	}
)

var (
	localMiddle IMiddle
)

func Middle() IMiddle {
	if localMiddle == nil {
		panic("implement not found for interface IMiddle, forgot register?")
	}
	return localMiddle
}

func RegisterMiddle(i IMiddle) {
	localMiddle = i
}
