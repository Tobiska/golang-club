package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}
	sRunes := []rune(s)
	var subSeqPtr int
	for _, c := range t {
		if c == sRunes[subSeqPtr] {
			subSeqPtr++
			if subSeqPtr == len(sRunes) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abk", "ahbgdc"))
}
