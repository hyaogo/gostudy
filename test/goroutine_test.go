package test

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	wg      sync.WaitGroup
	mutex   sync.Mutex
	counter int64
)

func TestGo(test *testing.T) {

	wg.Add(2)
	start := time.Now()
	go f1()
	go f1()
	wg.Wait()
	fmt.Println("耗时：", start.Sub(time.Now()))
}

func TestGo2(test *testing.T) {
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000000; i++ {
			atomic.AddInt64(&counter, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000000; i++ {
			atomic.AddInt64(&counter, 1)
			//counter++
		}
	}()
	wg.Wait()
	fmt.Println(counter)
}

func Test3(test *testing.T) {
	wg.Add(2)
	go counterAdd()
	go counterAdd()
	wg.Wait()
	fmt.Println(counter)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//无缓冲通道
func TestChan(test *testing.T) {
	ball := make(chan int)
	wg.Add(2)
	go Player("小明", ball)
	go Player("李磊", ball)
	ball <- 1
	wg.Wait()
}

func Player(name string, c chan int) {
	defer wg.Done()
	for {
		ball, ok := <-c
		if !ok {
			fmt.Printf("Player:%s win the game!\n", name)
			return
		}
		n := rand.Intn(100)
		if n > 90 {
			fmt.Printf("Player:%s is missed\n", name)
			close(c)
			return
		}
		ball++
		fmt.Printf("Player:%s hit %d\n", name, ball)
		c <- ball
	}
}

func f1() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func counterAdd() {
	defer wg.Done()
	mutex.Lock()
	for i := 0; i < 100000000; i++ {
		//atomic.AddInt64(&counter,1)
		counter++
	}
	mutex.Unlock()
}
