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
	for i := 2; i < value/2+1; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
