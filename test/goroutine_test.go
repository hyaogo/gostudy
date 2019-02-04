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

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func Test4(test *testing.T) {
	wg.Add(1)
	baton := make(chan int)
	go Runner(baton)
	//开始
	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {

	var newRunner int

	//等待接力棒
	runner := <-baton
	fmt.Printf("Runner %d 拿到接力棒开始跑\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d 走上跑道\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(1000 * time.Millisecond)

	if runner == 4 {
		fmt.Println("比赛结束:", runner)
		wg.Done()
		return
	}
	fmt.Printf("Runner %d 将接力棒交给下一位 Runner %d\n", runner, newRunner)
	baton <- newRunner
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
		if n > 98 {
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
