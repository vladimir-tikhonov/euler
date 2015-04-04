package main

import (
	"fmt"
)

func main() {
	var (
		currentValue      = 3
		currentPrimeIndex = 1
	)

	for ; ; currentValue += 2 {
		if isPrime(currentValue) {
			currentPrimeIndex++
			if currentPrimeIndex == 10001 {
				break
			}
		}
	}
	fmt.Println(currentValue)
}

func isPrime(value int) bool {
	if value%2 == 0 {
		return false
	}
	for i := 3; i < value/2+1; i += 2 {
		if value%i == 0 {
			return false
		}
	}
	return true
}
