package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func promCalc(cost, am int) int {
	return (am - am/3) * cost
}

func calcCost(costs []int) (result int) {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	curAm := 1
	for i := 1; i < len(costs); i++ {
		if costs[i-1] == costs[i] {
			curAm++
		} else {
			result += promCalc(costs[i-1], curAm)
			curAm = 1
		}
	}
	result += promCalc(costs[len(costs)-1], curAm)
	return
}

func ScanPurchases() (results []int, err error) {
	var countPurchases int
	_, err = fmt.Scanf("%d \n", &countPurchases)
	if err != nil {
		log.Fatalln(err)
	}

	results = make([]int, 0, countPurchases)
	for i := 0; i < countPurchases; i++ {
		var countProduct int
		_, err = fmt.Scanf("%d \n", &countProduct)
		if err != nil {
			return
		}

		productSl := make([]int, countProduct)
		for j := 0; j < countProduct; j++ {
			fmt.Fscanf(os.Stdin, "%d", &productSl[j])
		}
		results = append(results, calcCost(productSl))
		fmt.Fscanf(os.Stdin, "\n")
	}
	return
}

func searchInsert(nums []int, target int) int {
	leftInd, rightInd := 0, len(nums)-1
	for leftInd < rightInd {
		mid := leftInd + (leftInd+rightInd)/2 // 0
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			rightInd = mid // 1
		} else {
			leftInd = mid + 1 // 0
		}
	}

	return leftInd
}

func main() {

	//results, err := ScanPurchases()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for _, res := range results {
	//	fmt.Println(res)
	//}

	fmt.Println(searchInsert([]int{1, 2, 3, 5}, 4))
}
