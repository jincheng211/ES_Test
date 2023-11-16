package indexs

import (
	"ES/global"
	"context"
	"fmt"
)

// 判断索引是否存在
func ExistsIndex(index string) bool {
	exists, err := global.ESClient.IndexExists(index).Do(context.Background())
	if err != nil {
		fmt.Println("exists err:", err)
	}
	return exists
}
