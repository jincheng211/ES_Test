package main

import (
	"ES/core"
	"ES/docs"
)

func main() {
	core.EsConnect() // 连接
	// indexs.CreateIndex() // 创建
	docs.DocCreate() // 创建文档
	docs.DocFind()   // 查询
}
