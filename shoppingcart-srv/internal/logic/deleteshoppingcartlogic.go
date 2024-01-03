package logic

import (
	"context"
	"demo/model/mysql"
	"errors"
	"fmt"

	"demo/shoppingcart-srv/internal/svc"
	"demo/shoppingcart-srv/shoppingCart"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShoppingCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteShoppingCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShoppingCartLogic {
	return &DeleteShoppingCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 移除购物车
func (l *DeleteShoppingCartLogic) DeleteShoppingCart(in *shoppingCart.DeleteShoppingCartRequest) (*shoppingCart.DeleteShoppingCartResponse, error) {
	// todo: add your logic here and delete this line
	//根据id删除购物车商品
	err := mysql.RemoveShoppingCart(in.ShopId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("删除数据失败")
	}
	return &shoppingCart.DeleteShoppingCartResponse{
		Pong: true,
	}, nil
}
