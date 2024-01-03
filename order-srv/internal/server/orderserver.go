// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"demo/order-srv/internal/logic"
	"demo/order-srv/internal/svc"
	"demo/order-srv/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

// 创建订单
func (s *OrderServer) CreateOrder(ctx context.Context, in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	l := logic.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

// 修改订单状态
func (s *OrderServer) UpdateOrderStatus(ctx context.Context, in *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	l := logic.NewUpdateOrderStatusLogic(ctx, s.svcCtx)
	return l.UpdateOrderStatus(in)
}

// 支付，返回支付宝付款地址
func (s *OrderServer) AliPay(ctx context.Context, in *order.AliPayRequest) (*order.AliPayResponse, error) {
	l := logic.NewAliPayLogic(ctx, s.svcCtx)
	return l.AliPay(in)
}