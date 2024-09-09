package main

import (
	"fmt"
	"log"
	"os"
)

func binaryRepresentPrint(num int) string {
	binary := make([]byte, 0)
	var div int
	for div > 0 {
		binary = append(binary, byte(num%2))
		num /= 2
	}
	return string(binary)
}

func main() {
	var n int
	_, err := fmt.Fscan(os.Stdin, &n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(binaryRepresentPrint(n))
}
