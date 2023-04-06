package main

import "fmt"

func main1111() {
	slc := []int{1, 2, 3, 4}
	slcCopy1 := make([]int, 2)
	slcCopy2 := make([]int, 5)

	n := copy(slcCopy1, slc)
	copy(slcCopy2, slc)
	fmt.Printf("num: %d copy1: %v, copy2: %v src: %v", n, slcCopy1, slcCopy2, slc)

	fmt.Println(cap(slcCopy2[:2]))
}
