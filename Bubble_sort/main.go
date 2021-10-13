package main

import "fmt"

func main() {

	a := []int{4, 5, 2, 1, 67, 43, 28, 82, 67, -1, -2}

	for i := len(a); i > 0; i-- {
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]

			}

		}

	}
	fmt.Println("Sorted Numbers :", a)

}

