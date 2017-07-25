package main

import (
	"fmt"
)

func forString() {
	var str1 = "go is a beautiful language"
	fmt.Printf("the len of str1 is %d\n", len(str1))
	for ix := 0; ix < len(str1); ix++ {
		fmt.Printf("the location %d of str1's character is %c\n", ix, str1[ix])
	}
	var str2 = "重庆邮电大学"
	fmt.Printf("the len of str2 is %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("the charater localed %d in str2 is %c\n", ix, str2[ix])
	}
}

func forG() {
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}

func forG2() {
	var str = "GGG"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[0 : i+1])
	}
}
