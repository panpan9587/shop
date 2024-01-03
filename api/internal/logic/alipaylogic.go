package logic

import (
	"context"
	"demo/order-srv/order"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlipayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlipayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlipayLogic {
	return &AlipayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlipayLogic) Alipay(req *types.AlipayReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// TODO: 调用支付路由，获取订单支付地址
	res, err := l.svcCtx.Order.AliPay(l.ctx, &order.AliPayRequest{OrderId: req.OrderId})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "订单获取失败",
		}, err
	}
	if res.PayURL == "" {
		return &types.Response{
			Code:    500,
			Message: "支付获取失败，请检查您的网络",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "支付宝支付",
		Data: map[string]interface{}{
			"pay_url": res.PayURL,
		},
	}, nil
}
