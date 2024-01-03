package logic

import (
	"context"
	"demo/common"
	"demo/model/mysql"
	"errors"

	"demo/order-srv/internal/svc"
	"demo/order-srv/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type AliPayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAliPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AliPayLogic {
	return &AliPayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 支付，返回支付宝付款地址
func (l *AliPayLogic) AliPay(in *order.AliPayRequest) (*order.AliPayResponse, error) {
	// todo: add your logic here and delete this line
	// TODO: 根据订单Id获取订单
	o, err := mysql.GetOrderById(in.OrderId)
	if err != nil {
		return nil, errors.New("订单获取失败")
	}
	payUrl := common.AliPay(o)
	return &order.AliPayResponse{
		PayURL: payUrl,
	}, nil
}
