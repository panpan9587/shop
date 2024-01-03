package mysql

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	UserId     int64   `gorm:"type:int(11);not null;"`
	GoodsId    int64   `gorm:"type:int(11);not null;"`
	GoodsNum   int64   `gorm:"type:int(11);not null;"`
	GoodsPrice float64 `gorm:"type:decimal(10,2);not null;"`
}

// 添加购物车
func AddShoppingCart(shop *ShoppingCart) (err error) {
	err = MyDB.Create(&shop).Error
	return
}

// 添加商品库存
func AddShoppingCartTo(shop *ShoppingCart) (err error) {
	err = MyDB.Where(shop).Save(&shop).Error
	return
}

// 根据商品id查询购物车数据
func GetShoppingCartByGoodsId(goodsId int64) (shop *ShoppingCart, err error) {
	err = MyDB.Where("goods_id = ?", goodsId).First(&shop).Error
	return
}

// 根据id删除购物车商品
func RemoveShoppingCart(shopId int64) (err error) {
	err = MyDB.Where("id = ?", shopId).Delete(&ShoppingCart{}).Error
	return
}

// 清空购物车
func RemoveAllShoppingCartByUserId(userId int64) (err error) {
	err = MyDB.Where("user_id = ?", userId).Delete(&ShoppingCart{}).Error
	return
}

// 查询购物车中的所有商品
func GetShoppingCartListByUserId(userId int64) (list []ShoppingCart, err error) {
	err = MyDB.Where("user_id = ?", userId).Find(&list).Error
	return
}
