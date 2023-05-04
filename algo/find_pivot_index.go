package main

import "fmt"

func runningSum1(nums []int) []int {
	sumsNums := make([]int, len(nums))
	sum := 0
	for idx, num := range nums {
		sum += num
		sumsNums[idx] += sum
	}
	return sumsNums
}

func reverse1(nums []int) []int {
	ln := len(nums)
	reverseNums := make([]int, ln)
	copy(reverseNums, nums)
	for i := 0; i < ln/2; i++ {
		reverseNums[i], reverseNums[ln-i-1] = nums[ln-i-1], nums[i]
	}
	return reverseNums
}

func pivotIndex(nums []int) int {
	directRunSum := runningSum1(nums)
	reverseRunSum := runningSum1(reverse1(nums))
	for i := 0; i < len(nums); i++ {
		if directRunSum[i] == reverseRunSum[len(nums)-i-1] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{0, 4, 5, 4, 0}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
	fmt.Println(pivotIndex([]int{0, 0}))
	fmt.Println(pivotIndex([]int{0, 1}))
	fmt.Println(pivotIndex([]int{0, 1, 12}))
}
