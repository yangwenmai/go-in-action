// Package functions 介绍
// 其他语言中，比如Java，一般来说，函数就是方法，方法就是函数。
// 在Go语言中，函数是指不属于任何结构体、类型的方法，也就是说，函数是没有接收者的；而方法是有接收者的，我们说的方法要么是属于一个结构体的，要么属于一个新定义的类型的。
package main

import (
	"fmt"
	"github.com/yangwenmai/go-in-action/05functions/lib"
	"log"
	"os"
)

func main() {
	sum, _ := add(1, 3)
	fmt.Println(sum)

	sum = lib.Add(2, 3)
	fmt.Println(sum)

	fmt.Println("------------")

	// Go语言里有两种类型的接收者：值接收者和指针接收者。
	p := Person{"小吴"}
	fmt.Println(p.name)
	p.Modify()
	fmt.Println(p.String())
	fmt.Println(p.name)
	fmt.Println("------------")
	fmt.Println(p.name)
	p.Modify2() // 一样的：(&p).Modify2()
	fmt.Println(p.String())
	fmt.Println(p.name)

	file, err := os.Open("/tmp")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file.Name())

	//
	myPrint("2", "+", "2", "=", "?")
}

func myPrint(a ...interface{}) {
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()
}

func add(a, b int) (int, error) {
	return a + b, nil
}

// Person person:{name}
type Person struct {
	name string
}

// String string
func (p Person) String() string {
	return "the person's name is " + p.name
}

// Modify modify
func (p Person) Modify() {
	p.name = "小小"
}

// Modify2 modify2
func (p *Person) Modify2() {
	p.name = "小小"
}
