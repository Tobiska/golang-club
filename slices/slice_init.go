package main

import "fmt"

func changeSlice(s []int) {
	s[0] = 8
}

func appendSlice(s []int) {
	s = append(s, 8)
}

func Example1() {
	//s := []int{1, 2}
	s := make([]int, 3, 5)
	s[0] = 1
	s[1] = 2

	changeSlice(s)
	// [8,2]

	appendSlice(s)
	// [8,2,0,0,0,0,0,0]

	fmt.Println(s[2])
	// [8,2]
}

func main() {
	Example1()
}
