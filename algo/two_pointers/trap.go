package main

import "fmt"

func trap(height []int) int {
	leftPtr, rightPtr := 0, len(height)-1
	leftMax, rightMax := height[leftPtr], height[rightPtr]
	waterTrapped := 0
	for leftPtr < rightPtr {
		if height[leftPtr] < height[rightPtr] {
			leftPtr++
			if leftMax < height[leftPtr] {
				leftMax = height[leftPtr]
			}
			waterTrapped += leftMax - height[leftPtr]
		} else {
			rightPtr--
			if rightMax < height[rightPtr] {
				rightMax = height[rightPtr]
			}
			waterTrapped += rightMax - height[rightPtr]
		}
	}
	return waterTrapped
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
