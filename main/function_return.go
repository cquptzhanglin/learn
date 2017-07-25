package main

// import (
// 	"fmt"
// 	"log"
// 	"runtime"
// )

// func main() {
// 	rf := add2()
// 	fmt.Println(rf(3))
// 	rf = adder2()
// 	fmt.Println(rf(1))
// 	fmt.Println(rf(22))
// 	fmt.Println(rf(333))
// }

// func add2() (f func(int) int) {
// 	rf := func(a int) (b int) {
// 		b = a + 2
// 		return b
// 	}
// 	return rf
// }

// func adder(a int) func(int) int {
// 	rf := func(b int) (c int) {
// 		c = a + b
// 		return c
// 	}
// 	return rf
// }
// func adder2() func(int) int {
// 	var x int
// 	return func(delta int) int {
// 		where := func() {
// 			_, file, line, _ := runtime.Caller(1)
// 			log.Printf("%s:%d", file, line)
// 		}
// 		where()
// 		x += delta
// 		return x
// 	}
// }
