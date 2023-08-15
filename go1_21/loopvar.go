package main

import "fmt"

//GOEXPERIMENT=loopvar go run go1_21/loopvar.go

func main() {
	result := make([]*int, 0)
	for _, v := range []int{1, 2, 3, 4, 5} {
		result = append(result, &v)
	}
	for _, res := range result {
		fmt.Println(*res)
	}
}
