package main

// import "fmt"

// func trace(s string) string {
// 	fmt.Println("entering:", s)
// 	return s
// }
// func un(s string) {
// 	fmt.Println("leaving:", s)
// }
// func a() {
// 	defer un(trace("a"))
// 	fmt.Println("in a")
// }
// func b() {
// 	defer un(trace("b"))
// 	fmt.Println("in b")
// 	a()
// }
// func ccc(i int) (z int) {
// 	defer fmt.Println(z)
// 	z = z + i*5
// 	return
// }

// func f1() (result int) {
// 	defer func() {
// 		result++
// 	}()
// 	return 0
// }
// func f2() (r int) {
// 	t := 5
// 	defer func() {
// 		r = t + 5
// 	}()
// 	return t
// }
// func f3() (r int) {
// 	defer func(r int) {
// 		r = r + 5
// 	}(r)
// 	return 1
// }
// func main() {
// 	b()
// 	fmt.Printf("  z = %d\n", ccc(5))
// 	fmt.Printf("  f1 = %d\n", f1())
// 	fmt.Printf("  f2 = %d\n", f2())
// 	fmt.Printf("  f3 = %d\n", f3())
// 	a := []string{"a", "b", "c", "d"}
// 	for i := range a {
// 		fmt.Println("Array item", i, "is", a[i])
// 	}
// }
