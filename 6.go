package main

import (
	"fmt"
)

func main() {
	const count int = 100
	result := 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if i != j {
				result += i * j
			}
		}
	}
	fmt.Println(result)
}
