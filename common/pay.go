package common

import (
	"demo/config"
	"demo/model/mysql"
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

var AliPayClient *alipay.Client

func AliPay(order *mysql.Order) string {
	// 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(config.AliPayConfig.AppId, config.AliPayConfig.PrivateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	AliPayClient = client
	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://4658e545.r5.cpolar.top/pay/notify"
	p.ReturnURL = "http://xxx"
	p.Subject = order.OrderSn
	p.OutTradeNo = order.OrderSn
	p.TotalAmount = fmt.Sprintf("%.2f", order.TotalPrice)
	p.ProductCode = "QUICK_WAP_WAY"

	var url, errs = client.TradeWapPay(p)
	if errs != nil {
		fmt.Println(errs)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
	return payURL
}
