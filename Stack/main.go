package main

import (
	"fmt"
)

func main() {

	a := []int{}

	for i := 0; i <= 10; i++ {

		a = append(a, i)
	}

	fmt.Println(a)
	
	for i:=len(a)-1;i>0;i--{
	    fmt.Println(a[i])
	
	}

}

