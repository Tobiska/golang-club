package main

import (
	"fmt"
	"sort"
)

func twoSumMap(nums []int, target int) []int {
	numsMp := make(map[int]int, len(nums))
	for currentInd, num := range nums {
		if requiredInd, ok := numsMp[target-num]; ok {
			return []int{requiredInd, currentInd}
		}
		numsMp[num] = currentInd
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	sort.Ints(nums)

	leftPtr := 0
	rightPtr := len(nums) - 1
	for leftPtr < rightPtr {
		sum := nums[leftPtr] + nums[rightPtr]
		if sum == target {
			return []int{leftPtr, rightPtr}
		} else if sum < target {
			leftPtr++
		} else {
			rightPtr--
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSumMap([]int{3, 3, 2}, 6))
}
