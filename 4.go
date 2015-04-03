package main

import (
	"fmt"
	"strconv"
)

func main() {
	var maxPalindrome int64 = 0
	for i := 100; i < 1000; i++ {
		palindromePart := strconv.FormatInt(int64(i), 10)
		palindrome, err := strconv.ParseInt(palindromePart+reverse(palindromePart), 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		if isMultiplicationOfTheeDigits(palindrome) && maxPalindrome < palindrome {
			maxPalindrome = palindrome
		}
	}
	fmt.Println(maxPalindrome)
}

func isMultiplicationOfTheeDigits(value int64) (result bool) {
	lastDivisor, isLastFound := lastThreeDigitDivisor(value)
	if !isLastFound {
		return false
	}
	if value%lastDivisor != 0 {
		return false
	}
	secondDivisor := value / lastDivisor
	return secondDivisor <= 999 && secondDivisor >= 100
}

func lastThreeDigitDivisor(value int64) (result int64, found bool) {
	for i := 999; i > 99; i-- {
		if value%int64(i) == 0 {
			return int64(i), true
		}
	}
	return -1, false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
