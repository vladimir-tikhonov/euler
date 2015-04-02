package main

import (
	"fmt"
)

func main() {
	value := 600851475143
	for i := 2; i < value/2; i++ {
		if value%i == 0 {
			value /= i
		}
	}
	fmt.Println(value)
}
