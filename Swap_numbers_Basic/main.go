package main

import (
	"fmt"
)

func main() {
	var a, b int
	a = 10
	b = 20
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

