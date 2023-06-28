package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	//Pipe создает синхронный pipe в памяти.
	//Его можно использовать для соединения кода, ожидающего io.Reader, с кодом, ожидающим io.Writer.
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some text to be read\n")
		w.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	fmt.Print(buf.String())
}
