package logic

import (
	"context"
	"demo/model/mysql"
	"errors"
	"fmt"

	"demo/order-srv/internal/svc"
	"demo/order-srv/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改订单状态
func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *order.UpdateOrderStatusRequest) (*order.UpdateOrderStatusResponse, error) {
	// todo: add your logic here and delete this line
	//根据订单编号查询
	o, _ := mysql.GetOrderByOrderSn(in.OrderSn)
	err := mysql.UpdateOrderStatus(int64(o.ID), in.Status)
	if err != nil {
		fmt.Println(err)
		return &order.UpdateOrderStatusResponse{
			Pong: false,
		}, errors.New("订单状态修改失败")
	}
	return &order.UpdateOrderStatusResponse{
		Pong: true,
	}, nil
}
