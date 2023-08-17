package cmd

import (
	"context"

	"demo/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					// New 结构体绑定了很多方法，可以统一注册
					controller.New(),               // 图书管理
					controller.TypeCtl,             // 图书类型管理
					controller.UserCtl,             // 用户管理
					controller.UBorrow,             // 用户查询
					controller.BBInformation,       // 用户借阅信息查询
					controller.UHistory,            // 用户历史借阅信息查询
					controller.UBooksBorrowed,      // 用户当前借阅信息查询
					controller.UPersonalProfile,    // 用户个人信息
					controller.ReaderManagementCtl, //读者信息管理
				)
			})
			s.Run()
			return nil
		},
	}
)
