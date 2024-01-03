package redis

import "fmt"

//商品添加虚拟数据

func AddGoodsList(goodsId int64, stock int64) {
	key := fmt.Sprintf("stock:goodsId_%d", goodsId)
	for i := 0; i < int(stock); i++ {
		Client.RPush(key, goodsId)
	}
}

// 商品查询虚拟库存

func LenGoodsList(goodsId int64) int64 {
	key := fmt.Sprintf("stock:goodsId_%d", goodsId)
	return Client.LLen(key).Val()
}

func DeclGoodsList(goodsId int64, stock int64) {
	key := fmt.Sprintf("stock:goodsId_%d", goodsId)
	for i := 0; i < int(stock); i++ {
		Client.LPop(key).Result()
	}
}
