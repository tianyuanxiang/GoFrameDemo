package auth

import (
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

// GfJWTMiddleware 提供 Json-Web-Token 身份验证实现。如果失败，会返回 401 HTTP 响应。用户可以通过向 LoginHandler 发送 json 请求来获取令牌。
// 用户可以通过向LoginHandler发送json请求来获取令牌，然后需要在身份验证头中传递令牌。例如 Authorization:Bearer XXX_TOKEN_XXX
var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "ID",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

type UserLoginInput struct {
	UserName string
	Password string
}

// 根据登录信息对用户进行身份验证的回调函数
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in UserLoginInput
	)
	//Parse:在处理 HTTP 请求时，根据请求类型解析请求数据到结构体/结构体数组。
	// 其实该函数是针对无输入参数的
	if err := r.Parse(&in); err != nil {
		return "", err
	}
	res, err := g.Model("userinformation").Ctx(ctx).Fields("count(1)").Where("UserName", in.UserName).Where("UserIP").All()
	if err != nil {
		return "", err
	}
	if res != nil {
		return gconv.Int(res[0]["count(1)"]), nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// 处理不进行授权的逻辑
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// 登录期间负责设置私有载荷的函数，默认设置Authenticator函数回调的所有内容
// data里面是什么？
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	// 断言,看看data到底是啥类型
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// 解析并设置用户身份信息，并设置身份信息至每次请求中
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}
