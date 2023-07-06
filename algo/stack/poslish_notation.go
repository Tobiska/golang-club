package main

import (
	"fmt"
	"strconv"
)

type stack []int

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() int {
	if len(*s) <= 0 {
		return 0
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func evalRPN(tokens []string) int {
	var stk stack
	for _, t := range tokens {
		n, err := strconv.Atoi(t)
		if err == nil {
			stk.push(n)
			continue
		}
		switch t {
		case "+":
			stk.push(stk.pop() + stk.pop()) // order doesn't matter
		case "-":
			b, a := stk.pop(), stk.pop()
			stk.push(a - b)
		case "*":
			stk.push(stk.pop() * stk.pop()) // order doesn't matter
		case "/":
			b, a := stk.pop(), stk.pop()
			stk.push(a / b)
		}
	}
	return stk.pop()
}

func main() {
	fmt.Println(evalRPN([]string{"1", "2", "+"}))
}
