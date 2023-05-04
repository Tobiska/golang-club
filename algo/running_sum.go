package main

func runningSum(nums []int) []int {
	sumsNums := make([]int, len(nums))
	sum := 0
	for idx, num := range nums {
		sum += num
		sumsNums[idx] += sum
	}
	return sumsNums
}

func main() {
	runningSum([]int{1, 2, 3, 4, 5})
}
