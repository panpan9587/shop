package elasticsearch

import (
	"demo/config"
	"fmt"
	"github.com/elastic/go-elasticsearch"
)

var Es *elasticsearch.Client

func init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + config.ElasticsearchDB,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("es连接失败")
	}
	Es = es
	// ...
}
