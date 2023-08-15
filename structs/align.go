package main

import (
	"fmt"
	"unsafe"
)

type MyStructA struct {
	f  bool
	i1 int32
	fl float64
}

func main() {
	fmt.Println(unsafe.Sizeof(MyStructA{}))
}
