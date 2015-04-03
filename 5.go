package main

import (
	"fmt"
)

func main() {
	result := 1
	for i := 2; i <= 20; i++ {
		result *= getNewElement(i)
	}
	fmt.Println(result)
}

func getNewElement(value int) int {
	if value == 2 {
		return 2
	}
	result := value
	for i := 2; i < value; i++ {
		newElementOfI := getNewElement(i)
		if result%newElementOfI == 0 {
			result /= newElementOfI
		}
	}
	return result
}
