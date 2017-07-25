package main

// import (
// 	"fmt"
// 	"time"
// )

// const LIM = 41

// var fib [LIM]int64

// func main() {
// 	var res int64 = 0
// 	var i int64
// 	start := time.Now()
// 	for i = 0; i < 40; i++ {
// 		res = calfib(i)
// 		fib[i] = res
// 		fmt.Printf("fib[%d] = %d\n", i, res)
// 	}
// 	end := time.Now()
// 	duaeration := end.Sub(start)
// 	fmt.Printf("total used = %s", duaeration)
// }
// func calfib(i int64) (res int64) {
// 	if fib[i] != 0 {
// 		return fib[i]
// 	}
// 	if i <= 1 {
// 		res = 1
// 	} else {
// 		res = calfib(i-1) + calfib(i-2)
// 	}
// 	//fib[i] = res
// 	return
// }
// import "fmt"

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }
import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	//sort.Strings(keys)
	sort.StringSlice(keys).Sort()
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v \n ", k, barVal[k])
	}
}
