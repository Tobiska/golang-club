package main

import (
	"os"
	"regexp"
)

func main() {
	IncorrectReadFile()
	CorrectReadFile()
}

var (
	letters = regexp.MustCompile("[a-z]+")
)

func IncorrectReadFile() []byte {
	//При чтении файла выделяется массив байт равный размеру(~1mb)
	//регулярное выражение забирает слайс с тем же базовым массивом,
	//а полезная для нас инфа содержится только в маленьком слайсе,
	//но мы будем продолжать тащить за собой огромнывй массив.
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
