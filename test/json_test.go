package test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

var RD = `{ 
	"name":"三国演义",
	"title":"历史"
}`

type Book struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

//反序列化
func TestJson1(test *testing.T) {
	var book Book
	err := json.Unmarshal([]byte(RD), &book)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(book)
}

//序列化
func TestJson2(test *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "三国演义"
	m["title"] = "历史"
	m["leader"] = map[string]interface{}{
		"name": "曹操",
		"sex":  "男",
	}

	data, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(data))
}
