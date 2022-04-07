package main

import (
	"fmt"

	"example.com/stack"
)

func balanced() {

	var expr string
	fmt.Println("\nEnter an expression")
	fmt.Scanln(&expr)

	if check(expr) {
		fmt.Println("Expression is balanced")
	} else {
		fmt.Println("Expression is unbalanced")
	}

}

func check(expr string) bool {
	s := stack.Init(10)
	var i = 0
	var x = ' '
	for i = 0; i < len(expr); i++ {
		if expr[i] == '(' || expr[i] == '[' || expr[i] == '{' {
			s.Push(int(expr[i]))
			continue
		}

		if s.IsEmpty() {
			return false
		}

		switch expr[i] {
		case ')':

			x = rune(s.Peek())
			s.Pop()
			if x == '{' || x == '[' {
				return false
			}
		case '}':

			x = rune(s.Peek())
			s.Pop()
			if x == '(' || x == '[' {
				return false
			}
		case ']':

			x = rune(s.Peek())
			s.Pop()
			if x == '(' || x == '{' {
				return false

			}
		}
	}
	return s.IsEmpty()
}
