package main

import "fmt"

// O(n^2) O(n)
func lengthOfLongestSubstring(s string) int {
	sizeLongestSeq := 0
	mp := make(map[byte]struct{}, len(s))
	for startPtr, endPtr := 0, 0; endPtr < len(s); endPtr++ {
		if _, ok := mp[s[endPtr]]; ok {
			for k := range mp {
				delete(mp, k)
			}
			startPtr++
			endPtr = startPtr
		}
		mp[s[endPtr]] = struct{}{}
		if sizeLongestSeq < len(mp) {
			sizeLongestSeq = len(mp)
		}
	}
	return sizeLongestSeq
}

func indexOf(s []rune, needle rune) int {
	for ind, ch := range s {
		if ch == needle {
			return ind
		}
	}
	return -1
}

// O(n) O(n)
func lengthOfLongestSubstringA(s string) int {
	substring := make([]rune, 0)
	maxLen := 1
	for _, ch := range s {
		if index := indexOf(substring, ch); index != -1 {
			substring = substring[index+1:] // почему нет паники?
		}

		curLen := len(substring)
		if maxLen < curLen {
			maxLen = curLen
		}
		substring = append(substring, ch)
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstringA("11"))
}
