package mysql

import (
	"demo/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MyDB *gorm.DB

func init() {
	dns := config.MysqlDB
	fmt.Println(dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	MyDB = db
	db.AutoMigrate(&User{}, &Goods{}, &Order{}, &OrderGoodsInfo{}, &Identity{}, &ShoppingCart{})
}
