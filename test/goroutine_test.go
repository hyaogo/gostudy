package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func TestGo(test *testing.T) {

	wg.Add(2)
	start := time.Now()
	go f1()
	go f1()
	wg.Wait()
	fmt.Println("耗时：", start.Sub(time.Now()))
}

func f1() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
