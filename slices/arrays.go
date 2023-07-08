package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	arr5 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("size: %d bt\n", reflect.TypeOf(arr5).Size()) // 40 bt

	arr5Copy := arr5
	fmt.Printf("ptr1: %d, ptr2: %d\n",
		reflect.ValueOf(&arr5Copy).Pointer(), // not equal
		reflect.ValueOf(&arr5).Pointer(),
	)

	arr5Ptr := &arr5

	arr5[0] = 50
	fmt.Printf("arr: %v, arr copy: %v\n", arr5, arr5Copy) // !=
	fmt.Printf("arr: %v, arr copy: %v\n", arr5, *arr5Ptr) // ==

	arr5Struct := [5]struct{}{{}, {}, {}}
	fmt.Printf("arr: %v\n", unsafe.Sizeof(arr5Struct)) // 0

	var grid [8][8]int
	fmt.Printf("grid: %v size: %d\n", grid, unsafe.Sizeof(grid)) // 8 * 8 * 8 = 2^9 = 512

	var zeroValueArr [0]int
	fmt.Printf("zero value: %v\n", zeroValueArr) // 0
}
