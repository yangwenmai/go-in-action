package main

import "github.com/yangwenmai/go-in-action/08identifier_export/common"
import "fmt"

func main() {
	// c := common.count(10) // cannot refer to unexported name common.count
	// fmt.Println(c)

	c := common.New(13)
	fmt.Println(c)

	l := common.NewLoginer()
	l.Login()
}
