package main

import (
	"fmt"
)

func main() {
	sum := 2
	var (
		first_number  = 1
		second_number = 2
	)
	for second_number < 4000000 {
		first_number, second_number = second_number, first_number+second_number
		if second_number%2 == 0 {
			sum += second_number
		}
	}
	fmt.Println(sum)
}
