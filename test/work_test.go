package test

import (
	"github.com/goinaction/code/chapter7/patterns/work"
	"log"
	"sync"
	"testing"
	"time"
)

var names = []string{
	"小明",
	"丽丽",
	"李磊",
	"丹尼尔",
}

type namePrinter struct {
	name string
}

func (name *namePrinter) Task() {
	log.Println(name)
	time.Sleep(time.Second)
}

func TestWord(test *testing.T) {
	w := work.New(100)
	var wg sync.WaitGroup
	wg.Add(len(names) * 100)
	for i := 0; i < 100; i++ {
		for _, n := range names {
			np := namePrinter{name: n}
			go func() {
				w.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	w.Shutdown()
}
