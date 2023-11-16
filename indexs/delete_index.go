package indexs

import (
	"ES/global"
	"context"
	"fmt"
)

func DeleteIndex(index string) {
	_, err := global.ESClient.DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Println("delete err:", err)
		return
	}
	fmt.Println("索引删除成功")
}
