package main

import (
	"fmt"
)

func ifelse() {
	var first = 10
	var second int
	if first < 0 {
		fmt.Printf("first= %d less then 0\n", first)
	} else if first >= 0 && first <= 5 {
		fmt.Printf("first= %d between [0 ,5]\n", first)
	} else {
		fmt.Printf("first = %d greater than 5\n", first)
	}
	if second = 5; second > 0 {
		fmt.Printf("second = %d is greater than 0\n", second)
	} else {
		fmt.Printf("second = %d is less than 0\n", second)
	}
}
