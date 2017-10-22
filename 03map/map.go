package main

import (
	"fmt"
	"sort"
)

func main() {
	dict := make(map[string]int)
	dict["小张"] = 28
	fmt.Println(dict)

	dict = map[string]int{"小李": 29}
	fmt.Println(dict)

	dict = map[string]int{"小李": 29, "小王": 22}
	fmt.Println(dict)

	dict["小王"] = 23
	dict["小赵"] = 25
	fmt.Println(dict)

	age, exists := dict["小王"]
	fmt.Printf("is exists:%v, age:%d\n", exists, age)

	delete(dict, "小杨")
	fmt.Println(dict)

	for key, val := range dict {
		fmt.Println(key, val)
	}

	fmt.Println("-------------------")

	var names []string
	for key := range dict {
		names = append(names, key)
	}
	sort.Strings(names)
	for _, key := range names {
		fmt.Println(key, dict[key])
	}

	fmt.Println("-------------------")

	dict = map[string]int{"小五": 20, "小三": 23}
	modify(dict)
	fmt.Println(dict["小三"])
}
func modify(a map[string]int) {
	a["小三"] = 25
}
