package main

import (
	"fmt"
	"log"
)

func main() {
	var countPairs int
	_, err := fmt.Scanf("%d\n", &countPairs)
	if err != nil {
		log.Fatalln(err)
	}

	result := make([]int, 0, countPairs)
	for i := 0; i < countPairs; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		result = append(result, a+b)
	}

	for _, c := range result {
		fmt.Printf("%d\n", c)
	}
}
