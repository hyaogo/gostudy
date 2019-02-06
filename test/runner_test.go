package test

import (
	"fmt"
	"github.com/goinaction/code/chapter7/patterns/runner"
	"log"
	"os"
	"testing"
	"time"
)

const timeout = 10 * time.Second

func TestRunner(test *testing.T) {
	log.Println("Starting work.")
	start := time.Now()
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("run timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
	log.Println("耗时：", start.Sub(time.Now()))
}

func createTask() func(int) {
	return func(i int) {
		fmt.Println("running ........", i)
		time.Sleep(1 * time.Second)
	}
}
