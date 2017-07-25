package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// )

// func main() {
// 	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
// 	pat := "[0-9]+.[0-9]+" //正则

// 	f := func(s string) (ret string) {
// 		fmt.Println(s)
// 		v, _ := strconv.ParseFloat(s, 32)
// 		fmt.Println(v)
// 		ret = strconv.FormatFloat(v*2, 'f', 2, 32)
// 		fmt.Println(ret)
// 		return
// 	}

// 	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
// 		fmt.Println("Match Found!")
// 	}

// 	re, _ := regexp.Compile(pat)
// 	fmt.Println(re)
// 	//将匹配到的部分替换为"##.#"
// 	str := re.ReplaceAllString(searchIn, "##.#")
// 	fmt.Println(str)
// 	//参数为函数时
// 	str2 := re.ReplaceAllStringFunc(searchIn, f)
// 	fmt.Println(str2)
// }
