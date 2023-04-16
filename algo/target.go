package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

/*
	Найти 2 числа в слайсе data, сумма который даёт число t
*/

func findTarget(data []int, t int) (int, int) {
	slices.SortFunc(data, func(a, b int) bool {
		return a < b
	})

	leftPtr := 0
	rightPtr := len(data) - 1
	for leftPtr < rightPtr {
		sum := data[leftPtr] + data[rightPtr]
		if sum == t {
			return data[leftPtr], data[rightPtr]
		}
		if sum < t {
			leftPtr++
		} else {
			rightPtr--
		}
	}
	return -1, -1
}

func main() {
	fmt.Println(findTarget([]int{1, 4, 2}, 5))
}
