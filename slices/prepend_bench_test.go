package main

import "testing"

func BenchmarkPrepend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prepend(100, []int{1, 2, 3, 4, 5})
	}
}
