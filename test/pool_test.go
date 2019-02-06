package test

import (
	"github.com/goinaction/code/chapter7/patterns/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxGoroutines = 25
	poolResources = 2
)

type dbConnection struct {
	id int32
}

func (db *dbConnection) Close() error {
	log.Println("Close connection:", db.id)
	return nil
}

var idCounter int32

//工厂函数
func createDbConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create new connection:", id)
	return &dbConnection{id}, nil
}

func TestPool(test *testing.T) {
	var wg sync.WaitGroup

	wg.Add(maxGoroutines)
	p, err := pool.New(createDbConnection, poolResources)
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < maxGoroutines; i++ {
		go func(q int) {
			querys(q, p)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func querys(q int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println("err")
		return
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", q, conn.(*dbConnection).id)
}
