package main

import (
	"fmt"
	"sort"
)

// O(n) O(n)
func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]struct{}, len(nums))
	for _, c := range nums {
		if _, ok := numsMap[c]; ok {
			return true
		} else {
			numsMap[c] = struct{}{}
		}
	}
	return false
}

// O(n * log N), O(1)
func containsDuplicateSort(nums []int) bool {
	if len(nums) <= 0 {
		return false
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 4, 5, 6}))
}
