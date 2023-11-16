package core

import (
	"ES/global"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func EsConnect() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("", ""), // 认证
	)
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	global.ESClient = client
}
