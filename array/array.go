package main

import "fmt"

func main() {
	var array1 [5]int

	var array2 [5]int
	array2 = [5]int{1, 2, 3, 4, 5}

	array3 := [5]int{1, 2, 3, 4, 5}

	array4 := [...]int{1, 2, 3, 4, 5}

	array5 := [5]int{0, 1, 0, 4, 0}
	array6 := [5]int{1: 1, 3: 4}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(array6)

	fmt.Printf("array1 index:2 %d\n", array1[2])
	array1[2] = 1
	fmt.Printf("array1 index:2 %d\n", array1[2])

	for i := 0; i < 5; i++ {
		fmt.Printf("index: %d, val: %d\n", i, array6[i])
	}
	fmt.Println("-----------------")
	for i, v := range array6 {
		fmt.Printf("index: %d, val: %d\n", i, v)
	}

	// var array7 [5]int = array6 // success
	// var array8 [4]int = array6 // error

	array8 := [5]*int{1: new(int), 3: new(int)}

	array8[0] = new(int)
	*array8[0] = 0
	*array8[1] = 1

	array9 := [5]int{1: 2, 3: 4}
	modifyArray(array9)
	fmt.Println(array9)
	fmt.Println("-------------")
	modifyArray2(&array9)
	fmt.Println(array9)
}

func modifyArray(a [5]int) {
	a[1] = 3
	fmt.Println(a)
}

// 数组的指针和指针数组是两个概念，数组的指针是*[5]int,指针数组是[5]*int，注意*的位置。
func modifyArray2(a *[5]int) {
	a[1] = 3
	fmt.Println(*a)
}
