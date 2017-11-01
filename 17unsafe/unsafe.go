// Package example is a example.
// http://www.flysnow.org/2017/07/02/go-in-action-unsafe-memory-layout.html
// 不同的字段顺序，最终决定struct的内存大小，所以有时候合理的字段顺序可以减少内存的开销。
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6
	fmt.Println("u1 size is ", unsafe.Sizeof(u1)) // 16
	fmt.Println("u2 size is ", unsafe.Sizeof(u2)) // 24
	fmt.Println("u3 size is ", unsafe.Sizeof(u3)) // 16
	fmt.Println("u4 size is ", unsafe.Sizeof(u4)) // 24
	fmt.Println("u5 size is ", unsafe.Sizeof(u5)) // 16
	fmt.Println("u6 size is ", unsafe.Sizeof(u6)) // 16
}

type user1 struct {
	b byte
	i int32
	j int64
}
type user2 struct {
	b byte
	j int64
	i int32
}
type user3 struct {
	i int32
	b byte
	j int64
}
type user4 struct {
	i int32
	j int64
	b byte
}
type user5 struct {
	j int64
	b byte
	i int32
}
type user6 struct {
	j int64
	i int32
	b byte
}
