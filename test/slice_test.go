package test

import (
	"fmt"
	"testing"
)

func TestSlice(test *testing.T) {
	s1 := make([]string, 5)
	fmt.Println("s1:", s1)
	fmt.Println("len(s1):", len(s1))
	fmt.Println("cap(s1):", cap(s1))

	s2 := []int{10, 20, 30, 40, 50}
	s3 := s2[1:3]
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	s3[0] = 77
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	s3 = append(s3, 99)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	//多维切片
	s4 := [][]int{{10}, {100, 200}}
	fmt.Println("s4:", s4)
	fmt.Println("s4[0][0]:", s4[0][0])
	s4[1] = append(s4[1], 300)
	fmt.Println("s4:", s4)
	for _, v := range s4 {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
	SliceParam(s3)
}

func SliceParam(s []int) {
	for _, v := range s {
		fmt.Println("我在SliceParam里被打印:", v)
	}
}

func TestSliceLen(test *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 6}
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
	s2 := s[:5]
	fmt.Println("s2:", s2)
	fmt.Println("len(s2):", len(s2))
	fmt.Println("cap(s2):", cap(s2))
}
