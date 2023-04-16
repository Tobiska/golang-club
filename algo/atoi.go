package main

import (
	"fmt"
)

func tryToDigit(r rune) (int, bool) {
	numAscii := int(r - '0')
	if numAscii < 0 || numAscii > 9 {
		return numAscii, false
	}
	return numAscii, true
}

func updateResult(r1, r2 rune, factor, res int) (uRes, uFactor int, uSign bool) {
	uRes = res
	uFactor = factor

	numA1, f1 := tryToDigit(r1)
	numA2, f2 := tryToDigit(r2)
	if f1 && numA2 == -3 {
		uSign = true
	}
	if f1 {
		uRes += numA1 * uFactor
		uFactor *= 10
	}
	if f2 {
		uRes += numA2 * uFactor
		uFactor *= 10
	}
	return
}

func reverse(rs []rune) {
	ln := len(rs)
	for i := 0; i < ln/2; i++ {
		rs[i], rs[ln-i-1] = rs[ln-i-1], rs[i]
	}
}

//O(2 * n) = O(n) - время
//O(n) - память

func myAtoi(s string) (res int) {
	runes := []rune(s)

	if len(runes) == 1 {
		dt, f := tryToDigit(runes[0])
		if f {
			return dt
		}
		return 0
	}

	reverse(runes)
	factor := 1
	sign := false
	for i := 1; i < len(runes); i += 2 {
		res, factor, sign = updateResult(runes[i-1],
			runes[i],
			factor, res)
	}

	if len(runes)%2 == 1 {
		res, factor, sign = updateResult(
			runes[len(runes)-1],
			runes[len(runes)-2],
			factor, res)
	}

	if sign {
		res *= -1
	}
	return res
}

func main() {
	fmt.Printf("%d", myAtoi("-22376767676576767"))
}
