package main

import (
	"fmt"
	"runtime"
	"sync"
	// "sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count) // 优化之前可能是2，可能是3，可能是4；优化之后可能是2，可能是4
	// 我们对于同一个资源的读写必须是原子化的，也就是说，同一时间只能有一个goroutine对共享资源进行读写操作。
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		// value := count
		// 优化1：原子操作
		// value := atomic.LoadInt32(&count)
		// 优化2：加锁
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		// count = value
		// 优化1：原子操作
		// atomic.StoreInt32(&count, value)
		// 优化2：加锁
		count = value
		mutex.Unlock()
	}
}
