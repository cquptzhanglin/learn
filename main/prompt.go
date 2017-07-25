package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"runtime"
// )

// var prompt string = "Enter a digit, e.g. 3 " + "or %s to quit."

// func init() {
// 	if runtime.GOOS == "windows" {
// 		prompt = fmt.Sprintf(prompt, "ctrl + d")
// 	} else {
// 		prompt = fmt.Sprintf(prompt, "ctrl + z")
// 	}
// }

// func main() {
// 	// fmt.Println(prompt)
// 	//ifelse()
// 	//strconversion2()
// 	//switch2()
// 	// forString()
// 	// forG()
// 	// forG2()
// 	// fizzbuzz()
// 	// rangeString()
// 	// LABEL1:
// 	// 	for i := 0; i <= 5; i++ {
// 	// 		for j := 0; j <= 5; j++ {
// 	// 			if j == 4 {
// 	// 				continue LABEL1
// 	// 			}
// 	// 			fmt.Printf("i is: %d, and j is: %d\n", i, j)
// 	// 		}
// 	// 	}
// 	// var i int
// 	// i = 5
// 	// println(i)
// 	// doSomeChange(i)
// 	// println("after call func doSomeChange", i)
// 	// i, j := retTwoInt(1, 2)
// 	// println(i, j)
// 	// j = retTwoInt2(1, 2)
// 	// println(i, j)
// 	// greeting("hello", "simon", "zhanglin", "n", "o", "p")
// 	// typecheck(1, 2, "123", int(32), int16(16), int64(64), "hello go")
// 	defera(5)
// 	deferb()
// 	// a()
// 	// b()
// 	// func1("GO")
// 	// callback(1, 2, add)
// 	// cs := callback2(1, 2, add2)
// 	// fmt.Println("main end", cs)
// 	// fv := func(s string) {
// 	// 	fmt.Println(s)
// 	// }
// 	// fv("hello world")
// 	// anonymousFunction()
// }

// func doSomeChange(i int) {
// 	var ptrI *int
// 	ptrI = &i
// 	*ptrI = 100
// 	println("in func doSomeChange", i)
// }

// func retTwoInt(i, j int) (x, y int) {
// 	return i + 1, j + 2
// }
// func retTwoInt2(i, j int) (y int) {
// 	return y
// }

// func greeting(prefix string, who ...string) {
// 	//fmt.Print(prefix + ":")
// 	for _, val := range who {
// 		fmt.Print(prefix + ":" + val + " \n")
// 	}
// 	// fmt.Println()
// }

// func typecheck(i, j int, values ...interface{}) {
// 	fmt.Printf("i = %d\n", i)
// 	fmt.Printf("j = %d\n", j)
// 	for _, val := range values {
// 		switch val.(type) {
// 		case int:
// 			fmt.Printf("paramter type is int\n")
// 		case int16:
// 			fmt.Printf("paramter type is int16\n")
// 		case int32:
// 			fmt.Printf("paramter type is int32\n")
// 		case int64:
// 			fmt.Printf("paramter type is int64\n")
// 		}
// 	}
// }

// func defera(x int) (ret int) {
// 	//i := 0
// 	ret = x
// 	defer fmt.Printf("defered i = %d\n", ret)
// 	x++
// 	ret = x
// 	fmt.Printf("x = %d\n", x)
// 	x++
// 	ret = x
// 	fmt.Printf("x = %d\n", x)
// 	ret = x
// 	return
// }

// func deferb() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Printf("i = %d\n", i)
// 	}
// }

// func trace(s string) {
// 	fmt.Printf("enter func %s\n", s)
// }

// func untrace(s string) {
// 	fmt.Printf("leave func %s\n", s)
// }

// func a() {
// 	trace("a")
// 	defer untrace("a")
// 	fmt.Println("in a")
// }

// func b() {
// 	trace("b")
// 	defer untrace("b")
// 	fmt.Println("in b")
// }

// func func1(s string) (n int, err error) {
// 	defer func() {
// 		log.Printf("func1(%s) = %d, %v", s, n, err)
// 	}()
// 	return 7, io.EOF
// }

// func add(a, b int) {
// 	fmt.Printf(" %d + %d = %d\n", a, b, a+b)
// }
// func add2(a, b int) (s int) {
// 	s = a + b
// 	fmt.Printf("%d + %d = %d\n", a, b, s)
// 	return
// }
// func callback(a, b int, f func(int, int)) {
// 	f(a, b)
// }
// func callback2(a, b int, f func(int, int) int) (cs int) {
// 	cs = f(a, b)
// 	return
// }

// func anonymousFunction() {
// 	for i := 0; i < 4; i++ {
// 		g := func(x int) (y int) {
// 			fmt.Printf("anonymousFunction used by g i = %d\n", i)
// 			y = 2 * x
// 			return
// 		}
// 		j := g(i)
// 		fmt.Printf(" i = %d j = %d\n", i, j)
// 		// fmt.Printf("g type = %T g value = %v\n", g, g)
// 	}
// }
