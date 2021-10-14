package main

import "fmt"

func main() {

	Queue := []int{}

	for i := 0; i < 10; i++ {
		Queue = append(Queue, i)
	}

	for i := 0; i < len(Queue); i++ {
		fmt.Println(Queue[i])
	}

}

