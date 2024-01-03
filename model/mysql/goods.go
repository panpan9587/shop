package mysql

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsTitle  string  `gorm:"type:varchar(100);not null;"`
	GoodsPrice  float64 `gorm:"type:decimal(10,2);not null;"`
	GoodsStock  int64   `gorm:"type:int(11);not null;"`
	GoodsStatus int64   `gorm:"type:int(11);not null;"`
}

// 添加商品
func AddGoods(goods *Goods) (err error) {
	err = MyDB.Create(&goods).Error
	return
}

// 根据商品id查询商品信息
func GetGoodsById(goodsId int64) (goods *Goods, err error) {
	err = MyDB.Where("id = ?", goodsId).First(&goods).Error
	return
}
