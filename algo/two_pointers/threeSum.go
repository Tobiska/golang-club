package main

import (
	"fmt"
	"sort"
)

// O(n^3) O(n)
func threeSumBrut(nums []int) [][]int {
	sumsGroupSl := make([][]int, 0, len(nums)/3)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					sumsGroupSl = append(sumsGroupSl, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return sumsGroupSl
}

// O(n^2) O(n)
func threeSum(nums []int) [][]int {
	lenNums := len(nums)
	sort.Ints(nums)
	resultGroups := make([][]int, 0, len(nums)/3)
	for i := 0; i < lenNums-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, leftPtr, rightPtr := -nums[i], i+1, lenNums-1
		for leftPtr < rightPtr {
			sum := nums[leftPtr] + nums[rightPtr]
			if sum == target {
				resultGroups = append(resultGroups, []int{nums[i], nums[leftPtr], nums[rightPtr]})
				rightPtr--
				leftPtr++
				for leftPtr < rightPtr && nums[leftPtr] == nums[leftPtr-1] {
					leftPtr++
				}
				for leftPtr < rightPtr && nums[rightPtr] == nums[rightPtr+1] {
					rightPtr--
				}
			} else if sum > target {
				rightPtr--
			} else {
				leftPtr++
			}
		}
	}
	return resultGroups
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
