package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var b bytes.Buffer
	fmt.Fprint(&b, "Hello World.")
	fmt.Println(b.String())
	var w io.Writer
	w = &b
	fmt.Println(w)

	// 多态
	var a animal

	var c cat
	a = c
	a.printInfo()

	var d dog
	a = d
	a.printInfo()

	// other
	// 实体类型以值接收者实现接口的时候，不管是实体类型的值，还是实体类型值的指针，都实现了该接口。
	var e cat
	invoke(e)
	invoke(&e)

	// other
	// 实体类型以指针接收者实现接口的时候，只有指向这个类型的指针才被认为实现了该接口。
	var f fish
	invoke(&f)
	// 类型的值只能实现值接收者的接口；指向类型的指针，既可以实现值接收者的接口，也可以实现指针接收者的接口。
}

func invoke(a animal) {
	a.printInfo()
}

type animal interface {
	printInfo()
}

type cat int
type dog int

func (c cat) printInfo() {
	fmt.Println("a cat")
}

func (d dog) printInfo() {
	fmt.Println("a dog")
}

type fish int

func (f *fish) printInfo() {
	fmt.Println("a fish")
}
