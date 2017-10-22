package main

import (
	"fmt"
)

func main() {
	ad := admin{user{"小张", "xiaozhang@gmail.com"}, "管理员"}
	fmt.Println("可以直接调用,名字为：", ad.name)
	fmt.Println("可以直接调用,邮箱为：", ad.email)
	fmt.Println("也可以通过内部类型调用,名字为：", ad.user.name)
	fmt.Println("也可以通过内部类型调用,邮箱为：", ad.user.email)
	fmt.Println("但是新增加的属性只能直接调用，级别为：", ad.level)

	ad.user.sayHello()
	ad.sayHello()

	// 如果内部类型实现了某个接口，那么外部类型也被认为实现了这个接口。
	sayHello(ad.user)
	sayHello(ad)
}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

// Hello hello interface
type Hello interface {
	hello()
}

func sayHello(h Hello) {
	h.hello()
}

func (u user) hello() {
	fmt.Println("Hello, i am a user")
}

func (a admin) hello() {
	fmt.Println("Hello, i am a admin")
}

func (u user) sayHello() {
	fmt.Println("Hello, i am a user")
}

func (a admin) sayHello() {
	fmt.Println("Hello, i am a admin")
}

// 以下是 io 标准包的定义。

// Reader reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer writer
type Writer interface {
	Write(p []byte) (n int, err error)
}

// Closer closer
type Closer interface {
	Close() error
}

// ReadWriter readwriter
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser readcloser
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser writecloser
type WriteCloser interface {
	Writer
	Closer
}
