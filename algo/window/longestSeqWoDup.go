package window

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

func main() {
	fmt.Println(lengthOfLongestSubstring("11123456789111"))
}
