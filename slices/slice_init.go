package main

import "fmt"

func mutateSlice(s []int, idx, val int) {
	s[idx] = val
}

func appendSlice(s []int, val int) {
	s = append(s, val)
}

func Example1() {
	nums := make([]int, 2, 3)
	fmt.Println(nums)

	appendSlice(nums, 1024)
	fmt.Println(nums)

	mutateSlice(nums, 2, 512)
	fmt.Println(nums)
}

func main6() {
	Example1()
}
