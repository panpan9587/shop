package mysql

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId      int64   `gorm:"type:int(11);not null;"`
	OrderSn     string  `gorm:"type:varchar(30);not null;unique"`
	GoodsId     int64   `gorm:"type:int(11);not null;"` //单件商品
	GoodsNum    int64   `gorm:"type:int(11);not null;"`
	TotalPrice  float64 `gorm:"type:decimal(10,2);not null;"`
	PayType     int64   `gorm:"type:int(11);not null;"`
	OrderStatus int64   `gorm:"type:int(11);not null;"`
	RealName    string  `gorm:"type:varchar(50);not null;"` //收件人姓名
	Addr        string  `gorm:"type:varchar(50);not null;"` //收货人地址
}

// 订单商品记录
type OrderGoodsInfo struct {
	GoodsId     int64
	GoodsPrice  float64
	GoodsNum    int64
	GoodsStatus int64
}

// 根据订单id查询订单
func GetOrderById(orderId int64) (order *Order, err error) {
	err = MyDB.Where("id = ?", orderId).First(&order).Error
	return
}

// 修改订单状态
func UpdateOrderStatus(orderId, status int64) (err error) {
	err = MyDB.Model(&Order{}).Where("id = ?", orderId).Update("order_status", status).Error
	return
}

// 根据订单编号获取订单
func GetOrderByOrderSn(orderSn string) (order *Order, err error) {
	err = MyDB.Where("order_sn = ?", orderSn).First(&order).Error
	return
}

// 创建订单
func CreateOrder(order *Order) (err error) {
	err = MyDB.Create(&order).Error
	return
}
