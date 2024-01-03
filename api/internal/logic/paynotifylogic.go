package logic

import (
	"context"
	"demo/common"
	"demo/order-srv/order"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"net/http"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayNotifyLogic {
	return &PayNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayNotifyLogic) PayNotify(req *types.PayNotifyReq, w http.ResponseWriter, r *http.Request) (resp *types.PayNotifyResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("1111111111")
	r.ParseForm()
	noti, err := common.AliPayClient.DecodeNotification(r.Form)
	if err != nil {
		fmt.Println(err)
		return
	}
	orderStatus := 1
	switch noti.TradeStatus {
	case alipay.TradeStatusSuccess: //交易支付成功
		fmt.Println("success")
		orderStatus = 2
	case alipay.TradeStatusClosed:
		fmt.Println("未付款交易超市关闭，或支付完成后全额退款")
		orderStatus = 3
	case alipay.TradeStatusFinished:
		fmt.Println("交易结束，不可退款")
		orderStatus = 4
	case alipay.TradeStatusWaitBuyerPay:
		fmt.Println("交易创建，等待买家付款")
	}
	//调用订单修改状态的rpc
	res, err := l.svcCtx.Order.UpdateOrderStatus(l.ctx, &order.UpdateOrderStatusRequest{
		OrderSn: noti.OutTradeNo,
		Status:  int64(orderStatus),
	})
	fmt.Println("2222222222222222")
	if err != nil {
		fmt.Println("更新订单状态失败", "11111111111111")
		fmt.Println(err)
		return
	}
	if !res.Pong {
		fmt.Println("更新订单状态失败", "111111111111111111")
	}
	return
}
