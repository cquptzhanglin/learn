package main

import (
	"fmt"
)

func rangeString() {
	var str1 = "go is a beautyful language"
	for pos, val := range str1 {
		fmt.Printf("the charater in position %d is %c\n", pos, val)
	}
	var str2 = "cqupt:重庆邮电大学"
	for pos, val := range str2 {
		fmt.Printf("the character in position %d of str2 is %c\n", pos, val)
	}

	fmt.Println("index  int(rune)   rune        char    bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %5d        %U   '%c'    %X\n", index, rune, rune, rune, []byte(string(rune)))
	}
	for pos := range str2 {
		fmt.Printf("position %d of str2 is %c\n", pos, str2[pos])
	}
}
