package main

import "fmt"

func sumChars(s string) (sum int) {
	for _, s := range []rune(s) {
		sum += int(s)
	}
	return
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return sumChars(s) == sumChars(t)
}

func main() {
	fmt.Println(isAnagram("ana", "naa"))
	fmt.Println(isAnagram("ana", "nnn"))
}
