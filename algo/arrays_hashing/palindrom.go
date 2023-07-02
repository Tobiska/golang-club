package main

import (
	"fmt"
	"unicode"
)

func isLetterOrDigit(s byte) bool {
	if (s < 'A' || s > 'Z') && (s < 'a' || s > 'z') && (s < '0' || s > '9') {
		return false
	}

	return true
}

func toLowerCase(s byte) byte {
	if s <= unicode.MaxASCII {
		if s >= 'A' && s <= 'Z' {
			s += 'a' - 'A'
		}
		return s
	}
	return s
}

func isPalindrome(s string) bool {
	leftPtr := 0
	rightPtr := len(s) - 1

	for leftPtr < rightPtr {
		lcLeft := s[leftPtr]
		lcRight := s[rightPtr]

		if !isLetterOrDigit(lcLeft) {
			leftPtr++
			continue
		}

		if !isLetterOrDigit(lcRight) {
			rightPtr--
			continue
		}

		if toLowerCase(lcLeft) != toLowerCase(lcRight) {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("ba"))
}
