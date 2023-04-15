package main

import (
	"fmt"
	"unsafe"
)

type MyStructA struct {
	i1 int8
	i2 int8
	fl float64 //8
}

func main() {
	fmt.Println(unsafe.Sizeof(MyStructA{}))
}
