package main

import "fmt"

func mutate(lst []int) {
	lst[0] = 123
}

func add(lst []int) {
	lst = append(lst, 125)
}

func main() {
	lst := []int{1, 2, 3}

	fmt.Printf("before mutate: %v\n", lst)
	mutate(lst)
	fmt.Printf("after mutate: %v\n", lst)

	fmt.Printf("before append: %v\n", lst)
	add(lst) // внешний слайс не изменится
	fmt.Printf("after append: %v\n", lst)

	// [123 2] 3 -> [123 2 3] -> [123 2 125]
	fmt.Printf("before append part: %v\n", lst)
	add(lst[:2])
	fmt.Printf("after append part: %v\n", lst)

	lst2 := make([]int, 2, 2)
	fmt.Printf("before append copy: %v\n", lst2)
	lst22 := append(lst2, 124) // lst22 и lst2 не пересекаются по массиву. Т.к создался новый
	lst22[0] = 11
	fmt.Printf("lst2: %v lst22: %v\n", lst2, lst22)

	lst3 := make([]int, 2, 3)
	fmt.Printf("before append copy: %v\n", lst2)
	lst33 := append(lst3, 124) // lst22 и lst2 пересекаются по массиву. Т.к выделения памяти не произошло
	lst33[0] = 11
	fmt.Printf("lst3: %v lst33: %v\n", lst3, lst33)
}
