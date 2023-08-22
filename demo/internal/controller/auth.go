package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/logic/auth"
)

type authController struct{}

var Auth = authController{}

func (c *authController) Login(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	res = &v1.AuthLoginRes{}
	res.Message = "登录信息如下:"
	res.Token, res.Expire = auth.Auth().LoginHandler(ctx)
	return
}

func (c *authController) RefreshToken(ctx context.Context, req *v1.AuthRefreshTokenReq) (res *v1.AuthRefreshTokenRes, err error) {
	res = &v1.AuthRefreshTokenRes{}
	res.Token, res.Expire = auth.Auth().RefreshHandler(ctx)
	return
}

func (c *authController) Logout(ctx context.Context, req *v1.AuthRefreshTokenReq) (res *v1.AuthRefreshTokenRes, err error) {
	auth.Auth().LogoutHandler(ctx)
	return
}
