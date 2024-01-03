package logic

import (
	"context"
	"demo/shoppingcart-srv/shoppingCart"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveAllShoppingCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveAllShoppingCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveAllShoppingCartLogic {
	return &RemoveAllShoppingCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveAllShoppingCartLogic) RemoveAllShoppingCart(req *types.RemoveAllShoppingCartReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ShoppingCart.RemoveAllShoppingCart(l.ctx, &shoppingCart.RemoveAllShoppingCartRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "清空失败",
		}, err
	}
	if !res.Pong {
		return &types.Response{
			Code:    500,
			Message: "清空失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "清空成功",
		Data: map[string]interface{}{
			"totalPrice": res.TotalPrice,
		},
	}, nil
}
