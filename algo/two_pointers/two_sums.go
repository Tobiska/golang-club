package main

import "fmt"

// O(n) O(1)
func twoSum(numbers []int, target int) []int {
	leftPtr := 0
	rightPtr := len(numbers) - 1
	for leftPtr < rightPtr {
		sum := numbers[leftPtr] + numbers[rightPtr]
		if sum == target {
			return []int{leftPtr + 1, rightPtr + 1}
		} else if sum > target {
			rightPtr--
		} else {
			leftPtr++
		}
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
