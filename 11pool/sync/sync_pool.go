package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutine = 5
)

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	p := &sync.Pool{New: createConnection}
	//模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
}

// dbQuery 模拟数据库查询
func dbQuery(query int, pool *sync.Pool) {
	conn := pool.Get().(*dbConnection)
	defer pool.Put(conn)
	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.ID)
}

// dbConnection 数据库连接
type dbConnection struct {
	ID int32 //连接的标志
}

// Close 实现io.Closer接口
func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

// createConnection 生成数据库连接的方法，以供资源池使用
func createConnection() interface{} {
	//并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{ID: id}
}
