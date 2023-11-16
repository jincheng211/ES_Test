package docs

import (
	"ES/global"
	"ES/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func DocFind() {
	query := elastic.NewBoolQuery()
	// 精确匹配(针对keyword字段)
	// query := elastic.NewTermQuery("user_name", "fengfeng")

	// 模糊匹配(主要是查text，也能查keyword
	// 模糊匹配keyword字段，是需要查完整的
	// 匹配text字段则不用，搜完整的也会搜出很多)
	// query := elastic.NewMatchQuery("nick_name", "夜空中最亮的枫枫")
	res, err := global.ESClient.Search(models.UserModel{}.Index()).Query(query).From(0).Size(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	count := res.Hits.TotalHits.Value
	fmt.Println(count)
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}
