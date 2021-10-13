package main

import (
	"fmt"
)

var t1, t2, next_term int

func Fibo_Rec(inp int) int {

	if inp <= 1 {
		return inp
	}

	return Fibo_Rec(inp-1) + Fibo_Rec(inp-2)

}

func fibo_loop(inp int) {

	t1 = 0
	t2 = 1

	for i := 1; i <= inp; i++ {

		if i == 1 {
			fmt.Printf("%v", t1)
			continue
		}

		if i == 2 {
			fmt.Printf("%v", t2)
			continue
		}

		next_term = t1 + t2
		t1 = t2
		t2 = next_term
		fmt.Printf("%v", next_term)

	}
}

func main() {

	fmt.Println(Fibo_Rec(10))
	fibo_loop(10)

}

