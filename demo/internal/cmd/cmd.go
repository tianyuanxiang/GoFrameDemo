package cmd

import (
	"context"
	"demo/internal/service"

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
			s.Group("/api", func(group *ghttp.RouterGroup) {
				// 该中间件应用于全局接口
				group.Middleware(
					service.Middle().MiddlewareCORS,  // CORS权限
					service.Middle().ResponseHandler, // 返回处理
				)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					// 如果有需求，还可以在这里添加中间件
					group.Bind(
						// 该组没有编写特定的中间件，意味着不用判断鉴权
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
				group.Group("/v1/user", func(group *ghttp.RouterGroup) {
					// Auth的鉴权登录
					group.Middleware(service.Middle().Auth)
					group.Bind(
						controller.Auth, //用户登录
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
