package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	for i := range s {
		if t[i] != s[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("ab", "ba"))
}
