package main

import (
	"fmt"
	"strconv"
)

func strconversion2() {
	var origStr = "12333"
	var newStr string
	fmt.Printf("this size of strconv.int = %d\n", strconv.IntSize)
	an, err := strconv.Atoi(origStr)
	if err != nil {
		fmt.Printf("thea origStr is not a integer %s\n", origStr)
		return
	}
	newStr = "123dd"
	an, err = strconv.Atoi(newStr)
	fmt.Printf("the answer is %d %v\n", an,err)
}
