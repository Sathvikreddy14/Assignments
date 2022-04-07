package main

import "fmt"

func main() {

	var str1 string
	var str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)

	if anagram(str1, str2) {
		fmt.Println("Anagrams")
	} else {
		fmt.Println("Not Anagrams")
	}

}

func anagram(str1 string, str2 string) bool {

	var count [256]int
	var i = 0
	if len(str1) != len(str2) {
		return false
	}

	for ; i < len(str1); i++ {
		count[int(str1[i])]++
		count[int(str2[i])]--

	}
	for i = 0; i < 256; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true

}
