package main

import (
	"demo/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入MySQL驱动

	_ "demo/internal/logic"
	//"net/http"
	//
	//"github.com/gogf/gf/v2/frame/g"
	//"github.com/gogf/gf/v2/net/ghttp"
)

//func MiddlewareAuth(r *ghttp.Request) {
//	token := r.Get("token")
//	if token.String() == "123456" {
//		r.Response.Writeln("auth")
//		r.Middleware.Next()
//	} else {
//		r.Response.WriteStatus(http.StatusForbidden)
//	}
//}
//
//func MiddlewareCORS(r *ghttp.Request) {
//	r.Response.Writeln("cors")
//	r.Response.CORSDefault()
//	r.Middleware.Next()
// }

func main() {
	//s := g.Server()
	//s.Group("/api.v2", func(group *ghttp.RouterGroup) {
	//	// CORS:跨域中间件 //MiddlewareAuth: 鉴权中间件
	//	// 请求时将会按照中间件注册的先后顺序，先执行MiddlewareCORS全局中间件，再执行MiddlewareAuth分组中间件。
	//	group.Middleware(MiddlewareCORS, MiddlewareAuth)
	//	group.ALL("/user/list", func(r *ghttp.Request) {
	//		r.Response.Writeln("list")
	//	})
	//})
	//s.SetPort(8199)
	//s.Run()
	cmd.Main.Run(gctx.New())
}
