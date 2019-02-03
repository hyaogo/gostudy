package test

import (
	"fmt"
	"testing"
)

func TestMap(test *testing.T) {
	var m1 = make(map[string]string)
	m1["name"] = "xiaoming"
	fmt.Println(m1["name"])
	m2 := map[string]int{}
	m2["age"] = 88
	fmt.Println(m2["age"])

	name, exists := m1["fd"]
	fmt.Println("name:", name)
	fmt.Println("exists:", exists)

	for key, value := range m1 {
		fmt.Printf("key:%s,value:%s\n", key, value)
	}

	delete(m2, "age")
	fmt.Println("m2:", m2)

	for key, value := range m2 {
		fmt.Println("***", key, value)
	}

	ParamMap(m1)
	fmt.Println("m1:", m1)
}

//函数传递map
func ParamMap(m map[string]string) {
	m["hello"] = "hello"
}
