// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"demo/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/users/captcha",
				Handler: UserCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/forget/password",
				Handler: ForgetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/login",
				Handler: LoginUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/pay/notify",
				Handler: PayNotifyHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/users/bind/mobile",
					Handler: BindMobileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/identity",
					Handler: IdentityUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/goods/release",
					Handler: ReleaseGoodsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/goods/getinfo",
					Handler: GetGoodsInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shoppingcart/addshop",
					Handler: AddShoppingCartHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shoppingcart/deleteshop",
					Handler: DeleteShoppingCartHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shoppingcart/remove/all",
					Handler: RemoveAllShoppingCartHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/order/create",
					Handler: CreateOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/pay/alipay",
					Handler: AlipayHandler(serverCtx),
				},
			}...,
		),
	)
}
