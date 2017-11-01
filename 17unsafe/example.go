// Package example is a example.
// http://www.flysnow.org/2017/07/02/go-in-action-unsafe-memory-layout.html
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))

	var b bool
	var i8 int8
	var i16 int16
	var i64 int64
	var f32 float32
	var s string
	var m map[string]string
	var p *int32
	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Alignof(i8))
	fmt.Println(unsafe.Alignof(i16))
	fmt.Println(unsafe.Alignof(i64))
	fmt.Println(unsafe.Alignof(f32))
	fmt.Println(unsafe.Alignof(s))
	fmt.Println(unsafe.Alignof(m))
	fmt.Println(unsafe.Alignof(p))
	// unsafe.Alignof(x)等价于reflect.TypeOf(x).Align()
	fmt.Println(reflect.TypeOf(p).Align())

	fmt.Println("------------------------")
	var u user
	fmt.Println(unsafe.Offsetof(u.b))
	fmt.Println(unsafe.Offsetof(u.i))
	fmt.Println(unsafe.Offsetof(u.j))
	fmt.Println(reflect.TypeOf(u).Field(2).Offset)

	// unsafe.Offsetof(u1.i)等价于reflect.TypeOf(u1).Field(1).Offset

}

type user struct {
	b byte
	i int32
	j int64
}
