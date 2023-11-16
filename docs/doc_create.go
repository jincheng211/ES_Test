package docs

import (
	"ES/global"
	"ES/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

func DocCreate() {
	user := models.UserModel{
		ID:       12,
		UserName: "lisi",
		NickName: "夜空中最亮的lisi",
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	indexResponse, err := global.ESClient.Index().Index(user.Index()).BodyJson(user).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", indexResponse)
}

func DocCreateBatch(modelLists []models.UserModel) {
	bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, model := range modelLists {
		req := elastic.NewBulkCreateRequest().Doc(model)
		bulk.Add(req)
	}

	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Succeeded())
}
