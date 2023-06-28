package main

import (
	"errors"
	"fmt"
)

func main() {
	var EOF = errors.New("EOF")
	fmt.Println(EOF)

}
