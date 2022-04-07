package main

import "fmt"

func main() {
	var n int = 0
	var i int = 0
	var largest1 int = 0
	var largest2 int = 0
	var temp int = 0

	fmt.Println("Enter the size of the array")
	fmt.Scan(&n)
	var array = make([]int, n)
	fmt.Println("Enter the elements")
	for i = 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	largest1 = array[0]
	largest2 = array[1]

	if largest1 < largest2 {
		temp = largest1
		largest1 = largest2
		largest2 = temp
	}

	for i = 2; i < n; i++ {
		if array[i] > largest1 {
			largest2 = largest1
			largest1 = array[i]
		} else if array[i] > largest2 && array[i] != largest1 {
			largest2 = array[i]
		}
	}
	fmt.Println(" First Largest ", largest1)
	fmt.Println(" Second Largest ", largest2)

}
