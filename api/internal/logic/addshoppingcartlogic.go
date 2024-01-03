package logic

import (
	"context"
	"demo/shoppingcart-srv/shoppingCart"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddShoppingCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddShoppingCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddShoppingCartLogic {
	return &AddShoppingCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddShoppingCartLogic) AddShoppingCart(req *types.AddShoppingCartReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	//TODO: 添加购物车
	res, err := l.svcCtx.ShoppingCart.AddShoppingCart(l.ctx, &shoppingCart.AddShoppingCartRequest{
		UserId:     req.UserId,
		GoodsId:    req.GoodsId,
		GoodsPrice: req.GoodsPrice,
		GoodsNum:   req.GoodsNum,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "购物商品添加失败",
		}, err
	}
	if res.ShopId == 0 {
		return &types.Response{
			Code:    500,
			Message: "购物商品添加失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "购物车商品添加成功",
		Data: map[string]interface{}{
			"shopId": res.ShopId,
		},
	}, nil
}
