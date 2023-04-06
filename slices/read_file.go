package main

import (
	"os"
	"regexp"
)

func main2() {
	IncorrectReadFile()
	CorrectReadFile()
}

var (
	letters = regexp.MustCompile("[a-z]+")
)

func IncorrectReadFile() []byte {
	bt, _ := os.ReadFile("file.pdf")
	return letters.Find(bt)
}

func CorrectReadFile() []byte {
	//тут gc всё почистит
	bt, _ := os.ReadFile("file.pdf")
	lettersBt := letters.Find(bt)
	lettersBtCopy := make([]byte, len(lettersBt))
	copy(lettersBtCopy, lettersBt)

	return lettersBtCopy
}
