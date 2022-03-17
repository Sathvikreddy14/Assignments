package main

import "fmt"

func main() {
	
	var num int
	num = 8
	Fibonacci(num)
}
func Fibonacci(num int) {
	var t1 int=0
	var t2 int=1
  var next int
	for i := 0; i < num-1; i++ {
		
		next = t1 + t2
		t1 = t2
		t2 = next
        
	}
    fmt.Println(t1)
	
}
