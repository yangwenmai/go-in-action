package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 5)
	fmt.Printf("length:%d, cap:%d\n", len(slice1), cap(slice1))

	slice2 := make([]int, 5, 10)
	fmt.Printf("length:%d, cap:%d\n", len(slice1), cap(slice2))

	// slice3 := make([]int, 5, 4) // error

	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("length:%d, cap:%d\n", len(slice4), cap(slice4))

	slice5 := []int{5: 5}
	fmt.Printf("length:%d, cap:%d\n", len(slice5), cap(slice5))

	fmt.Println("-------------")

	slice := []int{1, 2, 3, 4, 5}
	slice6 := slice[:]
	slice7 := slice[0:]
	slice8 := slice[:5]
	slice9 := slice[2:5]
	fmt.Println(slice6)
	fmt.Println(slice7)
	fmt.Println(slice8)
	fmt.Println(slice9)
	// slice10 := slice[2:6] // panic: runtime error: slice bounds out of range
	// fmt.Println(slice10)

	slice = []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	newSlice[1] = 10
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Printf("length:%d, cap:%d\n", len(newSlice), cap(newSlice))

	newSlice2 := slice[1:2:3]
	fmt.Println(newSlice2)
	fmt.Printf("length:%d, cap:%d\n", len(newSlice2), cap(newSlice2))

	// newSlice2 = slice[1:2:6] // panic: runtime error: slice bounds out of range

	slice[2] = 3
	fmt.Println(slice[2])
	slice[2] = 10
	fmt.Println(slice[2])

	fmt.Println("------------")

	slice = []int{1, 2, 3, 4, 5}
	newSlice = slice[1:3]

	newSlice = append(newSlice, 10)
	fmt.Println(newSlice)
	fmt.Println(slice)

	newSlice = append(newSlice, 10, 20)
	fmt.Println(newSlice)

	newSlice = append(newSlice, slice...)
	fmt.Println(newSlice)
	fmt.Println(slice)

	fmt.Println("----------")

	slice = []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Printf("idx:%d, val:%d\n", i, v)
		v = i * 2
		fmt.Println(v)
	}

	fmt.Println(slice)

	fmt.Println("----------")

	slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)

	modify(slice)
	fmt.Println(slice)
}

func modify(a []int) {
	fmt.Printf("%p\n", &a)
	a[2] = 10
}
