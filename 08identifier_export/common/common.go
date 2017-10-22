package common

import (
	"fmt"
)

type count int

// New new count
func New(v int) count {
	return count(v)
}

// NewLoginer new loginer
func NewLoginer() Loginer {
	return defaultLogin(0)
}

// Loginer interface
type Loginer interface {
	Login()
}

type defaultLogin int

func (d defaultLogin) Login() {
	fmt.Println("login in...")
}
