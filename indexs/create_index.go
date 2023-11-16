package indexs

import (
	"ES/global"
	"ES/models"
	"context"
	"fmt"
)

func CreateIndex() {
	index := "user_index"
	// 判断索引是否存在
	if ExistsIndex(index) {
		DeleteIndex(index)
	}

	createIndex, err := global.ESClient.CreateIndex(index).
		BodyString(models.UserModel{}.Mapping()).Do(context.Background())
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	fmt.Println("索引创建成功", createIndex)
}
