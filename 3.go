package main

import (
	"fmt"
)

func main() {
	value := 600851475143
	for i := 72; i > 1; i-- {
		if value%i == 0 {
			fmt.Println(i)
			return
		}
	}
}
