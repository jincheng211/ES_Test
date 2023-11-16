package docs

import (
	"ES/global"
	"ES/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// DocDelete 单个删除(根据id)
func DocDelete(ID string) {
	deleteResponse, err := global.ESClient.Delete().Index(models.UserModel{}.Index()).Id(ID).Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse)
}

// DocDeleteBatch 删除多个文件(根据id)
func DocDeleteBatch(idList []string) {
	bulk := global.ESClient.Bulk().Index(models.UserModel{}.Index()).Refresh("true")
	for _, s := range idList {
		req := elastic.NewBulkDeleteRequest().Id(s)
		bulk.Add(req)
	}

	res, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Succeeded())
}
