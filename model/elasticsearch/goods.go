package elasticsearch

import (
	"bytes"
	"context"
	"demo/model/mysql"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"strconv"
	"strings"
)

// 添加商品到el搜索
func AddEsGoods(goods *mysql.Goods, index string) {
	data, _ := json.Marshal(goods)
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: strconv.Itoa(int(goods.ID)),
		Body:       strings.NewReader(string(data)),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), Es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}

// 根据索引获取es商品
func GetEsGoods(index, search string) (r map[string]interface{}) {
	//多层map嵌套
	//定义查询DSL
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"goods_title": search,
			},
		},
	}
	json.NewEncoder(&buf).Encode(query)
	res, err := Es.Search(
		Es.Search.WithContext(context.Background()),
		Es.Search.WithIndex(index),
		Es.Search.WithBody(&buf),
		Es.Search.WithTrackTotalHits(true),
		Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
	return r
}
