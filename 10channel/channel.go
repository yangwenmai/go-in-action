package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	go func() {
		var sum int
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()
	fmt.Println(<-ch)

	one := make(chan int)
	two := make(chan int)
	go func() {
		one <- 100
	}()
	go func() {
		v := <-one
		two <- v
	}()
	fmt.Println(<-two)
}
