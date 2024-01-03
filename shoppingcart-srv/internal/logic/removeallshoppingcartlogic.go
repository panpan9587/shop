package logic

import (
	"context"
	"demo/model/mysql"
	"errors"

	"demo/shoppingcart-srv/internal/svc"
	"demo/shoppingcart-srv/shoppingCart"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAllShoppingCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveAllShoppingCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAllShoppingCartLogic {
	return &RemoveAllShoppingCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveAllShoppingCartLogic) RemoveAllShoppingCart(in *shoppingCart.RemoveAllShoppingCartRequest) (*shoppingCart.RemoveAllShoppingCartResponse, error) {
	// todo: add your logic here and delete this line
	// TODO:清空购物车，计算总金额
	//查询列表计算总价格
	var totalPrice float64
	list, err := mysql.GetShoppingCartListByUserId(in.UserId)
	if err != nil {
		return nil, errors.New("购物车数据获取失败")
	}
	for _, val := range list {
		totalPrice += val.GoodsPrice
	}
	err = mysql.RemoveAllShoppingCartByUserId(in.UserId)
	if err != nil {
		return nil, errors.New("购物车清空失败")
	}
	return &shoppingCart.RemoveAllShoppingCartResponse{
		TotalPrice: totalPrice,
		Pong:       true,
	}, nil
}
