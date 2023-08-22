package middle

import (
	"demo/internal/logic/auth"
	"demo/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
)

type sMiddle struct{}

func New() *sMiddle {
	return &sMiddle{}
}
func init() {
	service.RegisterMiddle(New())
}

// 中间件,用于作返回处理
func (s *sMiddle) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	var (
		err = r.GetError()
		// 获取请求的响应对象
		res = r.GetHandlerResponse()
	)
	if err != nil {
		// 系统返回的错误码
		code := gerror.Code(err).Code()
		switch code {
		case 52:
			r.Response.WriteStatus(400, nil)
			r.Response.WriteJsonExit(g.Map{
				"message": "数据库校验错误",
			})
		case 51:
			r.Response.WriteStatus(400, nil)
			r.Response.WriteJsonExit(g.Map{
				"message": "参数校验错误" + gconv.String(err),
				"data":    make([]string, 0),
			})
		default:
			r.Response.WriteStatus(http.StatusOK)
			r.Response.WriteJsonExit(res)
		}
	} else {
		r.Response.WriteStatus(http.StatusOK)
		r.Response.WriteJsonExit(res)
	}
}

// 开启CORS权限
func (s *sMiddle) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 登录中间件
func (s *sMiddle) Auth(r *ghttp.Request) {
	auth.Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}
