package main

import (
	"fmt"
	"iter"
)

func countTo() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < 6; i++ {
			yield(i)
		}
	}
}

func main() {
	for d := range countTo() {
		fmt.Println(d)
	}
}
