package main

// import (
// 	"fmt"
// 	"time"
// )

// const (
// 	//HEIGHT is the screen height
// 	HEIGHT = 1920
// 	//WIDTH is the screen width
// 	WIDTH = 1020
// )

// type pixel int

// var screen [WIDTH][HEIGHT]pixel

// func main() {
// 	//var a, b int = 1, 2
// 	//println(a, b)
// 	var arr1 = [5]int{1, 2, 3, 4, 5}
// 	var arr3 [5]int
// 	var arr2 []int
// 	//arr2 := new([5]int)
// 	f1(arr1)
// 	arr3 = arr1
// 	arr2 = arr1[:]
// 	arr2[2] = 100
// 	f2(arr2)
// 	f1(arr3)
// 	sc()
// 	start := time.Now()
// 	fmt.Println("end main", sumsc(screen))
// 	end := time.Now()
// 	fmt.Println("total run time :", end.Sub(start))
// }
// func f1(arr [5]int) {
// 	fmt.Println(arr)
// }
// func f2(arr []int) {
// 	fmt.Println(arr)
// }

// func sc() {
// 	for x := 0; x < WIDTH; x++ {
// 		for y := 0; y < HEIGHT; y++ {
// 			screen[x][y] = pixel(x * y)
// 		}
// 	}
// }

// func sumsc(a [WIDTH][HEIGHT]pixel) (sum int) {
// 	for _, val := range a {
// 		for _, v := range val {
// 			sum += int(v)
// 		}
// 	}
// 	return
// }
