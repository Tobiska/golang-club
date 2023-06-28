package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	onlyCopy()
	copyBuffer()
	copyN()
}

func onlyCopy() {
	// При успешном копировании возвращается err == nil, а не err == EOF
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

func copyBuffer() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf используется здесь...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... переиспользуется также здесь.
	// Нет необходимости выделять дополнительный буфер.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func copyN() {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 5); err != nil {
		log.Fatal(err)
	}
}
