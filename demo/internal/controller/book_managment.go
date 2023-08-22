package controller

import (
	"context"
	"demo/api/v1"
	"demo/internal/service"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

type Contrller struct{}

func New() *Contrller {
	return &Contrller{}
}

// 插入数据
func (c *Contrller) InsertInformation(ctx context.Context, InReq *v1.BookInsertReq) (InRes *v1.BookInsertRes, err error) {
	info, err := service.User().Insert(ctx, *InReq)
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(500, nil)
		return
	}
	// return 为空也没事，因为已经赋过值了
	InRes = &v1.BookInsertRes{
		Message: info.Message,
		Date:    info.Date,
	}
	return
}

// 对请求数据结构进行处理(用户输入参数 : req *v1.Req)
// 得到请求参数，并使用:
func (c *Contrller) Query_Information(ctx context.Context, Q_req *v1.BookQueryReq) (Q_res *v1.BookQueryRes, err error) {
	// 这个输出的应该就是请求的完整url
	fmt.Println(Q_req.Name)
	fmt.Println(Q_req.ISBN)
	// service.User()：这是一个服务层的方法调用，用于获取用户服务的实例
	// Login(ctx, req.Name)：这是调用用户服务中的 Login 方法来进行用户登录操作。其中：ctx 是一个上下文对象，用于传递上下文信息，如请求的上下文环境、超时控制等。
	// req.Name 是从请求中获取的用户名，作为登录的参数。
	info, err := service.User().Query(ctx, Q_req.Name, Q_req.ISBN)
	// 得到最后的输出
	// 有错误
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteStatus(500, nil)
		return
	}
	Q_res = &v1.BookQueryRes{
		Message:     info.Message,
		Information: info.Information,
	}
	return
}

// 更新表中数据
func (c *Contrller) UpdateInformation(ctx context.Context, req *v1.BookUpdateReq) (res *v1.BookUpdateRes, err2 error) {
	res, err := service.User().Update(ctx, *req)
	if err != nil {
		err = err2
		g.RequestFromCtx(ctx).Response.WriteStatus(500, nil)
		return
	}
	res = &v1.BookUpdateRes{
		Message: res.Message,
		Date:    res.Date,
	}
	return
}

// 删除表中数据
func (c *Contrller) DeleteInfromation(ctx context.Context, req *v1.BookDeleteReq) (res *v1.BookDeleteRes, err3 error) {
	cnt, err := service.User().Delete(ctx, *req)
	if err != nil {
		err3 = err
		g.RequestFromCtx(ctx).Response.WriteStatus(500, nil)
		return
	}
	res = &v1.BookDeleteRes{
		Information: cnt.Information,
		Message:     cnt.Message,
	}
	return
}
