package logic

import (
	"context"
	"demo/common"
	"demo/order-srv/order"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// TODO: 生成数据创建订单
	res, err := l.svcCtx.Order.CreateOrder(l.ctx, &order.CreateOrderRequest{
		UserId:      req.UserId,
		OrderSn:     common.GetOrderSn(),
		GoodsId:     req.GoodsId,
		GoodsNum:    req.GoodsNum,
		TotalPrice:  req.TotalPrice,
		PayType:     req.PayType,
		OrderStatus: req.OrderStatus,
		RealName:    req.RealName,
		Addr:        req.Addr,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "订单创建失败",
		}, err
	}
	if res.OrderId == 0 {
		return &types.Response{
			Code:    500,
			Message: "订单创建失败",
		}, nil

	}
	return &types.Response{
		Code:    200,
		Message: "订单创建成功",
	}, nil
}
