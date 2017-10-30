package main

import (
	"fmt"
	"sync"
)

func main() {
	sm := NewSynchronizedMap()
	sm.Put("1", 2)
	sm.Put("2", 2)
	sm.Put(1, "1")
	sm.Put(2, "2")
	sm.Each(func(k, v interface{}) {
		fmt.Println(k, v)
	})
}

// SynchronizedMap struct
type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

// Put put
func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k] = v
}

// Get get
func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}

// Delete 删除
func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data, k)
}

// Each 循环遍历
func (sm *SynchronizedMap) Each(cb func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	for k, v := range sm.data {
		cb(k, v)
	}
}

// NewSynchronizedMap New
func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}
