package docs

import (
	"ES/global"
	"ES/models"
	"context"
	"fmt"
)

func DocUpdate(id string) {
	res, err := global.ESClient.Update().Index(models.UserModel{}.Index()).
		Id(id).Doc(map[string]any{
		"user_name": "你好枫枫",
	}).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", res)
}
