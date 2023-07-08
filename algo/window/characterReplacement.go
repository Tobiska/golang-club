package main

import "fmt"

// k = 1
// [A A C B D E]

func replacementSameSubstringValid(substring string, k int) bool {
	freqMp := make(map[byte]int, len(substring))
	maxf := 0
	for i := 0; i < len(substring); i++ {
		freqMp[substring[i]]++
		if maxf < freqMp[substring[i]] {
			maxf = freqMp[substring[i]]
		}
	}

	return len(substring)-maxf <= k
}

// O(n) O(n)
func characterReplacement(s string, k int) int {
	freqChar := make(map[byte]int, len(s))
	var startPtr, endPtr, maxFreq int
	for startPtr, endPtr = 0, 0; endPtr < len(s); endPtr++ {
		freqChar[s[endPtr]]++
		if maxFreq < freqChar[s[endPtr]] {
			maxFreq = freqChar[s[endPtr]]
		}

		if (endPtr-startPtr+1)-maxFreq > k {
			freqChar[s[startPtr]]--
			startPtr++
		}

	}
	return endPtr - startPtr
}

func main() {
	fmt.Println(characterReplacement("ACB", 1))
}
