package svc

import (
	"demo/api/internal/config"
	"demo/api/internal/middleware"
	"demo/goods-srv/goodsclient"
	"demo/order-srv/orderclient"
	"demo/shoppingcart-srv/shoppingcartclient"
	"demo/user-srv/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	AuthToken    rest.Middleware
	User         userclient.User
	Goods        goodsclient.Goods
	ShoppingCart shoppingcartclient.ShoppingCart
	Order        orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AuthToken:    middleware.NewAuthTokenMiddleware().Handle,
		User:         userclient.NewUser(zrpc.MustNewClient(c.User)),
		Goods:        goodsclient.NewGoods(zrpc.MustNewClient(c.Goods)),
		ShoppingCart: shoppingcartclient.NewShoppingCart(zrpc.MustNewClient(c.ShoppingCart)),
		Order:        orderclient.NewOrder(zrpc.MustNewClient(c.Order)),
	}
}
