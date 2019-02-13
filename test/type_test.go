package test

import (
	"fmt"
	"reflect"
	"testing"
)

type testType int8

func TestType(test *testing.T) {
	var i int8 = 127
	var t1 testType = 127
	fmt.Println(i)
	fmt.Println(t1)
	n := reflect.TypeOf(i)
	n1 := reflect.TypeOf(t1)
	fmt.Println(n)
	fmt.Println(n1)
}
