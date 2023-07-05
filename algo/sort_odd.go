package main

import (
	"fmt"
	"sort"
)

// Написать функцию, которая принимает массив с четными и нечетными числами и
// возвращает такой же массив, в котором четные числа стоят на своих местах,
// а нечетные остортированы в порядке возрастания

// sortArr([2, 3, 1, 7, 4, 9, 5, 8]) => [2, 1, 3, 5, 4, 7, 9, 8]

func sortArrPoiners(arr []int) []int {
	leftPtr := 0
	rightPtr := len(arr) - 1

	for leftPtr < rightPtr {

	}

	return arr
}

func sortArr(arr []int) []int {
	odd := make([]int, 0, len(arr))
	for _, el := range arr {
		if el%2 == 1 {
			odd = append(odd, el)
		}
	}
	sort.Ints(odd)

	var indOdd int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			arr[i] = odd[indOdd]
			indOdd++
		}
	}

	return arr
}

func main() {
	arr := []int{2, 3, 1, 7, 4, 9, 5, 8}
	fmt.Println(sortArr(arr))
}
