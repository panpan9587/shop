package logic

import (
	"context"
	"demo/model/mysql"
	"demo/order-srv/internal/svc"
	"demo/order-srv/order"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	// todo: add your logic here and delete this line
	// TODO:接收数据创建订单

	//TODO:判断订单是否已经存在
	o, _ := mysql.GetOrderByOrderSn(in.OrderSn)
	if o.ID != 0 {
		return nil, errors.New("订单已被创建，请勿重复创建")
	}

	//TODO:判断商品库存是否足够
	g, err := mysql.GetGoodsById(in.GoodsId)
	if err != nil {
		fmt.Println("库存查询失败")
		return nil, errors.New("商品库存加载失败")
	}
	if in.GoodsNum > g.GoodsStock {
		return nil, errors.New("商品库存不足，无法从创建订单")
	}
	orders := mysql.Order{
		UserId:      in.UserId,
		OrderSn:     in.OrderSn,
		GoodsId:     in.GoodsId,
		GoodsNum:    in.GoodsNum,
		TotalPrice:  in.TotalPrice,
		PayType:     in.PayType,
		OrderStatus: in.OrderStatus,
		RealName:    in.RealName,
		Addr:        in.Addr,
	}
	err = mysql.CreateOrder(&orders)
	if err != nil {
		return nil, errors.New("订单创建失败")
	}
	return &order.CreateOrderResponse{
		OrderId: int64(orders.ID),
	}, nil
}
