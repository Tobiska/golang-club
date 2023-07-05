package main

import "fmt"

// решаем с учётом чётности частот символов.
// если чётно частота, то вносим вклад.
// если нечётно делаем чётным.
// если встречался нечётные частоты, то ставим один символ в середину
func longestPalindrome1(s string) int {
	runeS := []rune(s)
	freqMp := make(map[rune]int, len(runeS))
	for _, c := range runeS {
		freqMp[c]++
	}

	odd := false
	sizePm := 0
	for _, f := range freqMp {
		if f%2 == 0 {
			sizePm += f
		} else {
			sizePm += f - 1
			odd = true
		}
	}
	if odd == true {
		sizePm++
	}

	return sizePm
}

// ищем пары.
// если остались без пары, но стаим символ в середину
func longestPalindrome(s string) int {
	pairs := make(map[byte]struct{}, 0)
	var pairCount int
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if _, ok := pairs[cur]; ok {
			pairCount++
			delete(pairs, cur)
		} else {
			pairs[cur] = struct{}{}
		}
	}
	pmSize := pairCount * 2
	if len(pairs) > 0 {
		pmSize++
	}
	return pmSize
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
