// Package example is a example.
// http://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html
// unsafe.Pointer的4个规则:

// 1. 任何指针都可以转换为unsafe.Pointer
// 2. unsafe.Pointer可以转换为任何指针
// 3. uintptr可以转换为unsafe.Pointer
// 4. unsafe.Pointer可以转换为uintptr
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))

	*fp = *fp * 3
	fmt.Println(fp)

	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "小张"
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 22
	fmt.Println(*u)
}

type user struct {
	name string
	age  int
}
