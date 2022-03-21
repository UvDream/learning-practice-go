package main

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	address := []string{"http://10.100.1.5:9200/"}
	config := elasticsearch.Config{

		Addresses: address,
	}
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "连接es失败")
	fmt.Println("连接成功", es)
	res, err := es.Info()
	failOnError(err, "获取es信息失败")
	fmt.Println(res.String())

	// 查询
	type visitor struct {
		Vid int `json:"vid"`
	}
	res, err = es.Get("visitors", "1")
	failOnError(err, "查询失败")
	fmt.Println("查询结果")
	fmt.Println(res.String())
}
