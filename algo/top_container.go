package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type Item struct {
	Val  string
	Freq uint
}

func countFrequency(cnt []string) []Item {
	freqLst := make([]Item, 0, len(cnt))
	freq := make(map[string]uint, len(cnt))
	for _, s := range cnt {
		freq[s]++
	}

	for k, v := range freq {
		freqLst = append(freqLst, Item{Val: k, Freq: v})
	}

	return freqLst
}

func mostFrequentlyStr(cnt []string, top int) []Item {
	lst := countFrequency(cnt)
	slices.SortFunc(lst, func(a, b Item) bool {
		return a.Freq > b.Freq
	})
	return lst[:top]
}

// O(n + k*log(k)) - вычисления
// O(n) - память
func main() {
	cnt := []string{
		"aba",
		"aaa",
		"aaa",
		"aaa",
		"bbb",
		"bbb",
		"bbb",
		"bbb",
		"bbb",
		"bbbbbbbbbbb",
		"lolololo",
	}

	fmt.Println(mostFrequentlyStr(cnt, 2))
}
