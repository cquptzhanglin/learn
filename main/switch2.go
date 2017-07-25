package main

import (
	"fmt"
)

func switch2() {
	var num = 13
	switch {
	case num < 10:
		fmt.Println("num <10")
	case num >= 10 && num < 20:
		fmt.Println("num>=10 && num <20")
	case num >= 20:
		fmt.Println("num >= 20")

	}

	month := 1
	switch month {
	default:
		fmt.Println("default case")
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("Spring")
	
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fmt.Println("summer")

	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fmt.Println("autumn")
	
	case 10:
		fallthrough
	case 11:
		fallthrough
	case 12:
		fmt.Println("winter")

	}
}
