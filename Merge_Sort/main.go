package main

import (
	"fmt"
)

func MergeSort(items []int) []int {

	num := len(items)

	if num == 1 {
		return items
	}
	middle := int(num / 2)
	var left = make([]int, middle)
	var right = make([]int, num-middle)

	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]

		} else {
			right[i-middle] = items[i]
		}

	}

	return merge(MergeSort(left), MergeSort(right))

}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}

		i++

	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result

}

func main() {

	a := []int{99, 56, 43, 27, -1, -8, -6, 0, 90, 100,9}

	fmt.Println(MergeSort(a))

}

