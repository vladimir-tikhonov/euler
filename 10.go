package main

import (
	"fmt"
	"math"
)

func main() {
	var sum int64 = 2
	for i := 3; i < 2000000; i += 2 {
		if isPrime(i) {
			sum += int64(i)
		}
	}
	fmt.Println(sum)
}

func isPrime(value int) bool {
	if value%3 == 0 || value%5 == 0 || value%7 == 0 || value%11 == 0 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(value))) + 1
	for i := 3; i < maxDivisor; i += 2 {
		if value%i == 0 {
			return false
		}
	}
	return true
}
