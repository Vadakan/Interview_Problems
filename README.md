# Interview_Problems

we cannot pass more than 2 argument in slice append function

either you can pass the slice unfurled or number of arguments comma separated to append 

below example wont work

package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	a = append(a[0:2], a[3:5], a[6:]...)

	fmt.Println(a)
}
