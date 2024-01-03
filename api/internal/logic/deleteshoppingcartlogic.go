package logic

import (
	"context"
	"demo/shoppingcart-srv/shoppingCart"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShoppingCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShoppingCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShoppingCartLogic {
	return &DeleteShoppingCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 删除购物车商品
func (l *DeleteShoppingCartLogic) DeleteShoppingCart(req *types.DeleteShoppingCartReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ShoppingCart.DeleteShoppingCart(l.ctx, &shoppingCart.DeleteShoppingCartRequest{
		ShopId: req.ShopId,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "购物车物品删除失败",
		}, err
	}
	if !res.Pong {
		return &types.Response{
			Code:    500,
			Message: "购物车物品删除失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "购物车物品删除成功",
	}, nil
}
