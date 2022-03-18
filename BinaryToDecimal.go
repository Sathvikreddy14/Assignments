package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("enter no")
	fmt.Scanln(&num)
	fmt.Println(BinaryToDecimal(num))

}

func BinaryToDecimal(num int) int {
	var remainder int
	var i float64
	var decimal int
	for i = 0; num != 0; i++ {
		remainder = num % 10
		num = num / 10
		decimal = decimal + remainder*int(math.Pow(2, i))
	}
	return decimal
}
