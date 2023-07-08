package main

import "fmt"

func del(sl []int, d int) []int {
	return append(sl[:d], sl[d+1:]...)
}

func main() {
	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(del(sl, 2))
}
