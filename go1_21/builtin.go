package main

import (
	"fmt"
)

func main() {
	fmt.Println(min(1, 2, 3, 0))
	fmt.Println(max(1, 2, 3, 0))
	mp := map[string]string{
		"a": "b",
		"c": "b",
	}

	clear(mp)
	fmt.Println(mp)

	sl := []int{1, 2, 3, 4, 5, 6}
	clear(sl)
	fmt.Println(sl)

}
