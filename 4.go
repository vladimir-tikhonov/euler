package main

import (
	"fmt"
)

func main() {
	maxPalindrome := 0
	for i := 100; i < 1000; i++ {
		palindrome := int(string(i) + reverse(string(i)))
		if isMultiplicationOfTheeDigits(palindrome) && maxPalindrome < palindrome {
			maxPalindrome = palindrome
		}
	}
	fmt.Println(maxPalindrome)
}

func isMultiplicationOfTheeDigits(value int) (result bool) {
	firstDivisor, isDivisorFound := firstThreeDigitDivisor(value)
	if !isDivisorFound {
		return false
	}
	if value%firstDivisor != 0 {
		return false
	}
	secondDivisor := value / firstDivisor
	return secondDivisor <= 999 && secondDivisor >= 100
}

func firstThreeDigitDivisor(value int) (result int, found bool) {
	for i := 100; i < 1000; i++ {
		if value%i == 0 {
			return i, true
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
