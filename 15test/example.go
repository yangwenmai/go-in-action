// Package example is a example.
package example

import (
	"fmt"
)

// Add return a + b
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(3, 4))
}
