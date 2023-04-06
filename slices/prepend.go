package main

import "fmt"

func prepend(x int, sl []int) []int {
	//newSl := make([]int, 0, len(sl) + 1)
	return append([]int{x}, sl...)
}

func main7() {
	fmt.Println(prepend(100, []int{1, 2, 3}))
}
