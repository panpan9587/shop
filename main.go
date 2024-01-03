package main

import (
	"demo/config"
	"demo/model/elasticsearch"
	"encoding/json"
	"fmt"
)

func main() {
	Code := map[string]string{
		"code": "1234",
	}
	data, _ := json.Marshal(Code)
	fmt.Println(string(data))
	fmt.Println(config.IdentityConfig.SecretID)
	fmt.Println(config.IdentityConfig.SecretKey)
	//o, _ := mysql.GetOrderById(2)
	//url := common.AliPay(o)
	//fmt.Println(url)
	r := elasticsearch.GetEsGoods("商品1", "商品1")
	fmt.Println(r)
}
