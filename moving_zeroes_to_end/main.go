package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Println(a)

	count := 0

	for index := 0; index < len(a); index++ {

		if a[index] != 0 {
			a[count] = a[index]
			count++
		}
	}

	for count != len(a) {
		a[count] = 0
		count = count + 1

	}

	fmt.Println(a)
}

