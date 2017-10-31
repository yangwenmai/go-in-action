// Package example is a example.
package main

import (
	"fmt"
	"reflect"
)

// User struct
type User struct {
	Name string
	Age  int
}

// Print print
func (u User) Print(prefix string) {
	fmt.Printf("%s: Name is %s, Age is %d\n", prefix, u.Name, u.Age)
}

func main() {
	u := User{
		Name: "小张",
		Age:  22,
	}

	vx := reflect.ValueOf(u)
	mPrint := vx.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))
	fmt.Println(mPrint.IsValid())
	t := reflect.TypeOf(u)
	fmt.Println(t)
	v := reflect.ValueOf(u)
	fmt.Println(v)

	fmt.Printf("%T\n", u)

	fmt.Printf("%v\n", v)

	u1 := v.Interface().(User)
	fmt.Println(u1)

	t1 := v.Type()
	fmt.Println(t1)

	fmt.Println(t.Kind())

	fmt.Println("-------------------------------")
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
		fmt.Println(t.Field(i).Type)
		fmt.Println(t.Field(i).Tag)
		fmt.Println(t.Field(i).Anonymous)
		fmt.Println(t.Field(i))
	}

	fmt.Println("-------------------------------")

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
		fmt.Println(t.Method(i).Type)
		fmt.Println(t.Method(i))
	}

	fmt.Println("===============================")

	x := 2
	v2 := reflect.ValueOf(&x)
	v2.Elem().SetInt(100)
	fmt.Println(x)
}
