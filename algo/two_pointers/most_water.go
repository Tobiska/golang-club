package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func area(n, left, right int) int {
	return n * min(left, right)
}

// O(n) O(1)
func maxArea(height []int) int {
	leftPtr := 0
	rightPtr := len(height) - 1
	maxA := -1
	for leftPtr < rightPtr {
		sa := area(rightPtr-leftPtr, height[leftPtr], height[rightPtr])
		if sa > maxA {
			maxA = sa
		}
		if height[leftPtr] <= height[rightPtr] {
			leftPtr++
		} else {
			rightPtr--
		}
	}
	return maxA
}

func main() {
	fmt.Println(maxArea([]int{2, 3, 1, 7, 0, 0, 7, 1, 3, 2}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7, 2}))
}
