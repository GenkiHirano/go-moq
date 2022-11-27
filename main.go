package main

//go:generate moq -out goimports -out mockexample_test.go . Example

import "fmt"

type Args struct {
	a int
	b string
	c []byte
}

type Returns struct {
	e int
	f string
	g []byte
}

type Example interface {
	Helle(Args) Returns
}

func main() {
	fmt.Println("hello, world!")
}
