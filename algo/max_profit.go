package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPrice := prices[0]
	maxDelta := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			delta := prices[i] - minPrice
			if delta > maxDelta {
				maxDelta = delta
			}
		}
	}
	return maxDelta
}

func main() {
	fmt.Println(maxProfit([]int{0, 1, 2, 3, 10})) //10
	fmt.Println(maxProfit([]int{10, 9, 5}))       //0
}
