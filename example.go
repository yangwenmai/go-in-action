package example

import (
	"fmt"
)

// Add return a + b
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(1, 3))
}
