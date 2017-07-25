package main

import (
	"fmt"
)

func switch1() {
	var i = 98
	switch i {
	case 98:
	case 99:
		fmt.Println("i = 98 or 99")
	case 100:
		fmt.Println("i = 100")
		fallthrough
	default:
		fmt.Println("default")
	}
}
