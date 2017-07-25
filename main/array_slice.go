package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	// var arr1 [6]int
// 	// var slice1 []int = arr1[2:2]
// 	// for i := 0; i < 6; i++ {
// 	// 	arr1[i] = i
// 	// }
// 	// fmt.Println("len of slice1 = ", len(slice1), "capacity of slice1 = ", cap(slice1))
// 	// for i := 0; i < len(slice1); i++ {
// 	// 	fmt.Println(slice1[i])
// 	// }
// 	// slice1 = slice1[:cap(slice1)]
// 	// fmt.Println("len of slice1 = ", len(slice1), "capacity of slice1 = ", cap(slice1))
// 	// for i := 0; i < len(slice1); i++ {
// 	// 	fmt.Println(slice1[i])
// 	// }

// 	// slfrom := []int{1, 2, 3}
// 	// slto := make([]int, 10)

// 	// n := copy(slto, slfrom)
// 	// fmt.Println(slto)
// 	// fmt.Printf("Copied %d elements\n", n) // n == 3

// 	// sl3 := []int{1, 2, 3}
// 	// fmt.Println(sl3, "len", len(sl3), "cap", cap(sl3))
// 	// sl3 = append(sl3, 4, 5, 6)
// 	// fmt.Println(sl3, "len", len(sl3), "cap", cap(sl3))
// 	s1, s2, _ := split2("hello go", 5)
// 	fmt.Println(s1)
// 	fmt.Println(s2)
// 	s2 = convert("Google")
// 	fmt.Println(s2)
// 	s2 = convert2("Google重庆邮电大学")
// 	fmt.Println(s2)
// }

// //编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，
// //然后返回两个分割后的字符串。
// func split2(s string, i int) (ret1, ret2 string, err error) {
// 	if i > len(s) {
// 		err = errors.New("split error")
// 	}
// 	ret1 = s[:i]
// 	ret2 = s[i:]
// 	return
// }

// //编写一个程序，要求能够反转字符串，
// //即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。
// func convert(s string) (ret string) {
// 	orig := []byte(s)
// 	l := len(orig)
// 	for i := 0; i < l/2; i++ {
// 		orig[i], orig[l-1-i] = orig[l-1-i], orig[i]
// 	}
// 	ret = string(orig)
// 	return
// }

// func convert2(s string) (ret string) {
// 	orig := []int32(s)
// 	l := len(orig)
// 	for i := 0; i < l/2; i++ {
// 		orig[i], orig[l-1-i] = orig[l-1-i], orig[i]
// 	}
// 	ret = string(orig)
// 	return
// }
