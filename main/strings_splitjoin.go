package main

// import (
// 	"fmt"
// 	"math"
// 	"runtime"
// 	"strings"
// 	"time"
// )

// func Sqrt(x float64) (b int, z float64) {
// 	z = float64(1)
// 	old := z
// 	b = 0
// 	for {
// 		b++
// 		old = z
// 		z = z - ((z*z - x) / 2 * z)
// 		// fmt.Printf("i = %d  old = %f   z = %f\n", i, old, z)
// 		if math.Abs(z-old) < 0.0001 {
// 			break
// 		}

// 	}

// 	return
// }

// func main() {
// 	str := "The quick brown fox jumps over the lazy dog"
// 	s1 := strings.Fields(str)
// 	fmt.Printf("Splited in slice %v \n", s1)
// 	for _, ss := range s1 {
// 		fmt.Printf("%s\n", ss)
// 	}

// 	str2 := "GO1|The ABC of Go|25"
// 	s2 := strings.Split(str2, "|")
// 	fmt.Printf("Splited in slice %v \n", s2)
// 	for _, val := range s2 {
// 		fmt.Printf("%s\n", val)
// 	}
// 	str3 := strings.Join(s2, ";")
// 	fmt.Printf("%s\n", str3)
// 	x, y := Sqrt(2)
// 	fmt.Printf("%d  %f\n", x, y)

// 	fmt.Print("Go runs on ")
// 	switch os := runtime.GOOS; os {

// 	case "windows":
// 		fmt.Println("windows")
// 	case "darwin":
// 		fmt.Println("OS X.")
// 	case "linux":
// 		fmt.Println("Linux.")
// 	default:
// 		// freebsd, openbsd,
// 		// plan9, windows...
// 		fmt.Printf("%s.", os)
// 	}

// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon.")
// 	default:
// 		fmt.Println("Good evening.")
// 	}
// }
