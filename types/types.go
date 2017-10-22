package main

import "fmt"

func main() {
	name := "小张"
	fmt.Println(modify(name))
	fmt.Println(name)

	ages := map[string]int{"小张": 22}
	fmt.Println(ages)
	modifyAges(ages)
	fmt.Println(ages)

	fmt.Println("--------------")

	jim := Person{"Jim", 22, map[string]int{"平均成绩": 88}}
	fmt.Println(jim)
	modifyPerson(jim)
	// 年龄没有变
	fmt.Println(jim)

	// 年龄被修改了
	modifyPerson2(&jim)
	fmt.Println(jim)

	modifyPerson3(jim)
	// 平均成绩被修改了
	fmt.Println(jim)

	fmt.Println("----------------")
	myPerson := MyPerson{"小李", 22, map[string]int{"y": 2}}
	// var person Person
	// person = myPerson // cannot use myPerson (type MyPerson) as type Person in assignment
	// fmt.Println(person)
	fmt.Println(myPerson)
}

func modify(s string) string {
	s += s
	return s
}

func modifyAges(m map[string]int) {
	m["小张"] = 12
}

func modifyPerson(p Person) {
	p.age = p.age + 10
}

func modifyPerson2(p *Person) {
	p.age = p.age + 10
}

func modifyPerson3(p Person) {
	p.m["平均成绩"] = 99
}

type Person struct {
	name string
	age  int
	m    map[string]int
}

// MyPerson MyPerson
type MyPerson Person
