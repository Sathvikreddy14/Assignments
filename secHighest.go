package main

import (
	"fmt"
	"math"
)

func main() {
	var elems int
	fmt.Println("Enter Size of array")
	fmt.Scan(&elems)
	var array = make([]int, elems)

	for i := 0; i < elems; i++ {
		fmt.Scan(&array[i])
	}

	first := math.MinInt
	second := math.MinInt
	for i := 1; i < len(array); i++ {
		if first < array[i] {
			second = first
			first = array[i]
		} else if second < array[i] {
			second = array[i]
		}
	}
	fmt.Println("Second largest element is: ", second)
}
