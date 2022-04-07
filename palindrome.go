package main

import "fmt"

var reverse = 0

func check2(Num int) int {
	var remainder int

	for ; Num > 0; Num = Num / 10 {
		remainder = Num % 10
		reverse = reverse*10 + remainder
	}
	return reverse
}
func main() {

	var Num int

	fmt.Scanln(&Num)

	reverse = check2(Num)
	fmt.Println(" Reverse ", reverse)

	if Num == reverse {
		fmt.Println(Num, " Palindrome")
	} else {
		fmt.Println(Num, "  Not Palindrome ")
	}
}
