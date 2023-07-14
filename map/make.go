package main

import (
	"fmt"
	"unsafe"
)

func makeMap(m *map[int]int) {
	*m = make(map[int]int)
}

func main() {
	var m map[int]int
	var p uintptr
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // equal

	makeMap(&m)
	fmt.Println(m == nil) //
}
