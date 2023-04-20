package main

import "fmt"

func foo(i int) {
	if i > 100000000000000000 {
		return
	}
	i++
	_ = make([]int, 1000000000000)
	foo(i)
	fmt.Println("foo")
}

func main() {
	foo(0)
}
