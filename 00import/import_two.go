package main

import (
	"fmt"
	_ "net/http"
	"strings"
)

func main() {
	cnt := strings.Count("Hello", "Hello")
	fmt.Printf("Hello, Go! \nIt Contain %d Hello.\n", cnt)
}
