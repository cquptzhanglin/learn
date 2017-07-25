package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print(" fizzbuzz ")
		case i%5 == 0:
			fmt.Print(" buzz ")
		case i%3 == 0:
			fmt.Print(" fizz ")
		default:
			fmt.Printf(" %d ", i)
		}
		if i%10 == 0 {
			fmt.Println()
		}
	}
}
