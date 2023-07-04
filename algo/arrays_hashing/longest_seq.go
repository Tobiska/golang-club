package main

import (
	"fmt"
	"sort"
)

// O(n * k) k << n O(n)
func longestConsecutiveHashing(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		seen[num] = struct{}{}
	}

	var maxSequenceCount int
	for num := range seen {
		if _, ok := seen[num-1]; ok {
			continue
		}
		currentSequence := 1
		for {
			if _, ok := seen[num+1]; !ok {
				break
			}
			num, currentSequence = num+1, currentSequence+1
		}
		if maxSequenceCount < currentSequence {
			maxSequenceCount = currentSequence
		}
	}
	return maxSequenceCount
}

// O(n log n) O(1)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	maxSequenceCount := 0
	currentSequence := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 == nums[i] {
			currentSequence++
		} else {
			if maxSequenceCount < currentSequence {
				maxSequenceCount = currentSequence
			}
			currentSequence = 1
		}
	}

	if maxSequenceCount < currentSequence {
		maxSequenceCount = currentSequence
	}

	return maxSequenceCount
}

func main() {
	fmt.Println(longestConsecutiveHashing([]int{1, 2, 3, 5, 0, 4}))
}
