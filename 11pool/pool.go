package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// ErrPoolClosed 资源池关闭
var ErrPoolClosed = errors.New("资源池已经关闭。。。")

const (
	maxGoroutine = 5
	poolRes      = 2
)

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p, err := New1(createConnection, poolRes)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}

func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第 %d 个查询，使用的是ID为 %d 的数据库连接\n", query, conn.(*dbConnection).ID)
}

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}

// Pool pool
type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

// New1 new
func New1(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size 的值太小了。。。")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

// Acquire acquire
func (p *Pool) Acquire() (io.Closer, error) {

	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

// Close close
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	close(p.res)
	for r := range p.res {
		r.Close()
	}
}

// Release release
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源放到池子里面了")
	default:
		log.Println("资源池满了，释放这个资源吧。")
	}
}
